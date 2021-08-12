package main

import (
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

func loginHandler(c *gin.Context) {
	setLog()
	conn, status, err := ConnectAD()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	username := c.PostForm("username")
	userPW := c.PostForm("password")
	result, groups, isAdmin, err := login(conn, username, userPW)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"login":    result, //fmt.Sprintf("%v %v %v %v",userID, result, groups, isAdmin),
			"userID":   -1,
			"username": username,
			"groups":   nil,
			"isAdmin":  false})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"login":    result, //fmt.Sprintf("%v %v %v %v",userID, result, groups, isAdmin),
		"userID":   1,
		"username": username,
		"groups":   groups,
		"isAdmin":  isAdmin,
	})
}

func addUserHandler(c *gin.Context) {
	setLog()
	conn, status, err := ConnectAD()
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
	u := searchUser(l)
	c.JSON(http.StatusGone, u)
	return /*gin.H{
		"userID": 1,
		"username": user.Username,
	})*/
}
func getUserHandler(c *gin.Context) {

	setLog()
	conn, status, err := ConnectAD()
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
	//userPW := c.PostForm("password")
	//userPhone := c.PostForm("phone")
	//userMail := c.PostForm("email")
	//
	//givenName := c.PostForm("givenName")
	//title := c.PostForm("title")
	//sn := c.PostForm("sn")
	//
	//if l.IsClosing() {
	//	l, err = setupLdap()
	//	if err != nil {
	//		log.Errorln("AD Connection Failed")
	//		c.JSON(http.StatusInternalServerError, gin.H{
	//			"msg":      "AD Connection Failed",
	//			"userID":   -1,
	//			"username": "",
	//		})
	//		return
	//	}
	//}
	//user := ADUSER{
	//	"username":        userID,
	//	"telephoneNumber": userPhone,
	//	"mail":            userMail,
	//	"givenName":       givenName,
	//	"title":           title,
	//	"sn":              sn,
	//}
	//err = modUser(l, user)
	//if err != nil {
	//	log.Errorln(err)
	//	err2 := delUser(l, user)
	//	if err2 != nil {
	//		return
	//	}
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"userID":   1,
	//		"username": userID,
	//	})
	//	return
	//}
	//err = setPassword(l, user, userPW)
	//if err != nil {
	//	log.Errorln(err)
	//	err3 := delUser(l, user)
	//	if err3 != nil {
	//		return
	//	}
	//	c.JSON(http.StatusRequestedRangeNotSatisfiable, gin.H{
	//		"userID":   1,
	//		"username": userID,
	//		"msg":      err.Error(),
	//	})
	//	return
	//}
	//u, err:=getUser(l, &user_)
	//if err != nil {
	//	log.Errorln(err)
	//	err3 := delUser(l, user)
	//	if err3 != nil {
	//		return
	//	}
	//	c.JSON(http.StatusRequestedRangeNotSatisfiable, gin.H{
	//		"userID":   1,
	//		"username": userID,
	//		"msg":      err.Error(),
	//	})
	//	return
	//}
	c.JSON(http.StatusAccepted, aduser)
	return
}
func setUserPasswordHandler(c *gin.Context) {

	setLog()
	conn, status, err := ConnectAD()
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
	conn, status, err := ConnectAD()
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
	u := searchUser(l)
	c.JSON(http.StatusGone, u)
	return
}
func listGroupHandler(c *gin.Context) {
	setLog()
	conn, status, err := ConnectAD()
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
	u := searchUser(l)
	c.JSON(http.StatusGone, u)
	return
}
func getGroupHandler(c *gin.Context) {
	setLog()
	conn, status, err := ConnectAD()
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
	u := searchUser(l)
	c.JSON(http.StatusGone, u)
	return
}
func setGroupHandler(c *gin.Context) {
	setLog()
	conn, status, err := ConnectAD()
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
	u := searchUser(l)
	c.JSON(http.StatusGone, u)
	return
}
func deleteGroupHandler(c *gin.Context) {
	setLog()
	conn, status, err := ConnectAD()
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
	u := searchUser(l)
	c.JSON(http.StatusGone, u)
	return
}
