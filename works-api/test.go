package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func testFunc(c *gin.Context) {
	userName := c.PostForm("userName")
	insertPolicyRemotefx(userName)
	c.JSON(http.StatusOK, gin.H{

		"result": "",
	})
}
