package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
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
	result := selectWorkspaceList()

	var resultData []Workspace
	err := json.Unmarshal([]byte(result), &resultData)
	if err != nil {

	}
	returnData["list"] = resultData
	returnData["listTotalCount"] = selectCountWorkspace()
	returnData["status"] = http.StatusOK

	c.JSON(http.StatusOK, gin.H{
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
	result := selectWorkspaceDetail(workspaceUuid)

	paramsTemplate := []MoldParams{
		{"templatefilter": "executable"},
		{"id": result.TemplateUuid},
	}

	templateResult := getTemplate(paramsTemplate)

	paramsServiceOffering := []MoldParams{
		{"id": result.ComputeOfferingUuid},
	}

	serviceOfferingResult := getComputeOffering(paramsServiceOffering)

	paramsNetwork := []MoldParams{
		{"id": result.NetworkUuid},
	}
	networkResult := getNetwork(paramsNetwork)

	resultReturn["workspaceInfo"] = result
	resultReturn["templateInfo"] = templateResult
	resultReturn["serviceOfferingInfo"] = serviceOfferingResult
	resultReturn["networkInfo"] = networkResult
	resultReturn["status"] = http.StatusOK

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

	workspace.Uuid = getUuid()
	workspace.Name = c.PostForm("name")
	workspace.Description = c.PostForm("description")
	workspace.Type = c.PostForm("type")
	workspace.TemplateUuid = c.PostForm("templateUuid")
	workspace.ComputeOfferingUuid = c.PostForm("computeOfferingUuid")
	workspace.Shared, _ = strconv.ParseBool(c.PostForm("shared"))
	workspace.NetworkUuid = selectNetworkDetail()

	result := insertWorkspace(workspace)
	//params := []MoldParams{
	//	{"serviceofferingid": workspace.ComputeOfferingUuid},
	//	{"templateid": workspace.TemplateUuid},
	//	{"zoneid": selectZoneId()},
	//}

	resultDeploy := getDeployVirtualMachine(workspace.Uuid)
	log.Infof("Mold 통신 결과값 [%v]\n", resultDeploy)
	if resultDeploy["deployvirtualmachineresponse"].(map[string]interface{})["errorcode"] != nil {
		resultDeploy["status"] = SignatureErrorCode
		c.JSON(http.StatusOK, gin.H{
			"result": resultDeploy,
		})
	} else {
		asyncJob := AsyncJob{}
		asyncJob.Name = VMDestroy
		asyncJob.ExecUuid = resultDeploy["deployvirtualmachineresponse"].(map[string]interface{})["id"].(string)
		asyncJob.Ready = 0
		asyncJob.Parameter = ""
		resultAsyncJob := insertAsyncJob(asyncJob)

		c.JSON(http.StatusOK, gin.H{
			"result":         result,
			"resultDeploy":   resultDeploy,
			"resultAsyncJob": resultAsyncJob,
		})
	}
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
	paramsJobId := c.PostForm("asyncJobId")
	paramsType := c.PostForm("type")
	resultReturn := map[string]interface{}{}

	resultUpdateWorkspaceAgent := updateWorkspaceAgent(paramsUuid, paramsJobId, paramsType)
	if resultUpdateWorkspaceAgent["status"] == http.StatusOK {
		resultDeleteAsyncJob := deleteAsyncJob(paramsJobId)
		if resultDeleteAsyncJob["status"] == http.StatusOK {
			resultReturn["asyncDeleteResult"] = resultDeleteAsyncJob
			if paramsType == InstanceString {
				resultInstance := selectInstanceDetail(paramsUuid)
				resultWorkspace := selectWorkspaceDetail(resultInstance.WorkspaceUuid)
				params := []MoldParams{
					{"resourceids": paramsUuid},
					{"tags[0].key": ServiceDaaS},
					{"tags[0].value": ServiceWorks},
					{"tags[1].key": WorkspaceName},
					{"tags[1].value": resultWorkspace.Name},
				}
				resultGetCreateTags := getCreateTags(params)
				resultReturn["resultGetCreateTags"] = resultGetCreateTags
			}
		}
		resultReturn["asyncUpdateResult"] = resultUpdateWorkspaceAgent
	}

	c.JSON(http.StatusOK, gin.H{
		"result": resultReturn,
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
	uuid := c.PostForm("uuid")
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	resultReturn := map[string]interface{}{}
	resultReturn["status"] = BaseErrorCode
	insertQuantity := 0
	log.WithFields(logrus.Fields{
		"workspaceController": "putInstances",
	}).Infof("uuid=[%v], quantity=[%v]", uuid, quantity)
	asyncJob := AsyncJob{}
	asyncJob.Id = getUuid()
	asyncJob.Name = VMsDeploy
	asyncJob.ExecUuid = uuid
	asyncJob.Ready = 1
	asyncJob.Parameter = strconv.Itoa(quantity)
	resultInsertAsyncJob := insertAsyncJob(asyncJob)
	log.Info("qweqweqweqweqweqweqweqweqweqwewqeqweqweqweqwe")
	log.Info(resultInsertAsyncJob)
	log.Info("qweqweqweqweqweqweqweqweqweqwewqeqweqweqweqwe")
	if resultInsertAsyncJob["status"] == http.StatusOK {
		insertQuantity = insertQuantity + 1
		resultReturn["status"] = http.StatusOK
	}
	resultReturn["message"] = strconv.Itoa(quantity) + " virtual machines have been created and registered in async job."
	c.JSON(http.StatusOK, gin.H{
		"result": resultReturn,
	})
}
