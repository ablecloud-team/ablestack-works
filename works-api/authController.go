package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

// getLogin godoc
// @Summary 사용자 로그인 하는 API
// @Description 사용자 로그인 하는 API 입니다.
// @Accept  json
// @Produce  json
// @Tags Login
// @Param id path string true "사용자 계정"
// @Param password path string true "사용자 비밀번호"
// @Router /api/login [POST]
// @Success 200 {object} map[string]interface{}
// @Success 400 {object} map[string]interface{} "DC 서버와 통신이 안된경우 발생"
// @Success 401 {object} map[string]interface{} "로그인처리는 정상적으로 되었으나 토큰생성에서 에러가 발생한경우"
// @Success 404 {object} map[string]interface{} "DC 서버와 통신이 통신은 정상적이나 계정이나 비밀번호가 일치 하지 않는경우"
func getLogin(c *gin.Context) {
	userId := c.PostForm("id")
	userPassword := c.PostForm("password")
	resultLogin, err := postLogin(userId, userPassword)
	result := map[string]interface{}{}
	if err != nil {
		log.Errorf("result [%v], error [%v]", result, err)
		result["message"] = "Communication with the DC server failed."
		c.JSON(http.StatusBadRequest, gin.H{
			"result": result,
		})
		return
	} else {
		if resultLogin.Status == OK200 {
			json.NewDecoder(resultLogin.Body).Decode(&result)
			loginBool := result["login"].(bool)
			log.Debugln(loginBool)
			if loginBool == false {
				result["message"] = "Login failed"
				c.JSON(http.StatusNotFound, gin.H{
					"result": result,
				})
				return
			} else if loginBool == true {
				token, err := createToken(userId)
				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						result["message"] = "token is expired"
						c.JSON(http.StatusUnauthorized, gin.H{
							"result": result,
						})
						return
					}
				}
				log.Infof("%v token is [%v]", userId, token)
				result["token"] = token
				//c.Header("access-token", token)
				//log.Info(c.sameSite)
				c.JSON(http.StatusOK, gin.H{
					"result": result,
				})
				return
			}
		}
	}
}

// getLogout godoc
// @Summary 사용자 로그아웃 하는 API
// @Description 사용자 로그아웃 하는 API 입니다.
// @Tags Login
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
// @Tags User
// @Produce  json
// @Param username path string true "사용자 계정"
// @Router /api/v1/user/:userName [get]
// @Success 200 {object} map[string]interface{}
func getUserDetail(c *gin.Context) {
	//result := map[string]interface{}{}
	username := c.Param("userName")
	var resultCode int
	//userId := c.Param("userId")
	result := getUserInfo(username)
	var res map[string]interface{}
	json.NewDecoder(result.Body).Decode(&res)
	log.Info(result.Status)
	log.Info(res["msg"])
	log.Info(res["target"])
	if result.Status == OK200 {
		resultCode = http.StatusOK
	} else if result.Status == NotFound404 {
		resultCode = http.StatusNotFound
		if res["msg"].(string) == NotFound {
			res["msg"] = MessageAccountNotFound
		}
	}

	c.JSON(resultCode, gin.H{
		"result": res,
	})
}

// getUserDetail godoc
// @Summary 사용자 토큰을 조회하는 API
// @Description 사용자 토큰을 조회를 위한 API 입니다.
// @Description 해더에 로그인시 생성된 토큰값을 "Authorization" 키값에 넣어주시면 됩니다.
// @Accept  json
// @Tags User
// @Produce  json
// @Router /api/v1/token [get]
// @Success 200 {object} map[string]interface{}
func getUserToken(c *gin.Context) {
	//result := map[string]interface{}{}
	//userId := c.Param("userId")
	cookieUserId := c.MustGet("cookie-user-id").(string)
	fmt.Println("cookieUserId = " + cookieUserId)
	result := getUserInfo(cookieUserId)
	var res map[string]interface{}
	json.NewDecoder(result.Body).Decode(&res)
	c.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

// getUser godoc
// @Summary 사용자 리스트를 조회 하는 API
// @Description 사용자 리스트를 조회를 위한 API 입니다.
// @Accept  json
// @Tags User
// @Produce  json
// @Router /api/v1/user [get]
// @Success 200 {object} map[string]interface{}
func getUser(c *gin.Context) {
	result := map[string]interface{}{}
	userList, err := getUserList()
	log.Info(result["result"])
	if err != nil {
		result["message"] = "Communication with the DC server failed."
		c.JSON(http.StatusBadRequest, gin.H{
			"result": result,
		})
	} else if len(userList) == 0 {
		result["message"] = "There is no user list."
		c.JSON(http.StatusNotFound, gin.H{
			"result": result,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": userList,
		})
	}
}

// deleteUser godoc
// @Summary 사용자 리스트를 조회 하는 API
// @Description 사용자 리스트를 조회를 위한 API 입니다.
// @Accept  json
// @Tags User
// @Produce  json
// @Param userName path string true "사용자 계정"
// @Router /api/v1/user/:userName [delete]
// @Success 200 {object} map[string]interface{}
func deleteUser(c *gin.Context) {
	username := c.Param("userName")
	//result := map[string]interface{}{}
	deleteDCUserResult, err := deleteDCUser(username)
	log.Infof("deleteDCUserResult [%v], err [%v]", deleteDCUserResult, err)
	//result["status"] = http.StatusOK
	c.JSON(http.StatusNoContent, gin.H{
		"result": "deleteDCUserResult",
	})
}

// putUser godoc
// @Summary 사용자 생성하는 하는 API
// @Description 사용자 생성을 위한 API 입니다.
// @Accept  json
// @Tags User
// @Produce  json
// @Param username path string true "사용자 계정"
// @Param password path string true "사용자 계정 비밀번호"
// @Param phone path string false "사용자 전화번호"
// @Param email path string true "사용자 이메일"
// @Param firstName path string true "사용자 성"
// @Param lastName path string true "사용자 이름"
// @Param title path string false "사용자 직급"
// @Param department path string false "사용자 부서"
// @Router /api/v1/user [put]
// @Success 200 {object} map[string]interface{}
func putUser(c *gin.Context) {
	resultValue := map[string]interface{}{}
	res := map[string]interface{}{}
	resultCode := http.StatusUnauthorized
	userInfo := UserInfo{}
	if c.PostForm("username") != "" {
		userInfo.Cn = c.PostForm("username")
	} else {
		resultValue["msg"] = "There is no username."
		c.JSON(http.StatusBadRequest, gin.H{
			"result": resultValue,
		})
		c.Abort()
		return
	}
	if c.PostForm("password") != "" {
		if RegexpPassword(c.PostForm("password")) {
			userInfo.Password = c.PostForm("password")
		} else {
			resultValue["msg"] = "The password format is incorrect.."
			c.JSON(http.StatusBadRequest, gin.H{
				"result": resultValue,
			})
			c.Abort()
			return
		}
	} else {
		resultValue["msg"] = "There is no password."
		c.JSON(http.StatusBadRequest, gin.H{
			"result": resultValue,
		})
		c.Abort()
		return
	}
	if c.PostForm("firstName") != "" {
		userInfo.Sn = c.PostForm("firstName")
	} else {
		resultValue["msg"] = "There is no firstName."
		c.JSON(http.StatusBadRequest, gin.H{
			"result": resultValue,
		})
		c.Abort()
		return
	}
	if c.PostForm("lastName") != "" {
		userInfo.GivenName = c.PostForm("lastName")
	} else {
		resultValue["msg"] = "There is no lastName."
		c.JSON(http.StatusBadRequest, gin.H{
			"result": resultValue,
		})
		c.Abort()
		return
	}
	if c.PostForm("email") != "" {
		userInfo.Mail = c.PostForm("email")
	} else {
		//resultValue["msg"] = "There is no email."
		//c.JSON(http.StatusBadRequest, gin.H{
		//	"result": resultValue,
		//})
		//c.Abort()
		//return
		userInfo.Mail = ""
	}
	if c.PostForm("phone") != "" {
		userInfo.TelephoneNumber = c.PostForm("phone")
	} else {
		userInfo.TelephoneNumber = ""
	}
	if c.PostForm("title") != "" {
		userInfo.Title = c.PostForm("title")
	} else {
		userInfo.Title = ""
	}
	if c.PostForm("department") != "" {
		userInfo.Department = c.PostForm("department")
	} else {
		userInfo.Department = ""
	}
	result, errInsertDCUser := postDCUser(userInfo)
	if errInsertDCUser != nil {
		log.Error(errInsertDCUser)
		resultValue["insertDCUser"] = errInsertDCUser

		c.Abort()
	}
	log.Debug(result)
	if result.Status == Created201 {
		err := json.NewDecoder(result.Body).Decode(&res)
		if err != nil {
			log.Errorf("An error occurred while converting http.Response to JSON.")
		}
		resultValue["dcResult"] = res
		resultInsertUserDB := insertUserDB(userInfo)
		if resultInsertUserDB["status"] == http.StatusOK {
			resultCode = http.StatusOK
			resultValue["insertDBResult"] = resultInsertUserDB
		} else {
			result, _ := deleteDCUser(userInfo.Cn)
			log.Errorf("DC 에 유저 생성 후 Works DB에 insert 중 에러가 발생하여 rollback 되었습니다.")
			resultCode = http.StatusUnauthorized
			resultValue["message"] = "After creating a user on DC, an error occurred while inserting into the Works DB and it was rolled back."
			resultValue["result"] = result
		}
	} else if result.Status == Conflict409 {
		resultCode = http.StatusConflict
		err := json.NewDecoder(result.Body).Decode(&res)
		if err != nil {
			log.Errorf("An error occurred while converting http.Response to JSON.")
		}
		log.Debug(res)
		resultValue["message"] = res["msg"]
	}
	c.JSON(resultCode, gin.H{
		"result": resultValue,
	})
}

// getGroup godoc
// @Summary 그룹 리스트를 조회 하는 API
// @Description 그룹 리스트를 조회를 위한 API 입니다.
// @Tags Group
// @Accept  json
// @Produce  json
// @Router /api/v1/group [get]
// @Success 200 {object} map[string]interface{}
func getGroup(c *gin.Context) {
	result := []map[string]interface{}{}
	resultCode := http.StatusUnauthorized
	resultSelectGroup, _ := selectGroupList()
	log.Debug(resultSelectGroup)
	log.Debug(resultSelectGroup.Status)
	if resultSelectGroup.Status == OK200 {
		resultCode = http.StatusOK
		err := json.NewDecoder(resultSelectGroup.Body).Decode(&result)
		if err != nil {
			log.Errorf("An error occurred while converting http.Response to JSON.")
		}
	} else if resultSelectGroup.Status == Gone410 {
		resultCode = http.StatusGone
	}
	c.JSON(resultCode, gin.H{
		"result": result,
	})
}

// getGroupDetail godoc
// @Summary 그룹 리스트를 조회 하는 API
// @Description 그룹 리스트를 조회를 위한 API 입니다.
// @Tags Group
// @Param groupName path string true "사용자 계정"
// @Accept  json
// @Produce  json
// @Router /api/v1/group/:groupName [get]
// @Success 200 {object} map[string]interface{}
func getGroupDetail(c *gin.Context) {
	result := map[string]interface{}{}
	groupName := c.Param("groupName")
	resultCode := http.StatusUnauthorized
	resultSelectGroupDetail, _ := selectGroupDetail(groupName)
	log.Debug(resultSelectGroupDetail)
	log.Debug(resultSelectGroupDetail.Status)
	if resultSelectGroupDetail.Status == OK200 {
		resultCode = http.StatusOK
		err := json.NewDecoder(resultSelectGroupDetail.Body).Decode(&result)
		if err != nil {
			log.Errorf("An error occurred while converting http.Response to JSON.")
		}
	} else if resultSelectGroupDetail.Status == Gone410 {
		resultCode = http.StatusGone
	}
	c.JSON(resultCode, gin.H{
		"result": result,
	})
}

// delGroupDetail godoc
// @Summary 그룹을 삭제하는 API
// @Description 그룹을 삭제하기 위한 API 입니다.
// @Tags Group
// @Param groupName path string true "사용자 계정"
// @Accept  json
// @Produce  json
// @Router /api/v1/group [DELETE]
// @Success 200 {object} map[string]interface{}
func delGroupDetail(c *gin.Context) {
	result := map[string]interface{}{}
	groupName := c.Param("groupName")
	resultCode := http.StatusUnauthorized
	resultSelectGroupDetail, _ := deleteGroup(groupName)
	log.Debug(resultSelectGroupDetail)
	log.Debug(resultSelectGroupDetail.Status)
	if resultSelectGroupDetail.Status == OK200 {
		resultCode = http.StatusOK
		err := json.NewDecoder(resultSelectGroupDetail.Body).Decode(&result)
		if err != nil {
			log.Errorf("An error occurred while converting http.Response to JSON.")
		}
	} else if resultSelectGroupDetail.Status == Gone410 {
		resultCode = http.StatusGone
	}
	c.JSON(resultCode, gin.H{
		"result": result,
	})
}

// putAddUserToGroup godoc
// @Summary 그룹을 삭제하는 API
// @Description 그룹을 삭제하기 위한 API 입니다.
// @Tags Group
// @Param groupName path string true "사용자를 추가할 그룹 이름"
// @Param userName path string true "그룹에 추가할 사용자 계정"
// @Accept  json
// @Produce  json
// @Router /api/v1/group/:groupName/:userName [put]
// @Success 200 {object} map[string]interface{}
func putAddUserToGroup(c *gin.Context) {
	result := map[string]interface{}{}
	groupName := c.Param("groupName")
	userName := c.Param("userName")
	resultCode := http.StatusUnauthorized
	log.Debugf("groupName [%v], userName [%v]", groupName, userName)
	resultSelectGroupDetail, _ := insertAddUserToGroup(groupName, userName)
	respBody, err := ioutil.ReadAll(resultSelectGroupDetail.Body)
	if err != nil {
		log.Errorln(err)
		log.Errorln(respBody)
	}

	if resultSelectGroupDetail.Status == Accepted202 {
		err := json.Unmarshal(respBody, &result)
		if err != nil {
			log.Errorf("An error occurred while converting http.Response to JSON.")
			return
		}
		resultCode = http.StatusOK
	} else if resultSelectGroupDetail.Status == Gone410 {
		resultCode = http.StatusGone
	}
	c.JSON(resultCode, gin.H{
		"result": result,
	})
}

// delDeleteUserToGroup godoc
// @Summary 그룹에서 유저를 삭제하는 API
// @Description 그룹에서 유저를 삭제하는 API 입니다.
// @Tags Group
// @Param groupName path string true "사용자를 삭제할 그룹 이름"
// @Param userName path string true "그룹에 삭제할 사용자 계정"
// @Accept  json
// @Produce  json
// @Router /api/v1/group/:groupName/:userName [DELETE]
// @Success 200 {object} map[string]interface{}
func delDeleteUserToGroup(c *gin.Context) {
	result := map[string]interface{}{}
	groupName := c.Param("groupName")
	userName := c.Param("userName")
	resultCode := http.StatusUnauthorized
	log.Debugf("groupName [%v], userName [%v]", groupName, userName)
	resultSelectGroupDetail, _ := deleteAddUserToGroup(groupName, userName)
	respBody, err := ioutil.ReadAll(resultSelectGroupDetail.Body)
	if err != nil {
		log.Errorln(err)
		log.Errorln(respBody)
	}

	if resultSelectGroupDetail.Status == Accepted202 {
		err := json.Unmarshal(respBody, &result)
		if err != nil {
			log.Errorf("An error occurred while converting http.Response to JSON.")
			return
		}
		resultCode = http.StatusOK
	} else if resultSelectGroupDetail.Status == Gone410 {
		resultCode = http.StatusGone
	}
	c.JSON(resultCode, gin.H{
		"result": result,
	})
}
