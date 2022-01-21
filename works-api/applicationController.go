package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// getApplicationServer godoc
// v1.GET("/applicationServer/:workspaceUuid", getApplication)
// @Summary Application List  조회하는 API
// @Description Application List 를 조회하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags Application
// @Router /api/v1/applicationServer/:workspaceUuid [GET]
// @Success 200 {object} map[string]interface{}
func getApplicationServer(c *gin.Context) {
	//resultData := map[string]interface{}{}
	resultCode := http.StatusNotFound
	workspaceUuid := c.Param("workspaceUuid")
	log.WithFields(logrus.Fields{
		"applicationController.go": "getApplication",
	}).Infof("workspaceUuid [%v]", workspaceUuid)

	applicationList := selectApplicationServerListService(workspaceUuid)
	instanceList, _ := selectInstanceList(workspaceUuid, WorkspaceString)
	c.JSON(resultCode, gin.H{
		"applicationList": applicationList,
		"instanceList":    instanceList,
	})
}

// getApplication godoc
// v1.GET("/applicationServer/:workspaceUuid", getApplication)
// @Summary Application List  조회하는 API
// @Description Application List 를 조회하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags Application
// @Router /api/v1/application/:workspaceUuid [GET]
// @Success 200 {object} map[string]interface{}
func getApplication(c *gin.Context) {
	//resultData := map[string]interface{}{}
	resultCode := http.StatusNotFound
	workspaceUuid := c.Param("workspaceUuid")
	log.WithFields(logrus.Fields{
		"applicationController.go": "getApplication",
	}).Infof("workspaceUuid [%v]", workspaceUuid)

	applicationList := selectApplicationListService(workspaceUuid)
	c.JSON(resultCode, gin.H{
		"applicationList": applicationList,
	})
}

// postApplication godoc
// v1.POST("/application", postApplication)
// @Summary Application 가상화를 등록하는 API
// @Description 워크스페이스 수, 데스크톱 수, 데스크톱 연결 수, APP 연결 수 정보를 제공하는 API 입니다.
// @Param workspaceUuid path string true "Workspace 의 Uuid"
// @Param name path string true "Application 이름"
// @Param path path string true "Application 경로"
// @Param workingDirectory path string true "Application workingDirectory"
// @Param RegApplicationServer path string true "Application RegApplicationServer"
// @Accept  json
// @Produce  json
// @Tags Application
// @Router /api/v1/application [POST]
// @Success 200 {object} map[string]interface{}
func postApplication(c *gin.Context) {
	resultData := map[string]interface{}{}
	resultCode := http.StatusNotFound
	application := Application{
		Uuid:                 getUuid(),
		WorkspaceUuid:        c.PostForm("workspaceUuid"),
		Name:                 c.PostForm("name"),
		Path:                 c.PostForm("path"),
		WorkingDirectory:     c.PostForm("workingDirectory"),
		RegApplicationServer: c.PostForm("regApplicationServer"),
	}
	log.WithFields(logrus.Fields{
		"applicationController.go": "postApplication",
	}).Infof("application [%v]", application)

	resultInsertApplication := insertApplicationService(application)

	if resultInsertApplication["status"] == http.StatusOK {
		resultData["message"] = resultInsertApplication["message"]
		resultCode = http.StatusOK
	}

	c.JSON(resultCode, gin.H{
		"result": resultData,
	})
}

// deleteApplication godoc
// v1.DELETE("/application/:applicationUuid", deleteApplication)
// @Summary Application 가상화를 등록하는 API
// @Description 워크스페이스 수, 데스크톱 수, 데스크톱 연결 수, APP 연결 수 정보를 제공하는 API 입니다.
// @Param name path string true "Application 이름"
// @Param path path string true "Application 경로"
// @Param workingDirectory path string true "Application workingDirectory"
// @Param RegApplicationServer path string true "Application RegApplicationServer"
// @Accept  json
// @Produce  json
// @Tags Application
// @Router /api/v1/application [DELETE]
// @Success 200 {object} map[string]interface{}
func deleteApplication(c *gin.Context) {
	resultData := map[string]interface{}{}
	resultCode := http.StatusNotFound
	application := Application{
		Uuid: c.Param("applicationUuid"),
	}
	log.WithFields(logrus.Fields{
		"applicationController.go": "deleteApplication",
	}).Infof("application [%v]", application)

	resultInsertApplication := insertApplicationService(application)

	if resultInsertApplication["status"] == http.StatusOK {
		resultData["message"] = resultInsertApplication["message"]
		resultCode = http.StatusOK
	}

	c.JSON(resultCode, gin.H{
		"appList": resultData,
	})
}
