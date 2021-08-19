package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func testFunc(c *gin.Context) {
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

	result, err := rowsToString(rows)
	fmt.Println(result)
	jsonUnmarshal := []Workspace{}
	err = json.Unmarshal([]byte(result), &jsonUnmarshal)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	var params []MoldParams
	templateResult := getTemplate(params)
	serviceOfferingResult := getComputeOffering(params)
	c.JSON(http.StatusOK, gin.H{
		"result":              jsonUnmarshal,
		"templateList":        templateResult,
		"serviceOfferingList": serviceOfferingResult,
	})
}
