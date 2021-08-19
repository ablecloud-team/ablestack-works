package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func loginController(c *gin.Context) {
	var result map[string]interface{}
	userId := c.PostForm("id")
	userPassword := c.PostForm("password")
	result = login(userId, userPassword)
	token, err := createToken(userId)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized,
				gin.H{"status": http.StatusUnauthorized, "error": "token is expired"})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		c.Abort()
		return
	}
	c.SetCookie("access-token", token, 1800, "", "", false, false)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
func logoutController(c *gin.Context) {
	result := map[string]interface{}{}
	cookieUserId := c.MustGet("cookie-user-id").(string)
	fmt.Println("cookieUserId = " + cookieUserId)
	result["login"] = false
	result["username"] = cookieUserId
	c.SetCookie("access-token", "", -1, "", "", false, false)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func userDetailController(c *gin.Context) {
	var result map[string]interface{}
	//userId := c.Param("userId")
	cookieUserId := c.MustGet("cookie-user-id").(string)
	fmt.Println("cookieUserId = " + cookieUserId)
	result = userInfo(cookieUserId)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
