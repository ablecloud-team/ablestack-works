package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
)

// getWorkspaces godoc
// @Summary 워크스페이스 리스트를 조회하는 API
// @Description 워크 스페이스 리스트를 조회하는 API 입니다.
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
// @Accept  json
// @Produce  json
// @Param workspaceUuid path string true "워크스페이스 UUID"
// @Router /api/v1/workspace/:workspaceUuid [get]
// @Success 200 {object} map[string]interface{}
func getWorkspacesDetail(c *gin.Context) {
	workspaceUuid := c.Param("workspaceUuid")
	resultReturn := map[string]interface{}{}
	workspaceList, _ := selectWorkspaceList(workspaceUuid)
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

	groupDetail, _ := selectGroupDetail(workspaceInfo.Name)
	var groupData map[string]interface{}
	json.NewDecoder(groupDetail.Body).Decode(&groupData)

	resultReturn["workspaceInfo"] = workspaceInfo
	resultReturn["templateInfo"] = templateResult["listtemplatesresponse"]
	resultReturn["serviceOfferingInfo"] = serviceOfferingResult["listserviceofferingsresponse"]
	resultReturn["networkInfo"] = networkResult["listnetworksresponse"]
	resultReturn["instanceList"] = instanceList
	resultReturn["groupDetail"] = groupData
	//resultReturn["listVirtualMachinesMetrics"] = listVirtualMachinesMetrics

	c.JSON(http.StatusOK, gin.H{
		"result": resultReturn,
	})
}

// getOffering godoc
// @Summary 템플릿 및 각종 오퍼링을 조회하는 API
// @Description 템플릿, 서비스 오퍼링, 디스크 오퍼링 정보를 Mold 에서 조회하여 제공하는 API 입니다.
// @Accept  json
// @Produce  json
// @Router /api/v1/offering [get]
// @Success 200 {object} map[string]interface{}
func getOffering(c *gin.Context) {
	var params []MoldParams
	paramsTemplate := []MoldParams{
		{"templatefilter": "all"},
	}
	result := map[string]interface{}{
		"status": http.StatusOK,
	}
	templateResult := getTemplate(paramsTemplate)
	serviceOfferingResult := getComputeOffering(params)
	//networkResult := getNetwork(params)
	//diskOfferingResult := getDiskOffering(params)
	c.JSON(http.StatusOK, gin.H{
		"result":              result,
		"templateList":        templateResult,
		"serviceOfferingList": serviceOfferingResult,
		//"networkList":         networkResult,
		//"diskOfferingList":    diskOfferingResult,
	})
}

// putWorkspaces godoc
// @Summary 워크스페이스를 추가하는 API
// @Description 워크스페이를 추가하는 API 입니다.
// @Accept  json
// @Produce  json
// @Param name path string true "워크스페이스 이름"
// @Param description path string true "워크스페이스 설명"
// @Param type path string true "워크스페이스 타입(Desktop or Application)"
// @Param computeOfferingUuid path string true "워크스페이스에서 사용할 Compute offering UUID"
// @Param templateUuid path string true "워크스페이스에서 사용할 Template UUID"
// @Param shared path bool true "워크스페이스에서 Shard 여부 전용이면 'false', 공용이면 'true'"
// @Router /api/v1/workspace [put]
// @Success 200 {object} map[string]interface{}
func putWorkspaces(c *gin.Context) {
	workspace := Workspace{}
	result := map[string]interface{}{}
	resultCode := http.StatusNotFound
	workspace.Uuid = getUuid()
	workspace.Name = c.PostForm("name")
	workspace.Description = c.PostForm("description")
	workspace.WorkspaceType = c.PostForm("type")
	workspace.TemplateUuid = c.PostForm("templateUuid")
	workspace.ComputeOfferingUuid = c.PostForm("computeOfferingUuid")
	workspace.Shared, _ = strconv.ParseBool(c.PostForm("shared"))
	workspace.NetworkUuid = selectNetworkDetail()
	resultInsertGroup, err := insertGroup(workspace.Name)
	if err != nil {
		log.Errorf("워크스페이스 그룹 생성 DC API 통신중 에러가 발생했습니다.[%v]", err)
	}
	res := map[string]interface{}{}
	json.NewDecoder(resultInsertGroup.Body).Decode(&res)
	result["resultInsertGroup"] = res
	if resultInsertGroup.Status == Created201 {
		resultInsertWorkspace, _ := insertWorkspace(workspace)
		log.Info(resultInsertWorkspace)
		result["insertWorkspace"] = resultInsertWorkspace
		if resultInsertWorkspace["status"] == http.StatusOK {
			instanceUuid := getUuid()
			resultDeploy := getDeployVirtualMachine(workspace.Uuid, instanceUuid, WorkspaceString)
			log.Infof("Mold 통신 결과값 [%v]\n", resultDeploy)
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
				json.Unmarshal([]byte(virtualMachineInfo), &listVirtualMachinesMetrics)

				workspaceList, _ := selectWorkspaceList(workspace.Uuid)
				workspaceInfo := workspaceList[0]
				instance := Instance{}
				instance.Uuid = instanceUuid
				instance.MoldUuid = resultDeploy["deployvirtualmachineresponse"].(map[string]interface{})["id"].(string)
				instance.Name = listVirtualMachinesMetrics.Virtualmachine[0].Displayname
				instance.WorkspaceUuid = workspaceInfo.Uuid
				instance.WorkspaceName = workspaceInfo.Name
				resultInsertInstance := insertInstance(instance)
				if resultInsertInstance["status"] == http.StatusOK {
					resultCode = http.StatusOK
					result["resultInsertDeploy"] = resultInsertInstance
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
// @Param workspaceUuid path string true "워크스페이스 UUID"
// @Router /api/v1/workspace [delete]
// @Success 200 {object} map[string]interface{}
func deleteWorkspaces(c *gin.Context) {
	//result := map[string]interface{}{}
	resultCode := http.StatusNotFound
	//workspaceUuid = c.PostForm("workspaceUuid")
	//res := map[string]interface{}{}

	c.JSON(resultCode, gin.H{
		"result": "res",
	})
}

// putWorkspacesAgent godoc
// @Summary 워크스페이스를 추가하는 API
// @Description 워크스페이를 추가하는 API 입니다.
// @Accept  json
// @Produce  json
// @Param uuid path string true "UUID"
// @Param asyncJobId path string true "async job UUID"
// @Param type path string true "workspace or instance"
// @Router /api/v1/workspaceAgent [POST]
// @Success 200 {object} map[string]interface{}
func putWorkspacesAgent(c *gin.Context) {
	paramsUuid := c.PostForm("uuid")
	paramsType := c.PostForm("type")
	paramsLogin := c.PostForm("login")
	paramsLogout := c.PostForm("logout")
	log.Debugf("paramsUuid [%v], paramsType [%v], paramsLogin [%v], paramsLogout [%v]", paramsUuid, paramsType, paramsLogin, paramsLogout)
	resultReturn := map[string]interface{}{}
	returnCode := http.StatusUnauthorized
	if paramsType == WorkspaceString {
		instanceList, _ := selectInstanceList(paramsUuid, InstanceString)
		instanceInfo := instanceList[0]
		workspaceTemplateCheck := updateWorkspaceTemplateCheck(instanceInfo.WorkspaceUuid)

		if workspaceTemplateCheck["status"] == http.StatusOK {
			asyncJob := AsyncJob{}
			asyncJob.Id = getUuid()
			asyncJob.Name = VMDestroy
			asyncJob.ExecUuid = instanceInfo.Uuid
			asyncJob.Ready = 1
			resultInsertAsyncJob := insertAsyncJob(asyncJob)
			log.Infof("AsyncJob Insert Result [%v]", resultInsertAsyncJob)
			updateWorkspacePostfix(instanceInfo.WorkspaceUuid, 0)
			returnCode = http.StatusOK
		}
	} else if paramsType == InstanceString {
		instanceCheck := updateInstanceCheck(paramsUuid, paramsLogin, paramsLogout)
		if instanceCheck["status"] == http.StatusOK {
			returnCode = http.StatusOK
			resultReturn["message"] = MessageAgentUpdateOK
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
// @Param workspaceUuid path string true "Instance UUID"
// @Router /api/v1/instance/detail/:instanceUuid [GET]
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
		json.Unmarshal([]byte(virtualMachineInfo), &listVirtualMachinesMetrics)
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
	resultMoldInfo := getListVirtualMachinesMetrics(paramsInstance)
	resultInstanceVolumeInfo := getlistVolumesMetrics(paramsVolume)
	returnData["instanceDBInfo"] = instanceInfo
	returnData["instanceMoldInfo"] = resultMoldInfo
	returnData["instanceInstanceVolumeInfo"] = resultInstanceVolumeInfo
	c.JSON(returnCode, gin.H{
		"result": returnData,
	})
}

// putInstances godoc
// @Summary 워크스페이스의 instance 를 추가하는 API
// @Description 워크스페이스의 instance 를 추가하는 API 입니다.
// @Accept  json
// @Produce  json
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
	asyncJob := AsyncJob{}
	asyncJob.Id = getUuid()
	asyncJob.Name = VMsDeploy
	asyncJob.ExecUuid = workspaceUuid
	asyncJob.Ready = 1
	asyncJob.Parameter = strconv.Itoa(quantity)
	resultInsertAsyncJob := insertAsyncJob(asyncJob)
	if resultInsertAsyncJob["status"] == http.StatusOK {
		insertQuantity = insertQuantity + 1
		returnCode = http.StatusOK
	}
	resultReturn["message"] = strconv.Itoa(quantity) + " virtual machines have been created and registered in async job."
	c.JSON(returnCode, gin.H{
		"result": resultReturn,
	})
}

// postInstances godoc
// @Summary instance 에 사용자를 할당하는 API
// @Description instance 에 사용자를 할당하는 API 입니다.
// @Accept  json
// @Produce  json
// @Param instanceUuid path string true "Instance UUID"
// @Param username path string true "Instance 에 할당할 userName"
// @Router /api/v1/instance [POST]
// @Success 200 {object} map[string]interface{}
func postInstances(c *gin.Context) {
	//returnCode := http.StatusNotFound
	instanceUuid := c.PostForm("instanceUuid")
	username := c.PostForm("username")
	resultReturn := map[string]interface{}{}
	log.WithFields(logrus.Fields{
		"workspaceController": "postInstances",
	}).Infof("uuid=[%v], username=[%v]", instanceUuid, username)
	instanceList, _ := selectInstanceList(instanceUuid, InstanceString)
	instanceInfo := instanceList[0]
	paramsMold := []MoldParams{
		{"id": instanceInfo.MoldUuid},
	}
	resultMoldInstanceInfo := getListVirtualMachinesMetrics(paramsMold)
	resultUserInfo := selectUserDBDetail(username)
	listVirtualMachinesMetrics := ListVirtualMachinesMetrics{}
	virtualMachineInfo, _ := json.Marshal(resultMoldInstanceInfo["listvirtualmachinesmetricsresponse"])
	json.Unmarshal([]byte(virtualMachineInfo), &listVirtualMachinesMetrics)
	parameter := "hostname=" + listVirtualMachinesMetrics.Virtualmachine[0].Nic[0].Ipaddress + ",port=3389,ignore-cert=true,username=" + resultUserInfo.UserName + ",password=" + resultUserInfo.Password + ",domain=" + os.Getenv("SambaDomain")
	resultUserAllocatedInstance := insertUserAllocatedInstance(username, instanceInfo.Name, parameter)
	log.Debugf("%v", resultUserAllocatedInstance)
	log.Debugf("[%v]", resultUserAllocatedInstance.Status)
	updateInstanceUser(instanceInfo.Uuid, resultUserInfo.UserName)
	//if strings.TrimSpace(resultGuacamole.Status) == "200"{
	//log.Debugf("[%v]","참참참참")
	//	resultDB := UpdateInstanceUser(uuid, username)
	//	if resultDB["status"] == http.StatusOK {
	//		returnCode = http.StatusOK
	//	}
	//}

	c.JSON(http.StatusOK, gin.H{
		"result": resultReturn,
	})
}

// patchInstances godoc
// @Summary instance 의 상태 변경하는 API
// @Description instance 의 상태를 변경하는 API 입니다.
// @Accept  json
// @Produce  json
// @Param action path string true "action 해당 값은 [VMStart, VMStop, VMDestroy] 으로 보내야 합니다."
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
