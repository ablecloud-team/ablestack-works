package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// getUserDesktop godoc
// @Summary 유저 Desktop 목록을 조회하는 API
// @Description 유저에게 할당된 Desktop 목록을 조회하는 API
// @Tags userDesktop
// @Accept  json
// @Produce  json
// @Router /api/v1/userdesktop/:userName [GET]
// @Success 200 {object} map[string]interface{}
func getUserDesktop(c *gin.Context) {
	returnCode := http.StatusNotFound
	userName := c.Param("userName")
	workspaceUuidList, err := selectUserDesktopWorkspaceList(userName)
	var workspaceList []Workspace
	if err != nil {

	} else {
		returnCode = http.StatusOK
		for _, workspaceUuid := range workspaceUuidList {
			userPassword, _ := selectUserPassword(userName)
			workspaceInfo, _ := selectWorkspaceList(workspaceUuid)
			instanceList, _ := selectUserDesktopListForWorkspace(workspaceInfo[0].Uuid, userName)
			for i, _ := range instanceList {
				instanceList[i].Password = userPassword
			}
			workspaceInfo[0].InstanceList = append(workspaceInfo[0].InstanceList, instanceList...)
			workspaceList = append(workspaceList, workspaceInfo[0])
		}
	}
	c.JSON(returnCode, gin.H{
		"workspaceList": workspaceList,
	})
}
