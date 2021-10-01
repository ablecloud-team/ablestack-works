package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func testFunc(c *gin.Context) {
	userName := c.PostForm("userName")
	uuid := c.PostForm("uuid")
	log.Infof("username [%v], uuid [%v]", userName, uuid)
	log.Info(userName)
	result := insertUserAllocatedInstance(userName, uuid, "")
	log.Error(result)
	//log.Error(result.Status)
	//log.Error(result.Body)
	//res := map[string]interface{}{}
	//res1 := map[string]interface{}{}
	//body, err := ioutil.ReadAll(result.Body)
	//log.Errorf("error [%v]",err)
	//json.Unmarshal(body, &res)
	//log.Error(string(body))
	//res1["instanceDBInfo"] = result
	//log.Error(res)
	//params := []MoldParams{
	//	{"id": result.MoldUuid},
	//}
	//moldInstanceInfo := getListVirtualMachinesMetrics(params)
	//res1["instanceMoldInfo"] = moldInstanceInfo
	c.JSON(http.StatusOK, gin.H{

		"result": result,
	})
}
