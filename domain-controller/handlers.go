package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func exeShellHandler(c *gin.Context) {
	var pscmd PSCMD
	if err := c.Bind(&pscmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	if pscmd.CMD == "" {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Command not found"})
		return
	}
	if pscmd.TIMEOUT == 0 {
		pscmd.TIMEOUT = 1
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
	var (
		stdout string
		stderr string
	)
	select {
	case res := <-c1:
		stdout = res[0]
		stderr = res[1]
	case <-time.After(time.Duration(pscmd.TIMEOUT) * time.Second):
		stdout = "Timeout Reached"
		stderr = "Timeout Reached"
	}

	log.Infof("cmd: %v, arg: %v, stdout: %v, stderr: %v", pscmd.CMD, pscmd.ARG, stdout, stderr)
	c.JSON(http.StatusOK, gin.H{"stdout": stdout, "stderr": stderr})
	return
}
func appListHandler(c *gin.Context) {
	setLog()
	shell, _ := setupShell()
	apps := getApps(shell)
	log.Infof("finish")
	c.JSON(http.StatusOK, gin.H{"app": apps})
	return
}

type loginModel struct {
	Login    bool     `json:"login"`
	Username string   `json:"username"`
	Groups   []string `json:"groups"`
	IsAdmin  bool     `json:"isAdmin"`
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
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	username := c.PostForm("username")
	userPW := c.PostForm("password")
	result, groups, isAdmin, err := login(conn, username, userPW)
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

	c.JSON(http.StatusOK, ret)
}

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
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg":      "AD Connection Failed",
				"userID":   -1,
				"username": "",
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
		err2 := delUser(l, user)
		if err2 != nil {
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"userID":   1,
			"username": userID,
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
		c.JSON(http.StatusRequestedRangeNotSatisfiable, gin.H{
			"userID":   1,
			"username": userID,
			"msg":      err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"userID":   1,
		"username": userID,
	})
}
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
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"err": err.Error(),
		})
		return

	}
	u := listUser(l)
	c.JSON(http.StatusGone, u)
	return /*gin.H{
		"userID": 1,
		"username": user.Username,
	})*/
}
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
		c.JSON(http.StatusNotFound, gin.H{
			"username": user.Username,
			"err":      err.Error(),
		})
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

	err = setPassword(l, user, c.PostForm("password"))
	if err != nil {
		log.Errorln(err)
		c.JSON(http.StatusRequestedRangeNotSatisfiable, gin.H{
			"userID":   1,
			"username": user["username"],
			"msg":      err.Error(),
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

	tmpval, err := c.MultipartForm()
	addgroup := make(map[string][]string)
	for key, val := range tmpval.Value {
		addgroup[key] = val
	}

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
	c.JSON(http.StatusGone, addgroup)
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
	c.JSON(http.StatusGone, u)
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
			"err":      err.Error(),
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
func addUserToGroupHandler(c *gin.Context){
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

	c.JSON(http.StatusAccepted, group_)
	return
}

func deleteUserFromGroupHandler(c *gin.Context){
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
	if err != nil  || group_.Groupname == ""{
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"err": "Groupname is not provided",
		})
		return
	} else if group_.Groupname != "" && err == nil{
		err = delGroup(l, group_.Groupname)
		if err != nil{
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
