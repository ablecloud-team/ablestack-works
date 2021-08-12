package main

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"path"
	"runtime"
	"strings"
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

	//l, err := setupLdap()

	setup()

	router := gin.Default()
	router.Use(CORSMiddleware())

	//testLDAP()

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
			v1.POST("/login", loginHandler)

			v1.GET("/cmd/", exeShellHandler)
			v1.GET("/app", appListHandler)

			v1.POST("/user", addUserHandler)
			v1.GET("/user/:username", getUserHandler)
			v1.GET("/user/", listUserHandler)
			v1.PUT("/user/:username", setUserHandler)
			v1.PATCH("/user/:username", setUserPasswordHandler)
			v1.DELETE("/user/:username", deleteUserHandler)

			v1.POST("/group", addGroupHandler)
			v1.GET("/group", listGroupHandler)
			v1.GET("/group/:groupname", getGroupHandler)
			v1.PUT("/group/:groupname", setGroupHandler)
			v1.DELETE("/group/:groupname", deleteGroupHandler)

			/* TODO:
			사용자를 그룹에 추가(or그룹에 사용자를 추가)
			GPO 생성(?그냥 그룹이랑 같이?)
			GPO 읽기(적용된 정책 확인)
			그룹에 정책 적용
			그룹에 정책 제거
			적용 가능한 정책 목록

			*/

		}
	}

	log.WithFields(logrus.Fields{
		"serverVersion": Version,
	}).Infof("Starting application")
	err = router.Run("0.0.0.0:8083")
	fmt.Println(err)
}
