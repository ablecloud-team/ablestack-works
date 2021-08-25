package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func getWorkspaces(c *gin.Context) {
	db, err := sql.Open(os.Getenv("MsqlType"), os.Getenv("DbInfo"))
	if err != nil {
		fmt.Println("DB connect error")
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println("DB connect success")

	rows, err := db.Query("SELECT * FROM workspaces")
	if err != nil {
		fmt.Println("worksInit Mold Setting Query Failed")
		fmt.Println(err.Error())
	}
	defer rows.Close()
	fmt.Println(Version)

	result, err := rowsToString(rows)
	fmt.Println(result)
	resultData := []Workspace{}
	err = json.Unmarshal([]byte(result), &resultData)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": resultData,
	})
}

func getOffering(c *gin.Context) {
	var params []MoldParams
	templateResult := getTemplate(params)
	serviceOfferingResult := getComputeOffering(params)
	networkResult := getNetwork(params)
	diskOfferingResult := getDiskOffering(params)
	c.JSON(http.StatusOK, gin.H{
		"templateList":        templateResult,
		"serviceOfferingList": serviceOfferingResult,
		"networkList":         networkResult,
		"diskOfferingList":    diskOfferingResult,
	})
}

func putWorkspaces(c *gin.Context) {
	workspace := Workspace{}

	workspace.Uuid = getUuid()
	workspace.Name = c.PostForm("name")
	workspace.Description = c.PostForm("description")
	workspace.Type = c.PostForm("type")
	workspace.TemplateUuid = c.PostForm("templateUuid")
	workspace.ComputeOfferingUuid = c.PostForm("computeOfferingUuid")
	workspace.DiskOfferingUuid = c.PostForm("diskOfferingUuid")
	workspace.NetworkUuid = selectNetworkDetail()

	result := insertWorkspace(workspace)
	params := []MoldParams{
		{"serviceofferingid": workspace.ComputeOfferingUuid},
		{"templateid": workspace.TemplateUuid},
		{"zoneid": selectZoneId()},
	}

	resultDeploy := getDeployVirtualMachine(params)
	asyncJob := AsyncJob{}
	asyncJob.Name = "VMDestroy"
	asyncJob.ExecUuid = resultDeploy["deployvirtualmachineresponse"].(map[string]interface{})["id"].(string)
	asyncJob.Ready = 0
	asyncJob.Parameter = "expunge=true"
	resultAsyncJob := insertAsyncJob(asyncJob)

	c.JSON(http.StatusOK, gin.H{
		"result":         result,
		"resultDeploy":   resultDeploy,
		"resultAsyncJob": resultAsyncJob,
	})
}

func putWorkspacesAgent(c *gin.Context) {
	uuidParams := c.PostForm("uuid")
	jobId := c.PostForm("jobId")

	result := insertWrokspaceAgent(uuidParams, jobId)

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
