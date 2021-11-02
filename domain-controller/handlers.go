package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type shellReturnModel struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

// exeShellHandler godoc
// @Summary powershell 명령 처리기
// @Description powershell 명령 처리기
// @Accept  multipart/form-data
// @Produce  json
// @Param cmd query string true "명령어"
// @Param arg query string false "인자"
// @Param timeout query int false "시간제한, 기본값"
// @Success 200 {object} shellReturnModel "명령 성공"
// @Failure 401 {object} errorModel "명령 실패"
// @Failure default {objects} string
// @Router /cmd/ [get]
func exeShellHandler(c *gin.Context) {

	var (
		pscmd  PSCMD
		stdout string
		stderr string
	)
	if err := c.Bind(&pscmd); err != nil {
		c.JSON(http.StatusBadRequest, errorModel{Msg: "입력 오류", Target: err.Error()})
		return
	}
	if pscmd.CMD == "" {
		c.JSON(http.StatusBadRequest, errorModel{Msg: "Command not found", Target: "powershell"})
		return
	}
	if pscmd.TIMEOUT == 0 {
		pscmd.TIMEOUT = 3
	}

	c1 := make(chan []string, 1)
	go func() {

		shell, _ := setupShell()
		stdout, err := shell.Exec(fmt.Sprintf("%v %v", pscmd.CMD, pscmd.ARG))

		if err == nil {
			c1 <- []string{stdout, ""}
		} else {
			c1 <- []string{stdout, err.Error()}
		}

	}()

	select {
	case res := <-c1:
		stdout = res[0]
		stderr = res[1]
	case <-time.After(time.Duration(pscmd.TIMEOUT) * time.Second):
		stdout = "Timeout Reached"
		stderr = "Timeout Reached"
	}

	log.Infof("cmd: %v, arg: %v, stdout: %v, stderr: %v", pscmd.CMD, pscmd.ARG, stdout, stderr)
	c.JSON(http.StatusOK, shellReturnModel{Stdout: stdout, Stderr: stderr})
	return
}

// appListHandler godoc
// @Summary 윈도우 앱 목록
// @Description 윈도우 시작메뉴에 등록된 프로그램의 정보를 출력하는 API
// @Produce  json
// @Success 200 {object} []APPVAL "목록 표시"
// @Failure 401 {object} errorModel "명령 실패"
// @Failure default {objects} string
// @Router /app [get]
func appListHandler(c *gin.Context) {
	setLog()
	shell, _ := setupShell()
	apps := getApps(shell)
	log.Infof("finish")
	c.JSON(http.StatusOK, gin.H{"app": apps})
	return
}

type simpleReturnModel struct {
	Msg string `json:"msg"`
}
type loginModel struct {
	Login    bool     `json:"login"`
	Username string   `json:"username"`
	Groups   []string `json:"groups"`
	IsAdmin  bool     `json:"isAdmin"`
}
type userModel struct {
	Username string `json:"username"`
	Msg      string `json:"msg"`
}
type policyModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type errorModel struct {
	Msg    string `json:"msg"`
	Target string `json:"target"`
}

// loginHandler godoc
// @Summary List accounts
// @Description 사용자 로그인 체크
// @Accept  multipart/form-data
// @Produce  json
// @Param username formData string true "사용자 이름"
// @Param password formData string true "사용자 암호"
// @Success 200 {object} loginModel "로그인 성공"
// @Failure 401 {object} loginModel "로그인 실패"
// @Failure default {objects} string
// @Router /login [post]
func loginHandler(c *gin.Context) {
	setLog()
	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorf("%v", err)
	}
	if !status {
		log.Errorf("%v, %v", status, err)
	}
	username := c.PostForm("username")
	userPW := c.PostForm("password")
	result, groups, isAdmin, err := login(conn, username, userPW)
	setLog(fmt.Sprintf("%v, %v, %v, %v", result, groups, isAdmin, err))
	if !status {
		log.Errorf("%v, %v", result, err)
	}
	//loginResult :=
	ret := loginModel{
		Login:    result,
		IsAdmin:  isAdmin,
		Username: username,
		Groups:   groups,
	}
	if err != nil || !result {
		c.JSON(http.StatusUnauthorized, ret)
		return
	}
	setLog(fmt.Sprintf("%v", ret))
	c.JSON(http.StatusOK, ret)
}

// addUserHandler godoc
// @Summary add New User account
// @Description 사용자 생성
// @Accept  multipart/form-data
// @Produce  json
// @Param username formData string true "사용자 이름"
// @Param password formData string true "사용자 암호"
// @Param phone formData string true "전화번호"
// @Param email formData string true "이메일"
// @Param givenName formData string true "이름"
// @Param sn formData string true "성"
// @Param title formData string true "직급"
// @Success 200 {object} userModel "사용자 생성 성공"
// @Failure 401 {object} userModel "사용자 생성 실패"
// @Failure default {objects} string
// @Router /user [post]
func addUserHandler(c *gin.Context) {
	setLog()
	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn
	//l, err := setupLdap()
	userID := c.PostForm("username")
	userPW := c.PostForm("password")
	userPhone := c.PostForm("phone")
	userMail := c.PostForm("email")

	givenName := c.PostForm("givenName")
	title := c.PostForm("title")
	sn := c.PostForm("sn")

	log.Infof("%v, %v, %v, %v", userID, userPW, userPhone, userMail)
	if l.IsClosing() {
		l, err = setupLdap()
		if err != nil {
			log.Errorln("AD Connection Failed")
			c.JSON(http.StatusInternalServerError, userModel{
				Msg:      "AD Connection Failed",
				Username: "",
			})
			return
		}
	}
	user := ADUSER{
		"username":        userID,
		"telephoneNumber": userPhone,
		"mail":            userMail,
		"givenName":       givenName,
		"title":           title,
		"sn":              sn,
	}
	err = addUser(l, user)
	if err != nil {
		log.Errorln(err)
		if strings.Contains(err.Error(), "Exists") {
			c.JSON(http.StatusConflict, errorModel{
				Msg:    err.Error(),
				Target: userID,
			})
			return
		}
		err2 := delUser(l, user)
		if err2 != nil {
			return
		}
		c.JSON(http.StatusBadRequest, userModel{
			Username: userID,
			Msg:      err.Error() + err2.Error(),
		})
		return
	}
	err = setPassword(l, user, userPW)
	if err != nil {
		log.Errorln(err)
		err3 := delUser(l, user)
		if err3 != nil {
			return
		}
		c.JSON(http.StatusRequestedRangeNotSatisfiable, userModel{
			Username: userID,
			Msg:      err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, userModel{
		Username: userID,
		Msg:      "OK",
	})
}

// listUserHandler godoc
// @Summary 사용자 목록 조회
// @Description 사용자 목록 조회
// @Accept  multipart/form-data
// @Produce  json
// @Success 200 {object} []ADUser "사용자 생성 성공"
// @Failure 401 {object} errorModel "사용자 생성 실패"
// @Failure default {objects} string
// @Router /user [get]
func listUserHandler(c *gin.Context) {
	setLog()
	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn
	//l, err := setupLdap()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, errorModel{Msg: err.Error(), Target: "listuser"})
		return

	}
	u := listUser(l)
	c.JSON(http.StatusOK, u)
	return /*gin.H{
		"userID": 1,
		"username": user.Username,
	})*/
}

// getUserHandler godoc
// @Summary 사용자 목록 조회
// @Description 사용자 목록 조회
// @Accept  multipart/form-data
// @Produce  json
// @Param username path string true "사용자 이름"
// @Success 200 {object} ADUser "사용자 생성 성공"
// @Failure 401 {object} errorModel "사용자 생성 실패"
// @Failure default {objects} string
// @Router /user/{username} [get]
func getUserHandler(c *gin.Context) {

	setLog()
	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn
	//l, err := setupLdap()
	var user USER
	err = c.ShouldBindUri(&user)
	u, err := getUser(l, &user)
	if err != nil {
		c.JSON(http.StatusNotFound, errorModel{Msg: err.Error(), Target: user.Username})
		return
	} else {
		isAdmin, _ := inGroup(conn, user.Username, "Administrators")
		u["username"] = user.Username
		u["isAdmin"] = isAdmin
		c.JSON(http.StatusOK, u)
		return
	}

}

func setUserHandler(c *gin.Context) {

	setLog()
	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn
	//l, err := setupLdap()
	var aduserStruct *ADUser
	aduser := ADUSER{}
	var user USER
	err = c.ShouldBindUri(&user)
	tmpval, err := c.MultipartForm()
	for key, val := range tmpval.Value {
		aduser[key] = val
	}
	aduser["username"] = user.Username
	//err = c.ShouldBind(&aduserStruct)
	log.Infoln(err)
	log.Infoln(aduser)
	aduserStruct = NewADUser(aduser)
	log.Infoln(aduserStruct)
	aduser, err = modUser(l, aduser)
	log.Errorln(err)
	c.JSON(http.StatusAccepted, aduser)
	return
}
func setUserPasswordHandler(c *gin.Context) {

	setLog()
	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn
	//l, err := setupLdap()
	var user = make(ADUSER)
	var user_ USER
	err = c.ShouldBindUri(&user_)
	user["username"] = user_.Username

	err2 := setPassword(l, user, c.PostForm("password"))
	log.Errorf("passworderr: %v", err2)
	if err2 != nil {
		c.JSON(http.StatusRequestedRangeNotSatisfiable, gin.H{
			"userID":   1,
			"username": user["username"],
			"msg":      err2.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"userID": 1,
	})
	return
}
func deleteUserHandler(c *gin.Context) {

	setLog()
	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn
	//l, err := setupLdap()
	var user = make(ADUSER)
	var user_ USER
	err = c.ShouldBindUri(&user_)
	user["username"] = user_.Username
	err = delUser(l, user)
	if err != nil {
		log.Errorln(err)
		c.JSON(http.StatusGone, gin.H{
			"userID":   -1,
			"username": user["username"],
			"msg":      err.Error(),
		})
		return
	}
	c.JSON(http.StatusGone, gin.H{
		"userID":   1,
		"username": user["username"],
	})
	return
}

func addGroupHandler(c *gin.Context) {
	setLog()

	tmpval := c.PostForm("groupname")
	log.Infof("%v, %v", tmpval)
	addgroup := make(map[string][]string)
	//for key, val := range tmpval.Value {
	//	addgroup[key] = val
	//}
	addgroup["groupname"] = []string{tmpval}
	//log.Infof("%v, %v", tmpval, addgroup)
	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn

	addstr, _ := json.Marshal(addgroup)
	groupname := ""

	if _, ok := addgroup["groupname"]; ok {
		groupname = addgroup["groupname"][0]
		err = addGroup(l, groupname)
		if err != nil {
			log.Infof("%v", err.Error())

			c.JSON(http.StatusServiceUnavailable, gin.H{
				"err": err.Error(),
			})
			return
		}
	} else if !ok {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"err": "Groupname is not provided",
		})
		return
	} else {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"err": "Unknown Error",
		})
		return

	}
	//u := listUser(l)
	addstr, _ = json.Marshal(addgroup)
	log.Infof(string(addstr))
	c.JSON(http.StatusCreated, addgroup)
	return
}

//func addConnectionHandler(c *gin.Context) {
//	setLog()
//
//	var user2 USER
//	err := c.ShouldBindUri(&user2)
//
//	connection := c.PostForm("connection")
//	guacparameterval := c.PostForm("parameter")
//	log.Infof("%v, %v", user2, guacparameterval)
//	guacparameter := strings.Split(guacparameterval, ",")
//	user := ADUser{username: user2.Username}
//
//	conn, status, err := ConnectAD()
//	defer conn.Conn.Close()
//	if err != nil {
//		log.Errorf("connection failed:[%v]", err)
//	}
//	if !status {
//		log.Errorf("connection failed:[%v]", status)
//	}
//	l := conn.Conn
//
//
//	if user.username != "" && connection != "" && guacparameterval != ""{
//		err = addConnection(l, user, connection, guacparameter)
//		if err != nil {
//			log.Infof("%v", err.Error())
//
//			c.JSON(http.StatusServiceUnavailable, gin.H{
//				"err": err.Error(),
//			})
//			return
//		}
//	} else if user.username == "" {
//		c.JSON(http.StatusServiceUnavailable, gin.H{
//			"err": "username is not provided",
//		})
//		return
//	} else if connection == "" {
//		c.JSON(http.StatusServiceUnavailable, gin.H{
//			"err": "connection is not provided",
//		})
//		return
//	} else if guacparameterval == "" {
//		c.JSON(http.StatusServiceUnavailable, gin.H{
//			"err": "parameter is not provided",
//		})
//		return
//	} else {
//		c.JSON(http.StatusServiceUnavailable, gin.H{
//			"err": "Unknown Error",
//		})
//		return
//
//	}
//	//u := listUser(l)
//	//addstr, _ = json.Marshal(addgroup)
//	//log.Infof(string(addstr))
//	c.JSON(http.StatusCreated, "Connection Created")
//	return
//}

func addConnectionHandler(c *gin.Context) {
	setLog()

	var Connection CONNECTION
	err := c.ShouldBindUri(&Connection)

	username := c.PostForm("username")
	guacparameterval := c.PostForm("parameter")
	log.Infof("%v, %v", username, guacparameterval)
	guacparameter := strings.Split(guacparameterval, ",")
	user := ADUser{username: username}
	connection := Connection.Connectionname

	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorf("connection failed:[%v]", err)
	}
	if !status {
		log.Errorf("connection failed:[%v]", status)
	}
	l := conn.Conn

	if user.username != "" && connection != "" && guacparameterval != "" {
		err = addConnection(l, user, connection, guacparameter)
		if err != nil {
			log.Infof("%v", err.Error())

			c.JSON(http.StatusServiceUnavailable, gin.H{
				"err": err.Error(),
			})
			return
		}
	} else if user.username == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"err": "username is not provided",
		})
		return
	} else if connection == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"err": "connection is not provided",
		})
		return
	} else if guacparameterval == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"err": "parameter is not provided",
		})
		return
	} else {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"err": "Unknown Error",
		})
		return

	}
	//u := listUser(l)
	//addstr, _ = json.Marshal(addgroup)
	//log.Infof(string(addstr))
	c.JSON(http.StatusCreated, "Connection Created")
	return
}

func deleteConnectionHandler(c *gin.Context) {
	setLog()

	var Connection CONNECTION
	err := c.ShouldBindUri(&Connection)

	connection := Connection.Connectionname

	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorf("connection failed:[%v]", err)
	}
	if !status {
		log.Errorf("connection failed:[%v]", status)
	}
	l := conn.Conn

	if connection != "" {
		err = delConnection(l, connection)
		if err != nil {
			log.Infof("%v", err.Error())

			c.JSON(http.StatusServiceUnavailable, gin.H{
				"err": err.Error(),
			})
			return
		}
	} else if connection == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"err": "connection is not provided",
		})
		return
	} else {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"err": "Unknown Error",
		})
		return

	}
	//u := listUser(l)
	//addstr, _ = json.Marshal(addgroup)
	//log.Infof(string(addstr))
	c.JSON(http.StatusCreated, "Connection Created")
	return
}

func listGroupHandler(c *gin.Context) {
	setLog()
	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn
	//l, err := setupLdap()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"err": err.Error(),
		})
		return

	}
	u := searchGroup(l)
	c.JSON(http.StatusOK, u)
	return
}
func getGroupHandler(c *gin.Context) {
	setLog()
	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn
	//l, err := setupLdap()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"err": err.Error(),
		})
		return

	}

	var group GROUP
	err = c.ShouldBindUri(&group)
	u, err := getGroup(l, &group)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"groupname": group.Groupname,
			"err":       err.Error(),
		})
		return
	} else {
		u["groupname"] = group.Groupname
		c.JSON(http.StatusOK, u)
		return
	}
}
func setGroupHandler(c *gin.Context) {
	setLog()
	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn
	//l, err := setupLdap()
	var adgroupStruct *ADGroup
	adgroup := ADGROUP{}
	var group GROUP
	err = c.ShouldBindUri(&group)
	tmpval, err := c.MultipartForm()
	for key, val := range tmpval.Value {
		adgroup[key] = val
	}
	adgroup["groupname"] = group.Groupname
	//err = c.ShouldBind(&adgroupStruct)
	log.Infoln(err)
	log.Infoln(adgroup)
	adgroupStruct = NewADGroup(adgroup)
	log.Infoln(adgroupStruct)
	adgroup, err = modGroup(l, adgroup)
	log.Errorln(err)
	c.JSON(http.StatusAccepted, adgroup)
	return
}
func addUserToGroupHandler(c *gin.Context) {
	setLog()
	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn

	adgroup := ADGROUP{}
	var group GROUP
	err = c.ShouldBindUri(&group)

	adgroup["groupname"] = group.Groupname

	aduser := ADUSER{}
	var user USER
	err = c.ShouldBindUri(&user)

	aduser["username"] = user.Username

	group_, err := addUserToGroup(l, aduser, adgroup)
	log.Infof("add user group: %v", group_)
	log.Infof("add user group: %v", err)
	if err == nil {
		log.Infof("add user group: %v", group_)
		c.JSON(http.StatusAccepted, group_)
		return
	} else if err.Error() == "Not Found" {
		log.Errorf("add user error: %v", err)
		c.JSON(http.StatusBadRequest, errorModel{Msg: "User Not found", Target: user.Username})
		return
	} else if strings.Contains(err.Error(), "No Such Object") {
		log.Errorf("add user error: %v", err)
		c.JSON(http.StatusBadRequest, errorModel{Msg: "Group Not found", Target: group.Groupname})
		return
	} else if strings.Contains(err.Error(), "Entry Already Exists") {
		log.Errorf("add user error: %v", err)
		c.JSON(http.StatusBadRequest, errorModel{Msg: "Already Exists", Target: fmt.Sprintf("%v@%v", user.Username, group.Groupname)})
		return
	} else if err != nil {
		log.Errorf("add user error: %v", err)
		c.JSON(http.StatusBadRequest, errorModel{Msg: "Unknown error", Target: "addUserToGroup"})
		return
	}
}

func deleteUserFromGroupHandler(c *gin.Context) {
	setLog()
	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn

	adgroup := ADGROUP{}
	var group GROUP
	err = c.ShouldBindUri(&group)

	adgroup["groupname"] = group.Groupname

	aduser := ADUSER{}
	var user USER
	err = c.ShouldBindUri(&user)

	aduser["username"] = user.Username

	group_, err := deleteUserFromGroup(l, aduser, adgroup)

	c.JSON(http.StatusAccepted, group_)
	return
}
func deleteGroupHandler(c *gin.Context) {
	setLog()

	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn
	//l, err := setupLdap()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"err": err.Error(),
		})
		return

	}

	var group_ GROUP
	err = c.ShouldBindUri(&group_)

	groupname := ""
	if err != nil || group_.Groupname == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"err": "Groupname is not provided",
		})
		return
	} else if group_.Groupname != "" && err == nil {
		err = delGroup(l, group_.Groupname)
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"err": "Delete group failed",
			})
			return
		}
	} else {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"err": "Unknown error",
		})
		return
	}

	c.JSON(http.StatusGone, groupname)
	return
}

func detachPolicyHandler(c *gin.Context) {
	policyname := c.Param("policyname")
	groupname := c.Param("groupname")
	//New-GPLink -name usb_block -Target "ou=dev3,dc=dc1,dc=local"
	setLog(fmt.Sprintf("policyname: %v, groupname%v", policyname, groupname))
	shell, err := setupShell()
	if err != nil {
		log.Errorf("%v", err)
	}
	stdout, err := shell.Exec(fmt.Sprintf("Remove-GPLink -name %v -Target \"ou=%v,%v\"", policyname, groupname, ADconfig.ADbasedn))
	if err != nil {
		errorlines := strings.Split(err.Error(), "\r\n")
		for i, line := range errorlines {

			log.Errorf("%v, %v", i, line)
		}
		c.JSON(http.StatusNotFound, errorModel{Msg: errorlines[0], Target: errorlines[3]})
		return
	}
	ret := PowershellOut2map(stdout)
	c.JSON(http.StatusOK, ret)
}

func attachPolicyHandler(c *gin.Context) {
	policyname := c.Param("policyname")
	groupname := c.Param("groupname")
	//New-GPLink -name usb_block -Target "ou=dev3,dc=dc1,dc=local"
	setLog(fmt.Sprintf("policyname: %v, groupname%v", policyname, groupname))
	shell, err := setupShell()
	if err != nil {
		log.Errorf("%v", err)
	}
	stdout, err := shell.Exec(fmt.Sprintf("New-GPLink -name %v -Target \"ou=%v,%v\" -Enforced Yes", policyname, groupname, ADconfig.ADbasedn))
	if err != nil {
		errorlines := strings.Split(err.Error(), "\r\n")
		for i, line := range errorlines {

			log.Errorf("%v, %v", i, line)
		}
		c.JSON(http.StatusNotFound, errorModel{Msg: errorlines[0], Target: errorlines[2]})
		return
	}
	ret := PowershellOut2map(stdout)
	c.JSON(http.StatusOK, ret)
}

func getGroupPolicyHandler(c *gin.Context) {
	setLog()
	groupname := c.Param("groupname")
	shell, err := setupShell()
	if err != nil {
		log.Errorf("%v", err)
	}

	stdout, err := shell.Exec(fmt.Sprintf("(Get-GPInheritance -Target \"ou=%v,%v\").GpoLinks.DisplayName", groupname, ADconfig.ADbasedn))

	//groups := PowershellOut2map(stdout)
	//log.Infof("%v", stdout)
	//log.Infof("%v", err)
	lines := strings.Split(stdout, "\r\n")
	var policylist []policyModel
	for _, line := range lines {
		if line != "Default Domain Controllers Policy" && line != "Default Domain Policy" && line != "" {
			log.Infof("policy: %v\n", line)
			description, err := shell.Exec(fmt.Sprintf("(get-gpo %v).description", line))
			if err != nil {
				log.Errorf("%v:, %v", description, err)
			}
			policylist = append(policylist, policyModel{Name: line, Description: strings.TrimSpace(description)})
		}
	}
	//log.Infof("stdout: %v, \nstderr: %v\n", stdout, err)

	c.JSON(http.StatusOK, policylist)
}

func getPolicyHandler(c *gin.Context) {
	setLog()
	policyname := c.Param("policyname")
	shell, err := setupShell()
	if err != nil {
		log.Errorf("%v", err)
	}

	stdout, err := shell.Exec(fmt.Sprintf("Get-GPO -name %v", policyname))

	groups := PowershellOut2map(stdout)
	log.Infof("%v", stdout)
	log.Infof("%v", err)
	//lines := strings.Split(stdout, "\r\n")
	//var policylist []policyModel
	//for _, line := range lines{
	//	if line != "Default Domain Controllers Policy" && line != "Default Domain Policy" && line != ""{
	//		log.Infof("policy: %v\n", line)
	//		description, err := shell.Exec(fmt.Sprintf("(get-gpo %v).description", line))
	//		if err != nil{
	//			log.Errorf("%v:, %v", description, err)
	//		}
	//		policylist = append(policylist, policyModel{Name: line, Description: strings.TrimSpace(description)})
	//	}
	//}
	////log.Infof("stdout: %v, \nstderr: %v\n", stdout, err)

	c.JSON(http.StatusOK, groups)
}

func bootStrapHandler(c *gin.Context) {

	shell, err := setupShell()
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"msg": fmt.Sprintf("%v", err)})
		return
	}
	aduser := fmt.Sprintf("%v\\%v", strings.Split(ADconfig.ADdomain, ".")[0], ADconfig.ADusername)
	service := fmt.Sprintf("c:\\Works-DC\\nssm.exe set Works-DC objectName %v %v", aduser, ADconfig.ADpassword)
	stdout1, err1 := shell.Exec(service)
	log.Infof("stdout: %v, \nstderr: %v\n", stdout1, err1)
	//service = fmt.Sprintf("c:\\Works-DC\\nssm.exe restart Works-DC > c:\\Works-DC\\nssm.txt")
	service = "shutdown /r /f /t 10"
	stdout2, err2 := shell.Exec(service)
	log.Infof("stdout: %v, \nstderr: %v\n", stdout2, err2)

	val := map[string]string{
		"stdout1": stdout1,
		"stdout2": stdout2,
		"err1":    fmt.Sprintf("%v", err1),
		"err2":    fmt.Sprintf("%v", err2),
	}
	ADconfig.BootStraped = true
	ADsave()
	c.JSON(http.StatusOK, val)
	return
}
func updatePolicyHandler(c *gin.Context) {
	err2 := ""
	currentWorkingDirectory, err := os.Getwd()
	policyfile := fmt.Sprintf("%v/%v/%v", currentWorkingDirectory, ADconfig.PolicyPATH, ADconfig.PolicyLIST)
	data, err := os.Open(policyfile)
	if err != nil {
		log.Fatalf("Can not find %v file, %v", policyfile, err)
		c.JSON(http.StatusNotFound, map[string]string{"msg": fmt.Sprintf("Can not find %v file, %v", policyfile, err)})
		return
	}

	byteValue, err := ioutil.ReadAll(data)
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"msg": fmt.Sprintf("%v", err)})
		return
	}
	var policyList []map[string]string
	err = json.Unmarshal(byteValue, &policyList)
	for _, policyItem := range policyList {
		shell, err := setupShell()
		if err != nil {
			c.JSON(http.StatusNotFound, map[string]string{"msg": fmt.Sprintf("%v", err)})
			return
		}
		policy := policyItem["name"]
		description := policyItem["description"]
		//aduser:=fmt.Sprintf("%v\\%v", strings.Split(ADconfig.ADdomain, ".")[0], ADconfig.ADusername)
		//service:=fmt.Sprintf("c:\\Works-DC\\nssm.exe set Works-DC objectName %v %v", aduser, ADconfig.ADpassword)
		//stdout, err := shell.Exec(service)
		//log.Infof("stdout: %v, \nstderr: %v\n", stdout, err)
		//service=fmt.Sprintf("c:\\Works-DC\\nssm.exe retate Works-DC")
		//stdout, err = shell.Exec(service)
		//log.Infof("stdout: %v, \nstderr: %v\n", stdout, err)

		//cmd := fmt.Sprintf("c:\\Works-DC\\psexec.exe -accepteula -u %v -p %v powershell import-gpo -BackupGpoName %v -TargetName %v -Path '%v/%v' -CreateIfNeeded", aduser, ADconfig.ADpassword, policy, policy, currentWorkingDirectory, ADconfig.PolicyPATH)
		cmd := fmt.Sprintf("import-gpo -BackupGpoName %v -TargetName %v -Path '%v/%v' -CreateIfNeeded", policy, policy, currentWorkingDirectory, ADconfig.PolicyPATH)
		log.Infof("%v", cmd)
		stdout, err := shell.Exec(cmd)
		log.Infof("stdout: %v, \nstderr: %v\n", stdout, err)
		if err != nil {
			err2 = fmt.Sprintf("%v, %v", err2, err)
			c.JSON(http.StatusInternalServerError,
				map[string]string{"stdout": stdout, "err2": err2})
			return
		}
		stdout, err = shell.Exec(fmt.Sprintf("$policy = Get-Gpo -Name '%v'", policy))
		if err != nil {
			err2 = fmt.Sprintf("%v, %v", err2, err)
			c.JSON(http.StatusInternalServerError,
				map[string]string{"stdout": stdout, "err2": err2})
			return
		}
		stdout, err = shell.Exec(fmt.Sprintf("$policy.Description = '%v' ", description))
		if err != nil {
			err2 = fmt.Sprintf("%v, %v", err2, err)
			c.JSON(http.StatusInternalServerError,
				map[string]string{"stdout": stdout, "err2": err2})
			return
		}

	}

	var policylist []policyModel
	c.JSON(http.StatusOK, policylist)
	return
}
func listPolicyHandler(c *gin.Context) {
	shell, err := setupShell()
	if err != nil {
		log.Errorf("%v", err)
	}
	stdout, err := shell.Exec("(get-gpo -all).displayname")
	if err != nil {
		log.Errorf("%v", err)
	}
	lines := strings.Split(stdout, "\r\n")
	var policylist []policyModel
	for _, line := range lines {
		if line != "Default Domain Controllers Policy" && line != "Default Domain Policy" && line != "" {
			log.Infof("policy: %v\n", line)
			description, err := shell.Exec(fmt.Sprintf("(get-gpo %v).description", line))
			if err != nil {
				log.Errorf("%v:, %v", description, err)
			}
			policylist = append(policylist, policyModel{Name: line, Description: strings.TrimSpace(description)})
		}
	}
	//log.Infof("stdout: %v, \nstderr: %v\n", stdout, err)

	c.JSON(http.StatusOK, policylist)
	return
}


func addComputerToGroupHandler(c *gin.Context) {
	setLog()
	computername := c.Param("computername")
	groupname := c.Param("groupname")

	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn

	//New-GPLink -name usb_block -Target "ou=dev3,dc=dc1,dc=local"
	setLog(fmt.Sprintf("computername: %v, groupname: %v", computername, groupname))
	computercn, _:= getComputerCN(l, computername)
	err = addComputerToGroup(l, computercn, groupname)
	if err != nil {
		//errorlines := strings.Split(err.Error(), "\r\n")
		//for i, line := range errorlines {
		//
		//	log.Errorf("%v, %v", i, line)
		//}
		c.JSON(http.StatusNotFound, errorModel{Msg: err.Error(), Target: "addComputerToGroup"})
		return
	}
	c.JSON(http.StatusOK, simpleReturnModel{Msg: "addComputerToGroup Succes"})
}

func delComputerFromGroupHandler(c *gin.Context) {
	setLog()
	computername := c.Param("computername")
	groupname := c.Param("groupname")

	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn

	//New-GPLink -name usb_block -Target "ou=dev3,dc=dc1,dc=local"
	setLog(fmt.Sprintf("computername: %v, groupname: %v", computername, groupname))
	computercn, _:= getComputerCN(l, computername)
	err = delComputerFromGroup(l, computercn, groupname)
	if err != nil {
		//errorlines := strings.Split(err.Error(), "\r\n")
		//for i, line := range errorlines {
		//
		//	log.Errorf("%v, %v", i, line)
		//}
		c.JSON(http.StatusNotFound, errorModel{Msg: err.Error(), Target: "delComputerFromGroup"})
		return
	}
	c.JSON(http.StatusOK, simpleReturnModel{Msg: "delComputerFromGroup Succes"})
}


// listComputerHandler godoc
// @Summary 사용자 목록 조회
// @Description 사용자 목록 조회
// @Accept  multipart/form-data
// @Produce  json
// @Success 200 {object} []ADUser "사용자 생성 성공"
// @Failure 401 {object} errorModel "사용자 생성 실패"
// @Failure default {objects} string
// @Router /user [get]
func listComputerHandler(c *gin.Context) {
	setLog()
	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn
	//l, err := setupLdap()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, errorModel{Msg: err.Error(), Target: "listuser"})
		return

	}
	u := listComputer(l)
	c.JSON(http.StatusOK, u)
	return /*gin.H{
		"userID": 1,
		"username": user.Username,
	})*/
}


// getComputerHandler godoc
// @Summary 사용자 목록 조회
// @Description 사용자 목록 조회
// @Accept  multipart/form-data
// @Produce  json
// @Success 200 {object} []ADUser "사용자 생성 성공"
// @Failure 401 {object} errorModel "사용자 생성 실패"
// @Failure default {objects} string
// @Router /user [get]
func getComputerHandler(c *gin.Context) {
	setLog()

	computername := c.Param("computername")
	conn, status, err := ConnectAD()
	defer conn.Conn.Close()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	l := conn.Conn
	//l, err := setupLdap()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, errorModel{Msg: err.Error(), Target: "getComputer"})
		return

	}
	u, err := getComputer(l, computername)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, errorModel{Msg: err.Error(), Target: "getComputer"})
		return

	}
	c.JSON(http.StatusOK, u)
	return /*gin.H{
		"userID": 1,
		"username": user.Username,
	})*/
}