package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// patchPolicy godoc
// @Summary 정책을 변경하는 API
// @Description 정책을 변경하는 API 입니다.
// @Tags policy
// @Param policyName path string true "정책명"
// @Param policyValue path string true "정책값"
// @Accept  json
// @Produce  json
// @Router /api/v1/policy/:workspaceUuid [patch]
// @Success 200 {object} map[string]interface{}
func patchPolicy(c *gin.Context) {
	policyInfo := Policy{}
	policyInfo.WorkspacesUuid = c.Param("workspaceUuid")
	policyInfo.Name = c.PostForm("policyName")
	policyInfo.Value = c.PostForm("policyValue")
	workspaceInfo := selectWorkspaceInfo(policyInfo.WorkspacesUuid)
	returnCode := http.StatusOK
	returnData := map[string]interface{}{}

	returnData["patchPolicy"] = updatePolicy(workspaceInfo, policyInfo)
	c.JSON(returnCode, gin.H{
		"result": returnData,
	})
}
