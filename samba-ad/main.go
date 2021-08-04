package main

import (
	"context"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"strings"

	"net/http"
)

var log = logrus.New() //.WithField("who", "Main")
var Version = "development"

var (
	nodeContextCancel context.CancelFunc
)

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
		CustomCallerFormatter: func(f *runtime.Frame) string {
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			return fmt.Sprintf(" [%s:%d][%s()]", path.Base(f.File), f.Line, funcName)
		},
	})
	log.SetReportCaller(true)
}
func main() {
	l,err:=setupLdap()
	user := NewADUser()
	user.accountname= "testuser"
	user.sn= "testsn"
	user.givenName= "이름"

	//setPassword(l, user, "Ablecloud1!")


	setup()

	router := gin.Default()
	router.Use(CORSMiddleware())


	router.Use(static.Serve("/", static.LocalFile("./app/dist/", true)))
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/version", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"version": Version})
			})
			v1.POST("/user", func(c *gin.Context) {
				userID := c.PostForm("username")
				userPW := c.PostForm("password")
				userPhone := c.PostForm("phone")
				userMail := c.PostForm("email")
				log.Infof("%v, %v, %v, %v", userID, userPW, userPhone, userMail)
				c.JSON(http.StatusCreated, gin.H{
					"userID":1,
					"username":userID,
				})
			})
			v1.PUT("/user", func(c *gin.Context) {
				user := NewADUser()
				user.username= c.PostForm("username")
				log.Infof("%v, %v", user, password)

				output, stderr, err:=setPassword(l, user, c.PostForm("password"))
				log.Infof("%v", user)
				//cp:=exec.Command("bash", "-c", cmd)
				//cpOut, err := cp.Output()
				log.Infof(output)
				log.Infof(stderr)
				log.Infoln(err)

				c.JSON(http.StatusGone, gin.H{
					"userID":1,
					"cmd":string(output),
					"err":err,
				})
			})
		}
	}

	log.WithFields(logrus.Fields{
		"serverVersion": Version,
	}).Infof("Starting application")
	err = router.Run("0.0.0.0:8082")
	//err = endless.ListenAndServe(":8083", router)
	fmt.Println(err)
}
