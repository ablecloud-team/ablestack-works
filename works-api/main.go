package main

import (
	"context"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	logrus "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

var log = logrus.New() //.WithField("who", "Main")
var Version = "development"

var (
	//nodeContextCancel
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
		CallerFirst: false,
	})
	log.SetReportCaller(true)
}
func main() {
	var (
		err error
	)
	setup()

	router := gin.Default()
	router.Use(CORSMiddleware())
	//router.LoadHTMLGlob("templates/*")
	router.Use(static.Serve("/", static.LocalFile("./app/dist/", true)))
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/version", func(c *gin.Context) {
				login1()
				c.JSON(http.StatusOK, gin.H{"version": Version})
			})
			v1.POST("/login", func(c *gin.Context) {
				var result map[string]interface{}
				userId := c.PostForm("id")
				userPassword := c.PostForm("password")
				result = login(userId, userPassword)
				c.JSON(http.StatusOK, gin.H{
					"result":  result,
				})
			})
			v1.GET("/test", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"version": Version})
				request := map[string]string{
					"apikey": "h1IxdO9RWxtznF4V_nPpIWO0GoHg5oHdR8kHeve1dJD3f4rH14owxcZAu2n4ALpuCA6GzIy8akGHp83dhbeJuA",
					"command": "listVolumes",
					"response": "json",
				}
				sig := makeSignature(request)
				fmt.Println("========================")
				fmt.Println(sig)
				baseurl := "https://mold.ablecloud.io/client/api?"
				var strurl string
				for key, value := range request{
					strurl = strurl + key+"="+value+"&"
				}
				endUrl := baseurl+strurl+"signature="+sig
				fmt.Println(endUrl)
				resp, err := http.Get(endUrl)
				if err != nil {
					panic(err)
				}

				defer resp.Body.Close()

				data, err := ioutil.ReadAll(resp.Body)
				if err != nil{
					panic(err)
				}
				fmt.Println("%s\n", string(data))
			})
		}
	}

	log.WithFields(logrus.Fields{
		"serverVersion": Version,
	}).Infof("Starting application")
	err = router.Run("0.0.0.0:8083")
	fmt.Println(err)
}
