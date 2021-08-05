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

	"net/http"
)

var log = logrus.New() //.WithField("who", "Main")
var Version = "development"

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
	var err error
	setup()
	//l, err := setupLdap()

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
					"userID":   1,
					"username": userID,
				})
			})
			v1.PATCH("/user", func(c *gin.Context) {
				username = c.PostForm("username")
				log.Infof("%v, %v", username, password)

				output, stderr, err := setPassword(username, c.PostForm("password"))
				log.Println("stdout : ", output)
				log.Println("stderr : ", stderr)
				log.Println("err : ", err)
				if err != nil {

					c.JSON(http.StatusBadRequest, gin.H{
						"username": username,
						"stdout":   output,
						"stderr":   stderr,
						"err":      err,
					})
					return
				}

				c.JSON(http.StatusAccepted, gin.H{
					"username": username,
					"stdout":   output,
					"stderr":   stderr,
					"err":      err,
				})
			})
		}
	}

	log.WithFields(logrus.Fields{
		"serverVersion": Version,
	}).Infof("Starting application")
	err = router.Run("0.0.0.0:8082")
	fmt.Println(err)
}
