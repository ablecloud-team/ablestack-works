package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func testFunc(c *gin.Context) {
	var params []MoldParams
	templateResult := getTemplate(params)
	//serviceOfferingResult := getComputeOffering(params)
	c.JSON(http.StatusOK, gin.H{
		"templateList": templateResult,
		//"serviceOfferingList": serviceOfferingResult,
	})
}
