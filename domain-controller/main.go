package main

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"strings"
	"time"

	//ps "github.com/zhongpei/go-powershell"
	//"github.com/zhongpei/go-powershell/backend"
	"github.com/ycyun/go-powershell"
	"net/http"
)

var log = logrus.New() //.WithField("who", "Main")
var Version = "development"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST, PATCH, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func setup() {
	log.SetFormatter(&nested.Formatter{
		HideKeys: false,
		//FieldsOrder: []string{"component", "category"},
		CallerFirst: false,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			return fmt.Sprintf(" [%s:%d][%s()]", path.Base(f.File), f.Line, funcName)
		},
	})
	log.SetReportCaller(true)
}
func main() {
	var (
		err error
	)

	l, err := setupLdap()

	shell, err := powershell.New()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := shell.Exit()
		if err != nil {
			log.Errorln(err)
			return
		}
	}()

	setup()

	router := gin.Default()
	router.Use(CORSMiddleware())

	//testLDAP()

	conn, status, err := ConnectAD()
	if err != nil {
		log.Errorln(err)
	}
	if !status {
		log.Errorln(status, err)
	}
	//router.LoadHTMLGlob("templates/*")

	/*
		go func() {
			//login check
			for i=0; i>0;  {

					Get-EventLog System -Source Microsoft-Windows-WinLogon | where {($_.instanceID -eq 7001) -or ($_.instanceID -eq 7002)} | select *| ogv
					Get-EventLog security -source microsoft-windows-security-auditing  | where {($_.instanceID -eq 4624) -or ($_.instanceID -eq 4625)} | select * |ogv

			}
		}()
	*/
	router.Use(static.Serve("/", static.LocalFile("./app/dist/", true)))
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/version", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"version": Version})
			})
			v1.POST("/login", func(c *gin.Context) {
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
			})
			v1.GET("/cmd/", func(c *gin.Context) {
				var pscmd PSCMD
				if err := c.Bind(&pscmd); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"msg": err})
					return
				}
				if pscmd.CMD == "" {
					c.JSON(http.StatusNotFound, gin.H{"msg": "Command not found"})
				}
				if pscmd.TIMEOUT == 0 {
					pscmd.TIMEOUT = 1
				}

				c1 := make(chan []string, 1)
				go func() {
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
			})
			v1.POST("/user", func(c *gin.Context) {
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
				err := addUser(l, user)
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
			})
			v1.PATCH("/user/:username", func(c *gin.Context) {

				var user = make(ADUSER)
				var user_ USER
				err := c.ShouldBindUri(&user_)
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
			})
			v1.DELETE("/user/:username", func(c *gin.Context) {

				var user = make(ADUSER)
				var user_ USER
				err := c.ShouldBindUri(&user_)
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
			})
			v1.GET("/user/:username", func(c *gin.Context) {
				setLog()
				var user USER
				err := c.ShouldBindUri(&user)
				u, err := getUser(l, &user)
				if err != nil {
					c.JSON(http.StatusNotFound, gin.H{
						"username": user.Username,
						"err":      err.Error(),
					})
					return
				} else {
					u["username"] = user.Username
					c.JSON(http.StatusOK, u)
				}

			})
			v1.PUT("/user/:username", func(c *gin.Context) {


				setLog()
				var aduserStruct *ADUser
				aduser := ADUSER{}
				var user USER
				err := c.ShouldBindUri(&user)
				tmpval, err := c.MultipartForm()
				for key, val :=range tmpval.Value{
					aduser[key]=val
				}
				aduser["username"] = user.Username
				//err = c.ShouldBind(&aduserStruct)
				log.Infoln(err)
				log.Infoln(aduser)
				aduserStruct = NewADUser(aduser)
				log.Infoln(aduserStruct)
				aduser, err =modUser(l, aduser)
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
			})
			v1.GET("/user/", func(c *gin.Context) {
				setLog()
				u := searchUser(l)
				c.JSON(http.StatusGone, u) /*gin.H{
					"userID": 1,
					"username": user.Username,
				})*/
			})
			v1.GET("/app", func(c *gin.Context) {
				apps := getApps(shell)
				log.Infof("finish")
				c.JSON(http.StatusOK, gin.H{"app": apps})
			})
		}
	}

	log.WithFields(logrus.Fields{
		"serverVersion": Version,
	}).Infof("Starting application")
	err = router.Run("0.0.0.0:8083")
	//err = endless.ListenAndServe(":8083", router)
	fmt.Println(err)
}
