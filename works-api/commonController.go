package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

// getServerCheck godoc
// api.GET("/serverCheck", getServerCheck)
// @Summary Works 서비스와 관련된 서버의 상태를 조회하는 API
// @Description Works 서비스와 관련된 서버의 상태를 조회하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags ServerCheck
// @Router /serverCheck [GET]
// @Success 200 {object} map[string]interface{}
func getServerCheck(c *gin.Context) {
	log.WithFields(logrus.Fields{
		"commonController.go": "getServerCheck",
	}).Infof("getServerCheck")
	var DCInfo = os.Getenv("DCUrl")
	var SambaInfo = "http://" + os.Getenv("SambaUrl") + ":9017"
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	result := map[string]int{
		WorksDc:    http.StatusNotFound,
		WorksSamba: http.StatusNotFound,
		Mold:       http.StatusNotFound,
		Guacamole:  http.StatusNotFound,
	}

	respDC, errDC := client.Get(DCInfo + "/v1/version")
	if errDC != nil {
		log.Error(errDC)
		log.Error(respDC)
	} else {
		result[WorksDc] = respDC.StatusCode
	}
	respSamba, errSamba := client.Get(SambaInfo + "/api/v1/version")
	if errDC != nil {
		log.Error(errSamba)
		log.Error(respSamba)
	} else {
		result[WorksSamba] = respSamba.StatusCode
	}
	respMold, errMold := getListApis()
	if errDC != nil {
		log.Error(errMold)
		log.Error(respMold)
	} else {
		result[Mold] = respMold.StatusCode
	}

	respGuacamole, errGuacamole := getGuacamoleToken()
	if errGuacamole != nil {
		log.Error(errGuacamole)
		log.Error(errGuacamole)
	} else {
		result[Guacamole] = respGuacamole.StatusCode
	}
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
