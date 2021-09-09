package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

// getLogin godoc
// @Summary 사용자 로그인 하는 API
// @Description 사용자 로그인 하는 API 입니다.
// @Accept  json
// @Produce  json
// @Param id path string true "사용자 계정"
// @Param password path string true "사용자 비밀번호"
// @Router /api/login [put]
// @Success 200 {object} map[string]interface{}
func getLogin(c *gin.Context) {
	var result map[string]interface{}
	userId := c.PostForm("id")
	userPassword := c.PostForm("password")
	result = login(userId, userPassword)
	loginBool := result["login"].(bool)
	log.Debugln(loginBool)
	if loginBool == false {
		c.JSON(http.StatusAccepted, gin.H{
			"result":  result,
			"message": "Login failed",
		})
		c.Abort()
	} else if loginBool == true {
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
		log.Infof("%v token is %v", userId, token)
		//log.Info(c.sameSite)
		//c.Header("access-token", token)
		result["token"] = token
		result["status"] = http.StatusOK
		c.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	}
}

// getLogout godoc
// @Summary 사용자 로그아웃 하는 API
// @Description 사용자 로그아웃 하는 API 입니다.
// @Accept  json
// @Produce  json
// @Router /api/v1/logout [get]
// @Success 200 {object} map[string]interface{}
func getLogout(c *gin.Context) {
	result := map[string]interface{}{}
	cookieUserId := c.MustGet("cookie-user-id").(string)
	fmt.Println("cookieUserId = " + cookieUserId)
	result["login"] = false
	result["username"] = cookieUserId
	//c.SetCookie("access-token", "", -1, "/", allowUrl, false, true)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// getUserDetail godoc
// @Summary 사용자 상세조회 하는 API
// @Description 사용자 상세 조회를 위한 API 입니다.
// @Description 해더에 로그인시 생성된 토큰값을 "Authorization" 키값에 넣어주시면 됩니다.
// @Accept  json
// @Produce  json
// @Router /api/v1/user/:username [get]
// @Success 200 {object} map[string]interface{}
func getUserDetail(c *gin.Context) {
	result := map[string]interface{}{}
	//userId := c.Param("userId")
	cookieUserId := c.MustGet("cookie-user-id").(string)
	fmt.Println("cookieUserId = " + cookieUserId)
	result = selectUserDetail(cookieUserId)
	result["status"] = http.StatusOK
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// getUserDetail godoc
// @Summary 사용자 토큰을 조회하는 API
// @Description 사용자 토큰을 조회를 위한 API 입니다.
// @Description 해더에 로그인시 생성된 토큰값을 "Authorization" 키값에 넣어주시면 됩니다.
// @Accept  json
// @Produce  json
// @Router /api/v1/token [get]
// @Success 200 {object} map[string]interface{}
func getUserToken(c *gin.Context) {
	result := map[string]interface{}{}
	//userId := c.Param("userId")
	cookieUserId := c.MustGet("cookie-user-id").(string)
	fmt.Println("cookieUserId = " + cookieUserId)
	result = selectUserDetail(cookieUserId)
	result["status"] = http.StatusOK
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// getUser godoc
// @Summary 사용자 리스트를 조회 하는 API
// @Description 사용자 리스트를 조회를 위한 API 입니다.
// @Accept  json
// @Produce  json
// @Router /api/v1/user [get]
// @Success 200 {object} map[string]interface{}
func getUser(c *gin.Context) {
	result := map[string]interface{}{}
	result["result"] = selectUserList()
	log.Info(result["result"])
	result["status"] = http.StatusOK
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

// putUser godoc
// @Summary 사용자 생성하는 하는 API
// @Description 사용자 생성을 위한 API 입니다.
// @Accept  json
// @Produce  json
// @Param username path string true "사용자 계정"
// @Param password path string true "사용자 계정 비밀번호"
// @Param phone path string true "사용자 전화번호"
// @Param email path string true "사용자 이메일"
// @Param firstName path string true "사용자 성"
// @Param lastName path string true "사용자 이름"
// @Param title path string true "사용자 직급"
// @Router /api/v1/user [put]
// @Success 200 {object} map[string]interface{}
func putUser(c *gin.Context) {
	resultValue := map[string]interface{}{}
	userInfo := UserInfo{}
	userInfo.Username = c.PostForm("username")
	userInfo.Password = c.PostForm("password")
	userInfo.Sn = c.PostForm("firstName")
	userInfo.GivenName = c.PostForm("lastName")
	userInfo.Email = c.PostForm("email")
	if c.PostForm("phone") != "" {
		userInfo.Phone = c.PostForm("phone")
	} else {
		userInfo.Phone = "000-0000-0000"
	}
	if c.PostForm("title") != "" {
		userInfo.Title = c.PostForm("title")
	} else {
		userInfo.Phone = "OfficeWorker"
	}
	result := putDCUser(userInfo)
	resultValue["status"] = http.StatusOK
	resultValue["result"] = result
	c.JSON(http.StatusOK, gin.H{
		"result": resultValue,
	})
}
