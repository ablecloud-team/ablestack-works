package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// getConfiguration godoc
// @Summary Works 의 환경 설정을 조회하는 API
// @Description Works 의 환경 설정을 조회하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags Configurationt
// @Router /api/v1/configuration [GET]
// @Success 200 {object} map[string]interface{}
func getConfiguration(c *gin.Context) {
	returnCode := http.StatusNotFound
	configuration, err := selectConfigurationList()
	if err != nil {

	} else {
		returnCode = http.StatusOK
	}
	c.JSON(returnCode, gin.H{
		"result": configuration,
	})
}

// patchConfiguration godoc
// @Summary Works 의 환경 설정을 변경하는 API
// @Description Works 의 환경 설정을 변경하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags Configuration
// @Param id path string true "설정 값 ID"
// @Param value path string true "변경될 설정 값"
// @Router /api/v1/configuration/:id [PATCH]
// @Success 200 {object} map[string]interface{}
func patchConfiguration(c *gin.Context) {
	returnCode := http.StatusNotFound
	configuration := Configuration{}
	configuration.Id = c.Param("id")
	configuration.Value = c.PostForm("value")
	resultInt64, err := updateConfiguration(configuration)
	if err != nil {

	} else if resultInt64 >= int64(0) {
		returnCode = http.StatusOK
	}
	c.JSON(returnCode, gin.H{
		"result": resultInt64,
	})
}
