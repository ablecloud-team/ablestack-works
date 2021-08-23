package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func testFunc(c *gin.Context) {
	params := []MoldParams{
		{"serviceofferingid": "898e11e7-bbd2-425b-b457-4ea3953bf6d0"},
		{"templateid": "147fec65-81c7-4493-a6b7-60eaefa5fe3c"},
		{"zoneid": "6c7dfea1-b6dd-45c4-9c05-0faf1a3f6b6e"},
	}
	resultDeploy := getDeployVirtualMachine(params)
	c.JSON(http.StatusOK, gin.H{
		"resultDeploy": resultDeploy,
	})
}
