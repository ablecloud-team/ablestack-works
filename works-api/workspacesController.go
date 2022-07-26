package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"strings"
	//"works-api/workspaces"
)

// getWorkspaces godoc
// @Summary 워크스페이스 리스트를 조회하는 API
// @Description 워크 스페이스 리스트를 조회하는 API 입니다.
// @Tags workspaces
// @Accept  json
// @Produce  json
// @Router /api/v1/workspace [get]
// @Success 200 {object} map[string]interface{}
func getWorkspaces(c *gin.Context) {
	returnData := map[string]interface{}{}
	returnCode := http.StatusNotFound
	resultWorkspaceList, err1 := selectWorkspaceList("all")
	resultWorkspaceCount, err2 := selectCountWorkspace()
	if err1 != nil && err2 != nil {
		log.Errorf("workspace error [%v], count error [%v]", err1, err2)
		returnData["listError"] = err1
		returnData["countError"] = err2
	} else {
		returnData["list"] = resultWorkspaceList
		returnData["listTotalCount"] = resultWorkspaceCount
		returnCode = http.StatusOK
	}

	c.JSON(returnCode, gin.H{
		"result": returnData,
	})
}

// getWorkspacesDetail godoc
// @Summary 워크스페이스 리스트를 조회하는 API
// @Description 워크 스페이스 리스트를 조회하는 API 입니다.
// @Tags workspaces
// @Accept  json
// @Produce  json
// @Param workspaceUuid path string true "워크스페이스 UUID"
// @Router /api/v1/workspace/:workspaceUuid [get]
// @Success 200 {object} map[string]interface{}
func getWorkspacesDetail(c *gin.Context) {
	workspaceUuid := c.Param("workspaceUuid")
	returnCode := http.StatusNotFound
	resultReturn := map[string]interface{}{}
	workspaceList, _ := selectWorkspaceList(workspaceUuid)
	if len(workspaceList) != 0 {
		workspaceInfo := workspaceList[0]
		paramsTemplate := []MoldParams{
			{"templatefilter": "executable"},
			{"id": workspaceInfo.TemplateUuid},
		}

		templateResult := getTemplate(paramsTemplate)

		paramsServiceOffering := []MoldParams{
			{"id": workspaceInfo.ComputeOfferingUuid},
		}

		serviceOfferingResult := getComputeOffering(paramsServiceOffering)

		paramsNetwork := []MoldParams{
			{"id": workspaceInfo.NetworkUuid},
		}
		networkResult := getNetwork(paramsNetwork)

		instanceList, _ := selectInstanceList(workspaceUuid, WorkspaceString)
		paramsInstanceList := []MoldParams{
			{"domainid": os.Getenv("MoldDomainId")},
		}
		virtualMachineList := getListVirtualMachinesMetrics(paramsInstanceList)
		listVirtualMachinesMetrics := ListVirtualMachinesMetrics{}
		virtualMachineInfo, _ := json.Marshal(virtualMachineList["listvirtualmachinesmetricsresponse"])
		json.Unmarshal([]byte(virtualMachineInfo), &listVirtualMachinesMetrics)
		for i, v := range instanceList {
			for _, v1 := range listVirtualMachinesMetrics.Virtualmachine {
				if v.MoldUuid == v1.Id {
					instanceList[i].MoldStatus = v1.State
					break
				}
			}
		}

		groupDetail, _ := selectGroupDetail(workspaceInfo.Name)
		var groupData map[string]interface{}
		err := json.NewDecoder(groupDetail.Body).Decode(&groupData)
		if err != nil {

		}

		//workspacePolicy, _ := selectWorkspacePolicyList(workspaceInfo.Uuid)
		policyList := policyList(workspaceInfo.Uuid)

		//log.Debugf("[%v] [%v]", workspacePolicy, workspacePolicy)
		//var workspacePolicyData []map[string]interface{}
		//err = json.NewDecoder(workspacePolicy.Body).Decode(&workspacePolicyData)
		//if err != nil {
		//	log.Errorf("workspacePolicy error [%v]", err)
		//}

		resultReturn["workspaceInfo"] = workspaceInfo
		resultReturn["templateInfo"] = templateResult["listtemplatesresponse"]
		resultReturn["serviceOfferingInfo"] = serviceOfferingResult["listserviceofferingsresponse"]
		resultReturn["networkInfo"] = networkResult["listnetworksresponse"]
		resultReturn["instanceList"] = instanceList
		resultReturn["groupDetail"] = groupData
		resultReturn["workspacePolicy"] = policyList
		returnCode = http.StatusOK
	} else {
		resultReturn["message"] = fmt.Sprintf("There is no workspace for that UUID. [%v]", workspaceUuid)
	}

	c.JSON(returnCode, gin.H{
		"result": resultReturn,
	})
}

// getOffering godoc
// @Summary 템플릿 및 각종 오퍼링을 조회하는 API
// @Description 템플릿, 서비스 오퍼링, 디스크 오퍼링 정보를 Mold 에서 조회하여 제공하는 API 입니다.
// @Tags offering
// @Accept  json
// @Produce  json
// @Router /api/v1/offering [get]
// @Success 200 {object} map[string]interface{}
func getOffering(c *gin.Context) {
	returnCode := http.StatusOK
	//params1 := []MoldParams{
	//	{"command": "listServiceOfferings"},
	//}
	paramsComputerOffering := []MoldParams{}
	paramsTemplate := []MoldParams{}
	//{"templatefilter": "all"},
	//result := map[string]interface{}{}
	templateResult := getListDesktopMasterVersions(paramsTemplate)
	serviceOfferingResult := getComputeOffering(paramsComputerOffering)
	c.JSON(returnCode, gin.H{
		//"result":              result,
		"templateList":        templateResult,
		"serviceOfferingList": serviceOfferingResult,
	})
}

// postWorkspaces godoc
// @Summary 워크스페이스를 추가하는 API
// @Description 워크스페이를 추가하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags workspaces
// @Param name path string true "워크스페이스 이름"
// @Param description path string true "워크스페이스 설명"
// @Param type path string true "워크스페이스 타입(Desktop or Application)"
// @Param computeOfferingUuid path string true "워크스페이스에서 사용할 Compute offering UUID"
// @Param templateUuid path string true "워크스페이스에서 사용할 Template UUID"
// @Param shared path bool true "워크스페이스에서 Shard 여부 전용이면 'false', 공용이면 'true'"
// @Param rdpPort path int true "워크스페이스의 데스크탑의 RDP 접속 Port"
// @Param rdpAccessAllow path int true "워크스페이스의 RDP 접속 허용 여부"
// @Router /api/v1/workspace [POST]
// @Success 200 {object} map[string]interface{}
func postWorkspaces(c *gin.Context) {
	workspace := Workspace{}
	result := map[string]interface{}{}
	resultCode := http.StatusNotFound
	workspace.Uuid = getUuid()
	workspace.Name = c.PostForm("name")
	workspace.Description = c.PostForm("description")
	workspace.WorkspaceType = c.PostForm("type")
	workspace.TemplateUuid = c.PostForm("templateUuid")
	workspace.MasterTemplateName = c.PostForm("masterTemplateName")
	workspace.ComputeOfferingUuid = c.PostForm("computeOfferingUuid")
	workspace.Shared, _ = strconv.ParseBool(c.PostForm("shared"))
	workspace.NetworkUuid = selectNetworkDetail()
	workspace.Postfix = 0

	resultInsertGroup, err := insertGroup(workspace.Name)
	if resultInsertGroup.Status == Created201 {
		resultInsertPolicyRemotefx, err := insertPolicyRemotefx(workspace.Name)
		log.Infof("resultInsertPolicyRemotefx [%v], [%v]", resultInsertPolicyRemotefx, err)
	}

	if err != nil {
		log.Errorf("An error occurred during DC API communication created a workspace group. [%v]", err)
	}
	res := map[string]interface{}{}
	err = json.NewDecoder(resultInsertGroup.Body).Decode(&res)
	result["resultInsertGroup"] = res
	if resultInsertGroup.Status == Created201 {
		resultInsertWorkspace, _ := insertWorkspace(workspace)
		//insertWorkspacePolicy(workspace)
		log.Info(resultInsertWorkspace)
		result["insertWorkspace"] = resultInsertWorkspace
		if resultInsertWorkspace["status"] == http.StatusOK {
			//instanceUuid := getUuid()
			resultDeploy, instanceUuid := getDeployVirtualMachine(workspace, WorkspaceString)
			log.Infof("Mold 통신 결과값 [%v]", resultDeploy)
			if resultDeploy["deployvirtualmachineresponse"].(map[string]interface{})["errorcode"] != nil {
				result["resultDeploy"] = resultDeploy
				result["resultDeploy"].(map[string]interface{})["message"] = MessageSignatureError
			} else {
				paramsMold := []MoldParams{
					{"id": resultDeploy["deployvirtualmachineresponse"].(map[string]interface{})["id"].(string)},
				}
				resultMoldInstanceInfo := getListVirtualMachinesMetrics(paramsMold)
				listVirtualMachinesMetrics := ListVirtualMachinesMetrics{}
				virtualMachineInfo, _ := json.Marshal(resultMoldInstanceInfo["listvirtualmachinesmetricsresponse"])
				err = json.Unmarshal([]byte(virtualMachineInfo), &listVirtualMachinesMetrics)

				workspaceList, _ := selectWorkspaceList(workspace.Uuid)
				workspaceInfo := workspaceList[0]
				instanceInfo := Instance{}
				instanceInfo.Uuid = instanceUuid
				instanceInfo.MoldUuid = resultDeploy["deployvirtualmachineresponse"].(map[string]interface{})["id"].(string)
				instanceInfo.Name = listVirtualMachinesMetrics.Virtualmachine[0].Displayname
				instanceInfo.WorkspaceUuid = workspaceInfo.Uuid
				instanceInfo.WorkspaceName = workspaceInfo.Name
				instanceInfo.Ipaddress = listVirtualMachinesMetrics.Virtualmachine[0].Ipaddress
				resultInsertInstance := insertInstance(instanceInfo)
				if resultInsertInstance["status"] == http.StatusOK {
					resultCode = http.StatusOK
					result["resultInsertDeploy"] = resultInsertInstance
					params := []MoldParams{
						{"resourceids": instanceInfo.MoldUuid},
						{"resourcetype": UserVm},
						{"tags[0].key": ServiceDaaS},
						{"tags[0].value": AblecloudWorks},
						{"tags[1].key": WorkspaceName},
						{"tags[1].value": workspaceInfo.Name},
						{"tags[2].key": ClusterName},
						{"tags[2].value": os.Getenv("ClusterName")},
					}
					resultGetCreateTags := getCreateTags(params)
					log.Infof("Create Tag Result [%v], params [%v]", resultGetCreateTags, params)
					go handshakeVdi(instanceInfo, WorkspaceString)
				}
			}
		}
	}

	c.JSON(resultCode, gin.H{
		"result": result,
	})
}

// deleteWorkspaces godoc
// @Summary 워크스페이스를 추가하는 API
// @Description 워크스페이를 추가하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags workspaces
// @Param workspaceUuid path string true "워크스페이스 UUID"
// @Router /api/v1/workspace [delete]
// @Success 200 {object} map[string]interface{}
func deleteWorkspaces(c *gin.Context) {
	returnData := map[string]interface{}{}
	resultCode := http.StatusNotFound
	workspaceUuid := c.Param("workspaceUuid")
	workspaceList, _ := selectWorkspaceList(workspaceUuid)
	_, errDeleteGroup := deleteGroup(workspaceList[0].Name)
	if errDeleteGroup != nil {
		log.Errorf("%v", errDeleteGroup)
	} else {
		resultDeleteWorkspace := deleteWorkspace(workspaceUuid)
		//deleteWorkspacePolicy(workspaceUuid)
		if resultDeleteWorkspace["status"] == http.StatusOK {
			returnData["message"] = "workspace delete success"
			resultCode = http.StatusOK
		}
	}

	c.JSON(resultCode, gin.H{
		"result": returnData,
	})
}

// postWorkspacesAgent godoc
// @Summary 워크스페이스를 추가하는 API
// @Description 워크스페이를 추가하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags workspacesAgent
// @Param uuid path string true "UUID"
// @Param asyncJobId path string true "async job UUID"
// @Param type path string true "workspace or instance"
// @Router /api/v1/workspaceAgent [POST]
// @Success 200 {object} map[string]interface{}
func postWorkspacesAgent(c *gin.Context) {
	instanceUuid := c.Param("instanceUuid")
	paramsUuid := c.PostForm("uuid")
	paramsType := c.PostForm("type")
	paramsLogin := c.PostForm("login")
	//paramsLogout := c.PostForm("logout")
	paramsHash := c.PostForm("hash")
	log.Debugf("instanceUuid [%v], paramsType [%v], paramsLogin [%v], paramsHash [%v]", instanceUuid, paramsType, paramsLogin, paramsHash)
	resultReturn := map[string]interface{}{}
	returnCode := http.StatusUnauthorized
	if paramsType == WorkspaceString {
		instanceList, _ := selectInstanceList(instanceUuid, InstanceString)
		if instanceList == nil {
			log.Errorf("Instance 조회결과가 없습니다.")
			returnCode = http.StatusNotFound
			resultReturn["message"] = "There are no instance search results."
		} else {
			//instanceInfo := instanceList[0]
			//workspaceTemplateCheck := updateWorkspaceTemplateCheck(instanceInfo.WorkspaceUuid, AgentOK)
			//
			//if workspaceTemplateCheck["status"] == http.StatusOK {
			//	//asyncJob := AsyncJob{}
			//	//asyncJob.Id = getUuid()
			//	//asyncJob.Name = VMDestroy
			//	//asyncJob.ExecUuid = instanceInfo.Uuid
			//	//asyncJob.Ready = 1
			//	//resultInsertAsyncJob := insertAsyncJob(asyncJob)
			//	//log.Infof("AsyncJob Insert Result [%v]", resultInsertAsyncJob)
			//	updateWorkspacePostfix(instanceInfo.WorkspaceUuid, 0)
			//	returnCode = http.StatusOK
			//}
		}

	} else if paramsType == InstanceString {
		instanceList, _ := selectInstanceList(instanceUuid, InstanceString)
		if instanceList == nil {
			log.Errorf("Instance 조회결과가 없습니다.")
			returnCode = http.StatusNotFound
			resultReturn["message"] = "There are no instance search results."
		} else {
			instanceInfo := instanceList[0]

			loginInfoMap := []map[string]string{}
			//logoutInfoMap := map[string]interface{}{}
			//layout := "2006/01/02 15:04:05"

			err1 := json.Unmarshal([]byte(paramsLogin), &loginInfoMap)
			if err1 != nil {
			}

			userCheck := false

			for i, v := range loginInfoMap {

				if v["id"] == instanceInfo.OwnerAccountId || v["id"] == instanceInfo.OwnerAccountId+"@"+os.Getenv("SambaDomain") {
					instanceCheck := updateInstanceCheck(paramsUuid, paramsLogin, paramsHash, 1)
					if instanceCheck["status"] == http.StatusOK {
						returnCode = http.StatusOK
						resultReturn["message"] = MessageAgentUpdateOK
					}
					userCheck = true
				}
				log.Errorf("i [%v], v [%v], instanceInfo [%v]", i, v, instanceInfo)

			}

			if !userCheck {
				instanceCheck := updateInstanceCheck(paramsUuid, paramsLogin, paramsHash, 0)
				if instanceCheck["status"] == http.StatusOK {
					returnCode = http.StatusOK
					resultReturn["message"] = MessageAgentUpdateOK
				}
			}

		}
	}

	c.JSON(returnCode, gin.H{
		"result": resultReturn,
	})
}

// getInstances godoc
// @Summary 워크스페이스의 instance 를 조회하는 API
// @Description 워크스페이스의 instance 를 조회하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags Instances
// @Param workspaceUuid path string true "Instance UUID"
// @Router /api/v1/instance/:instanceUuid [GET]
// @Success 200 {object} map[string]interface{}
func getInstances(c *gin.Context) {
	returnCode := http.StatusNotFound
	instanceUuid := c.Param("instanceUuid")
	instanceList, err := selectInstanceList(instanceUuid, WorkspaceString)
	returnData := map[string]interface{}{}

	log.Infof("instanceList = [%v], error = [%v]", instanceList, err)
	if err != nil {
		returnData["instanceInfo"] = err
	} else {
		returnData["instanceInfo"] = instanceList
		paramsInstanceList := []MoldParams{
			{"domainid": os.Getenv("MoldDomainId")},
		}
		virtualMachineList := getListVirtualMachinesMetrics(paramsInstanceList)
		listVirtualMachinesMetrics := ListVirtualMachinesMetrics{}
		virtualMachineInfo, _ := json.Marshal(virtualMachineList["listvirtualmachinesmetricsresponse"])
		err = json.Unmarshal([]byte(virtualMachineInfo), &listVirtualMachinesMetrics)
		for i, v := range instanceList {
			for _, v1 := range listVirtualMachinesMetrics.Virtualmachine {
				if v.MoldUuid == v1.Id {
					instanceList[i].MoldStatus = v1.State
					break
				}
			}
		}
		returnCode = http.StatusOK
	}
	c.JSON(returnCode, gin.H{
		"result": returnData,
	})
}

// getInstancesDetail godoc
// @Summary 워크스페이스의 instance 를 추가하는 API
// @Description 워크스페이스의 instance 를 추가하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags Instances
// @Param instanceUuid path string true "instance UUID"
// @Router /api/v1/instance/detail/:instanceUuid [GET]
// @Success 200 {object} map[string]interface{}
func getInstancesDetail(c *gin.Context) {
	returnCode := http.StatusOK
	instanceUuid := c.Param("instanceUuid")
	instanceList, _ := selectInstanceList(instanceUuid, InstanceString)
	instanceInfo := instanceList[0]
	log.Infof("instanceList [%v]", instanceList)
	returnData := map[string]interface{}{}

	paramsInstance := []MoldParams{
		{"id": instanceInfo.MoldUuid},
	}
	paramsVolume := []MoldParams{
		{"virtualmachineid": instanceInfo.MoldUuid},
	}
	virtualMachineList := getListVirtualMachinesMetrics(paramsInstance)
	virtualMachineVolumeList := getListVolumesMetrics(paramsVolume)

	listVirtualMachinesMetrics := ListVirtualMachinesMetrics{}
	virtualMachineInfo, _ := json.Marshal(virtualMachineList["listvirtualmachinesmetricsresponse"])
	json.Unmarshal([]byte(virtualMachineInfo), &listVirtualMachinesMetrics)
	instanceInfo.MoldStatus = listVirtualMachinesMetrics.Virtualmachine[0].State

	instanceInstanceVolumeInfo := InstanceInstanceVolumeInfo{}
	virtualMachineVolumeInfo, _ := json.Marshal(virtualMachineVolumeList["listvolumesmetricsresponse"])
	json.Unmarshal([]byte(virtualMachineVolumeInfo), &instanceInstanceVolumeInfo)

	returnData["instanceDBInfo"] = instanceInfo
	returnData["instanceMoldInfo"] = listVirtualMachinesMetrics
	returnData["instanceInstanceVolumeInfo"] = instanceInstanceVolumeInfo

	c.JSON(returnCode, gin.H{
		"result": returnData,
	})
}

// putInstances godoc
// @Summary 워크스페이스의 instance 를 추가하는 API
// @Description 워크스페이스의 instance 를 추가하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags Instances
// @Param uuid path string true "UUID"
// @Param quantity path string true "워크스페이스에 추가할 Instance 수량"
// @Router /api/v1/instance [PUT]
// @Success 200 {object} map[string]interface{}
func putInstances(c *gin.Context) {
	returnCode := http.StatusNotFound
	workspaceUuid := c.PostForm("uuid")
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	resultReturn := map[string]interface{}{}
	resultReturn["status"] = BaseErrorCode
	insertQuantity := 0
	log.WithFields(logrus.Fields{
		"workspaceController": "putInstances",
	}).Infof("uuid [%v], quantity [%v]", workspaceUuid, quantity)
	for i := 0; i < quantity; i++ {
		asyncJob := AsyncJob{}
		asyncJob.Id = getUuid()
		asyncJob.Name = VMsDeploy
		asyncJob.ExecUuid = workspaceUuid
		asyncJob.Ready = 1
		asyncJob.Parameter = "1"
		resultInsertAsyncJob := insertAsyncJob(asyncJob)
		if resultInsertAsyncJob["status"] == http.StatusOK {
			insertQuantity = insertQuantity + 1
			returnCode = http.StatusOK
		}
	}
	resultReturn["message"] = strconv.Itoa(quantity) + " virtual machines have been created and registered in async job."
	c.JSON(returnCode, gin.H{
		"result": resultReturn,
	})
}

// getConnectionRdp godoc
// @Summary instance 에 사용자를 할당하는 API
// @Description instance 에 사용자를 할당하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags Connection
// @Param instanceUuid path string true "Instance UUID"
// @Param userName path string true "Instance 에 할당할 userName"
// @Router /api/v1/connection/:instanceUuid/:username [GET]
// @Success 200 {object} map[string]interface{}
func getConnectionRdp(c *gin.Context) {
	returnCode := http.StatusNotFound
	instanceUuid := c.Param("instanceUuid")
	userName := c.Param("userName")
	instanceInfo, err := selectConnectionRdpDesktop(instanceUuid, userName)
	workspaceInfo := selectWorkspaceInfo(instanceInfo.WorkspaceUuid)
	//workspaceInfo := workspaceList[0]
	//var workspaceList []Workspace
	if err != nil {

	} else {
		returnCode = http.StatusOK
		userPassword, _ := selectUserPassword(userName)
		resp, _ := selectPasswordConvert(userPassword)

		var res map[string]string

		err = json.NewDecoder(resp.Body).Decode(&res)

		log.WithFields(logrus.Fields{
			"workspaceController": "getConnectionRdp",
		}).Infof("res [%v], userName [%v], userPassword [%v], res2 [%v]", res, userName, userPassword, strings.TrimSpace(res["stdout"]))

		var rdpPort string
		policyList := policyList(workspaceInfo.Uuid)
		for _, policyInfo := range policyList {
			if policyInfo.Name == "rdp_port" {
				rdpPort = policyInfo.Value
			}
		}

		instanceInfo.Password = userPassword
		instanceInfo.PublicPort = selectPublicPort(instanceInfo, workspaceInfo)
		instanceInfo.PrivatePort, _ = strconv.Atoi(rdpPort)
		updateRdpConnected(instanceInfo, 1)

		go checkRdpConnect(instanceInfo)
	}
	c.JSON(returnCode, gin.H{
		"instance": instanceInfo,
	})
}

// putConnection godoc
// @Summary instance 에 사용자를 할당하는 API
// @Description instance 에 사용자를 할당하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags Connection
// @Param instanceUuid path string true "Instance UUID"
// @Param username path string true "Instance 에 할당할 userName"
// @Router /api/v1/connection/:instanceUuid/:username [PUT]
// @Success 200 {object} map[string]interface{}
func putConnection(c *gin.Context) {
	//returnCode := http.StatusNotFound
	instanceUuid := c.Param("instanceUuid")
	userName := c.Param("username")
	resultReturn := map[string]interface{}{}
	log.WithFields(logrus.Fields{
		"workspaceController": "putConnection",
	}).Infof("instanceUuid [%v], userName [%v]", instanceUuid, userName)
	instanceList, _ := selectInstanceList(instanceUuid, InstanceString)
	instanceInfo := instanceList[0]
	paramsMold := []MoldParams{
		{"id": instanceInfo.MoldUuid},
	}
	resultMoldInstanceInfo := getListVirtualMachinesMetrics(paramsMold)
	resultUserInfo := selectUserDBDetail(userName)
	listVirtualMachinesMetrics := ListVirtualMachinesMetrics{}
	virtualMachineInfo, _ := json.Marshal(resultMoldInstanceInfo["listvirtualmachinesmetricsresponse"])
	json.Unmarshal([]byte(virtualMachineInfo), &listVirtualMachinesMetrics)
	parameter := "hostname=" + listVirtualMachinesMetrics.Virtualmachine[0].Nic[0].Ipaddress + ",port=" + os.Getenv("portForRDP") +
		",ignore-cert=true,username=" + resultUserInfo.UserName + ",password=" + resultUserInfo.Password + ",domain=" + os.Getenv("SambaDomain") + ",resize-method=display-update"
	//VDI 파라메터

	// APP 파라메터 추가필요
	//parameter = parameter + ",remote-app=C:\\Users\\dcmic\\AppData\\Local\\SourceTree\\SourceTree...,remote-app-dir=c:\\"
	//
	resultInstanceAllocatedUser := insertConnection(userName, instanceInfo.Name, parameter)
	log.Debugf("%v", resultInstanceAllocatedUser)
	log.Debugf("[%v]", resultInstanceAllocatedUser.Status)
	updateInstanceUser(instanceInfo.Uuid, resultUserInfo.UserName)

	c.JSON(http.StatusOK, gin.H{
		"result": resultReturn,
	})
}

// putConnection godoc
// @Summary instance 에 사용자를 할당하는 API
// @Description instance 에 사용자를 할당하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags Connection
// @Param instanceUuid path string true "Instance UUID"
// @Param username path string true "Instance 에 할당할 userName"
// @Router /api/v1/appConnection/:instanceUuid/:username [PUT]
// @Success 200 {object} map[string]interface{}
func putAppConnection(c *gin.Context) {
	//returnCode := http.StatusNotFound
	instanceUuid := c.Param("instanceUuid")
	userName := c.Param("username")
	connectionName := c.Param("connection")
	resultReturn := map[string]interface{}{}
	log.WithFields(logrus.Fields{
		"workspaceController": "putAppConnection",
	}).Infof("instanceUuid [%v], userName [%v]", instanceUuid, userName)
	instanceList, _ := selectInstanceList(instanceUuid, InstanceString)
	instanceInfo := instanceList[0]
	paramsMold := []MoldParams{
		{"id": instanceInfo.MoldUuid},
	}
	resultMoldInstanceInfo := getListVirtualMachinesMetrics(paramsMold)
	resultUserInfo := selectUserDBDetail(userName)
	listVirtualMachinesMetrics := ListVirtualMachinesMetrics{}
	virtualMachineInfo, _ := json.Marshal(resultMoldInstanceInfo["listvirtualmachinesmetricsresponse"])
	json.Unmarshal([]byte(virtualMachineInfo), &listVirtualMachinesMetrics)
	parameter := "hostname=" + listVirtualMachinesMetrics.Virtualmachine[0].Nic[0].Ipaddress + ",port=" + os.Getenv("portForRDP") +
		",ignore-cert=true,username=" + resultUserInfo.UserName + ",password=" + resultUserInfo.Password + ",domain=" + os.Getenv("SambaDomain") + ",resize-method=display-update" +
		",remote-app=C:\\NewGen\\Rebirth\\Rebirth.exe,remote-app-dir=c:\\NewGen\\Rebirth\\"
	//",remote-app=C:\\Program Files\\Microsoft Office\\root\\Office16\\EXCEL.EXE"
	//VDI 파라메터

	// APP 파라메터 추가필요
	//parameter = parameter + ",remote-app=C:\\Users\\dcmic\\AppData\\Local\\SourceTree\\SourceTree...,remote-app-dir=c:\\"
	//
	resultInstanceAllocatedUser := insertConnection(userName, connectionName, parameter)
	log.Debugf("%v", resultInstanceAllocatedUser)
	log.Debugf("[%v]", resultInstanceAllocatedUser.Status)
	updateInstanceUser(instanceInfo.Uuid, resultUserInfo.UserName)

	c.JSON(http.StatusOK, gin.H{
		"result": resultReturn,
	})
}

// deleteConnection godoc
// @Summary instance 에 사용자를 할당하는 API
// @Description instance 에 사용자를 할당하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags Connection
// @Param instanceUuid path string true "Connection 을 삭제할 instance Uuid"
// @Router /api/v1/connection/:instanceUuid [DELETE]
// @Success 200 {object} map[string]interface{}
func deleteConnection(c *gin.Context) {
	returnCode := http.StatusNotFound
	instanceUuid := c.Param("instanceUuid")
	//userName := c.Param("username")
	resultReturn := map[string]interface{}{}
	log.WithFields(logrus.Fields{
		"workspaceController": "deleteConnection",
	}).Infof("instanceUuid [%v]", instanceUuid)
	instanceList, _ := selectInstanceList(instanceUuid, InstanceString)
	instanceInfo := instanceList[0]

	resultDelConnection := delConnection(instanceInfo.Name)
	log.WithFields(logrus.Fields{
		"workspaceController": "deleteConnection",
	}).Debugf("resultDelConnection [%v]", resultDelConnection)
	//log.Debugf("%v", resultDelConnection)
	if resultDelConnection.StatusCode == http.StatusCreated {
		updateInstanceUser(instanceInfo.Uuid, "")
		returnCode = http.StatusNoContent
	}

	c.JSON(returnCode, gin.H{
		"result": resultReturn,
	})
}

// patchInstances godoc
// @Summary instance 의 상태 변경하는 API
// @Description instance 의 상태를 변경하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags Instances
// @Param action path string true "action 해당 값은 [VMStart, VMStop, VMDestroy, VMReboot] 으로 보내야 합니다."
// @Param instanceUuid path string true "Instance UUID"
// @Router /api/v1/instance/:action/:instanceUuid [PATCH]
// @Success 200 {object} map[string]interface{}
func patchInstances(c *gin.Context) {
	returnCode := http.StatusUnauthorized
	result := map[string]interface{}{}
	action := c.Param("action")
	instanceUuid := c.Param("instanceUuid")
	log.Debugf("action [%v], instanceUuid [%v]", action, instanceUuid)
	asyncJob := AsyncJob{}
	asyncJob.Id = getUuid()
	asyncJob.Name = action
	asyncJob.ExecUuid = instanceUuid
	asyncJob.Ready = 1
	resultInsertAsyncJob := insertAsyncJob(asyncJob)
	if resultInsertAsyncJob["status"] == http.StatusOK {
		returnCode = http.StatusOK
		result["asyncJobId"] = resultInsertAsyncJob["jobid"]
		result["message"] = resultInsertAsyncJob["message"]
	}

	c.JSON(returnCode, gin.H{
		"result": result,
	})
}

// patchHandshake godoc
// @Summary instance 의 handshake 를 재실행 하는 API
// @Description instance 의 handshake 를 재실행 하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags Handshake
// @Param instanceUuid path string true "Instance UUID"
// @Param instanceType path string true "Instance Type InstanceString OR WorkspaceString"
// @Router /api/v1/handshake/:instanceUuid/:instanceType [PATCH]
// @Success 200 {object} map[string]interface{}
func patchHandshake(c *gin.Context) {
	result := map[string]interface{}{}
	instanceUuid := c.Param("instanceUuid")
	instanceType := c.Param("instanceType")
	log.Debugf("instanceUuid [%v]", instanceUuid)
	instanceList, _ := selectInstanceList(instanceUuid, InstanceString)
	instanceInfo := instanceList[0]
	go handshakeVdi(instanceInfo, instanceType)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// getDashboard godoc
// @Summary dashboard 조회하는 API
// @Description 워크스페이스 수, 데스크톱 수, 데스크톱 연결 수, APP 연결 수 정보를 제공하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags Dashboard
// @Router /api/v1/dashboard [get]
// @Success 200 {object} map[string]interface{}
func getDashboard(c *gin.Context) {

	resultData := map[string]interface{}{}
	resultCode := http.StatusNotFound
	returnCountWorkspace, workspaceErr := selectCountWorkspace()
	returnCountInstance, instanceErr := selectCountInstance()
	returnCountConnected, connectedErr := selectCountDesktopConnected()
	userList, userListErr := getUserList()
	log.WithFields(logrus.Fields{
		"workspacesController.go": "getDashboard",
	}).Infof("clientIP [%v]", c.ClientIP())

	if workspaceErr == nil && instanceErr == nil && connectedErr == nil && userListErr == nil {
		resultCode = http.StatusOK
		resultData["workspaceCount"] = returnCountWorkspace
		resultData["instanceCount"] = returnCountInstance
		resultData["connectedCount"] = returnCountConnected
		resultData["usersCount"] = len(userList)
	}
	c.JSON(resultCode, gin.H{
		"result": resultData,
	})
}
