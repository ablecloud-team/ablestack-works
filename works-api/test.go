package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func testFunc(c *gin.Context) {
	//params := []MoldParams{
	//	{"serviceofferingid": "898e11e7-bbd2-425b-b457-4ea3953bf6d0"},
	//	{"templateid": "147fec65-81c7-4493-a6b7-60eaefa5fe3c"},
	//	{"zoneid": "6c7dfea1-b6dd-45c4-9c05-0faf1a3f6b6e"},
	//}
	//resultDeploy := getDeployVirtualMachine(params)
	m := map[string]interface{}{
		"id":    "bc9b16e5-52c5-408c-bb88-700ffec40df8",
		"jobid": "94f1a444-2156-4b4d-9244-8cf33727bd4e",
	}
	jsonStr, _ := json.Marshal(m)
	log.Info("00000000000000000000000000000000000000000")
	log.Info(jsonStr)
	log.Info("00000000000000000000000000000000000000000")
	mapInterfaceToJson(m)
	//result := mapInterfaceToJson(m)
	c.JSON(http.StatusOK, gin.H{
		//"resultDeploy": result,
	})
}
