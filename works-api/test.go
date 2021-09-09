package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func testFunc(c *gin.Context) {
	params := []MoldParams{
		{"resourceids": "4a54e716-ebd3-4ee6-99ad-9e7fa90a3131"},
		{"tags[0].key": "321"},
		{"tags[0].value": "123"},
	}
	result := getCreateTags(params)
	c.JSON(http.StatusOK, gin.H{
		"testResult": result,
	})
}
