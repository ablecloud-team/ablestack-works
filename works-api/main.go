package main

import (
	"context"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	logrus "github.com/sirupsen/logrus"
	//ps "github.com/zhongpei/go-powershell"
	//"github.com/zhongpei/go-powershell/backend"
	"github.com/ycyun/go-powershell"
	"net/http"
	"os"
	"os/exec"
)

var log = logrus.New() //.WithField("who", "Main")
var Version = "development"

var (
	nodeContextCancel context.CancelFunc
)

func StartClientApp() (*exec.Cmd, error) {
	var err error

	cmd := exec.Command("npm", "run", "serve")
	cmd.Dir = "./app"
	cmd.Stdout = os.Stdout

	if err = cmd.Start(); err != nil {
		return cmd, fmt.Errorf("error starting NPM: %w", err)
	}

	return cmd, nil
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, DELETE, POST")

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
	})
	log.SetReportCaller(true)
}
func main() {
	var (
		err error
	)
	var pscmd PSCMD

	shell, err := powershell.New()
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	setup()

	router := gin.Default()
	router.Use(CORSMiddleware())
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
				userID := c.PostForm("id")
				userPW := c.PostForm("pw")
				result, groups, isAdmin, err := login(conn, userID, userPW)
				if err != nil {
					c.JSON(http.StatusOK, gin.H{
						"login":   result, //fmt.Sprintf("%v %v %v %v",userID, result, groups, isAdmin),
						"userID":  userID,
						"groups":  nil,
						"isAdmin": false})
				}
				c.JSON(http.StatusOK, gin.H{
					"login":   result, //fmt.Sprintf("%v %v %v %v",userID, result, groups, isAdmin),
					"userID":  userID,
					"groups":  groups,
					"isAdmin": isAdmin,
				})
			})
			v1.GET("/cmd/:cmd/:arg", func(c *gin.Context) {
				if err := c.ShouldBindUri(&pscmd); err != nil {
					c.JSON(400, gin.H{"msg": err})
					return
				}

				stdout, err := shell.Exec(fmt.Sprintf("%v %v", pscmd.CMD, pscmd.ARG))
				if err != nil {
					panic(err)
				}

				log.Infof("cmd: %v, arg: %v, stdout: %v", pscmd.CMD, pscmd.ARG, stdout)
				c.JSON(http.StatusOK, gin.H{"version": stdout})
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
