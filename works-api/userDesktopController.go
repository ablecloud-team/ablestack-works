package main

import (
	"encoding/json"
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
			workspacePolicy, _ := selectWorkspacePolicyList(workspaceUuid)
			workspaceInfo[0].Policy.RdpPort = workspacePolicy[0].Policy.RdpPort
			workspaceInfo[0].Policy.RdpAccessAllow = workspacePolicy[0].Policy.RdpAccessAllow
			instanceList, _ := selectUserDesktopListForWorkspace(workspaceInfo[0].Uuid, userName)
			for _, instanceInfo := range instanceList {
				instanceInfo.Password = userPassword
				paramsMold := []MoldParams{
					{"id": instanceInfo.MoldUuid},
				}
				resultMoldInstanceInfo := getListVirtualMachinesMetrics(paramsMold)
				listVirtualMachinesMetrics := ListVirtualMachinesMetrics{}
				virtualMachineInfo, _ := json.Marshal(resultMoldInstanceInfo["listvirtualmachinesmetricsresponse"])
				json.Unmarshal([]byte(virtualMachineInfo), &listVirtualMachinesMetrics)
				instanceInfo.MoldStatus = listVirtualMachinesMetrics.Virtualmachine[0].State
				workspaceInfo[0].InstanceList = append(workspaceInfo[0].InstanceList, instanceInfo)
			}
			workspaceList = append(workspaceList, workspaceInfo[0])
		}
	}
	c.JSON(returnCode, gin.H{
		"workspaceList": workspaceList,
	})
}

// getClient godoc
// @Summary 유저 Desktop 목록을 조회하는 API
// @Description 유저에게 할당된 Desktop 목록을 조회하는 API
// @Tags userDesktop
// @Accept  json
// @Produce  json
// @Router /api/client [GET]
// @Success 200 {object} map[string]interface{}postClient
func getClient(c *gin.Context) {
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
			for _, instanceInfo := range instanceList {
				instanceInfo.Password = userPassword
				paramsMold := []MoldParams{
					{"id": instanceInfo.MoldUuid},
				}
				resultMoldInstanceInfo := getListVirtualMachinesMetrics(paramsMold)
				listVirtualMachinesMetrics := ListVirtualMachinesMetrics{}
				virtualMachineInfo, _ := json.Marshal(resultMoldInstanceInfo["listvirtualmachinesmetricsresponse"])
				json.Unmarshal([]byte(virtualMachineInfo), &listVirtualMachinesMetrics)
				instanceInfo.MoldStatus = listVirtualMachinesMetrics.Virtualmachine[0].State
				workspaceInfo[0].InstanceList = append(workspaceInfo[0].InstanceList, instanceInfo)
			}
			workspaceList = append(workspaceList, workspaceInfo[0])
		}
	}
	c.JSON(returnCode, gin.H{
		"workspaceList": workspaceList,
	})
}

// postClient godoc
// @Summary 유저 Desktop 목록을 조회하는 API
// @Description 유저에게 할당된 Desktop 목록을 조회하는 API
// @Tags userDesktop
// @Accept  json
// @Produce  json
// @Router /api/v1/userdesktop/:userName [GET]
// @Success 200 {object} map[string]interface{}deleteClient
func postClient(c *gin.Context) {
	returnCode := http.StatusOK
	userName := c.PostForm("userName")
	instanceUuid := c.PostForm("instanceUuid")

	instanceInfo, _ := selectConnectionRdpDesktop(instanceUuid, userName)

	updateRdpConnected(instanceInfo, 0)

	c.JSON(returnCode, gin.H{
		"userName":     userName,
		"instanceUuid": instanceUuid,
	})
}

// deleteClient godoc
// @Summary 유저 Desktop 목록을 조회하는 API
// @Description 유저에게 할당된 Desktop 목록을 조회하는 API
// @Tags userDesktop
// @Accept  json
// @Produce  json
// @Router /api/v1/userdesktop/:userName [GET]
// @Success 200 {object} map[string]interface{}
func deleteClient(c *gin.Context) {
	returnCode := http.StatusOK
	userName := c.Param("userName")
	instanceUuid := c.Param("instanceUuid")

	instanceInfo, _ := selectConnectionRdpDesktop(instanceUuid, userName)

	deletePortForwading(instanceInfo)
	updatePortForwardingInstanceUuidDelete(instanceInfo.Uuid)
	updateRdpConnected(instanceInfo, 0)

	c.JSON(returnCode, gin.H{
		"userName":     userName,
		"instanceUuid": instanceUuid,
	})
}
