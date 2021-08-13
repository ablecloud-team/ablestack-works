package main

import (
	"context"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/dgrijalva/jwt-go"
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
	setenv()

	router := gin.Default()
	router.Use(SetHeader)
	//router.LoadHTMLGlob("templates/*")
	router.Use(static.Serve("/", static.LocalFile("./app/dist/", true)))
	api := router.Group("/api")
	{
		api.POST("/login", func(c *gin.Context) {
			var result map[string]interface{}
			userId := c.PostForm("id")
			userPassword := c.PostForm("password")
			result = login(userId, userPassword)
			token, err := createToken(userId)
			if err != nil{
				if err == jwt.ErrSignatureInvalid{
					c.JSON(http.StatusUnauthorized,
						gin.H{"status": http.StatusUnauthorized, "error": "token is expired"})
					c.Abort()
					return
				}
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				c.Abort()
				return
			}
			c.SetCookie("access-token", token, 1800, "", "", false, false)
			c.JSON(http.StatusOK, gin.H{
				"result":  result,
			})
		})
		v1 := api.Group("/v1")
		v1.Use(checkToken)
		{
			v1.GET("/version", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"version": Version})
			})
			v1.GET("/logout", func(c *gin.Context) {
				result := map[string]interface{}{}
				cookieUserId := c.MustGet("cookie-user-id").(string)
				fmt.Println("cookieUserId = "+cookieUserId)
				result["login"] = false
				result["username"] = cookieUserId
				c.SetCookie("access-token", "", -1, "", "", false, false)
				c.JSON(http.StatusOK, gin.H{
					"result":  result,
				})
			})
			v1.GET("/user", func(c *gin.Context) {
				var result map[string]interface{}
				//userId := c.Param("userId")
				cookieUserId := c.MustGet("cookie-user-id").(string)
				fmt.Println("cookieUserId = "+cookieUserId)
				result = userInfo(cookieUserId)
				c.JSON(http.StatusOK, gin.H{
					"result":  result,
				})
			})
		}
		test := api.Group("/test")
		test.Use(checkToken)
		{
			test.GET("/version", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"version": Version})
			})
			test.GET("/test", func(c *gin.Context) {
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
					fmt.Println(err)
				}
				fmt.Println("%s\n", string(data))
			})
			test.GET("/test1", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"version": Version})
				request := map[string]string{
					"apikey": "h1IxdO9RWxtznF4V_nPpIWO0GoHg5oHdR8kHeve1dJD3f4rH14owxcZAu2n4ALpuCA6GzIy8akGHp83dhbeJuA",
					"command": "listVolumes",
					"response": "json",
				}
				sig := makeSignature(request)
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
			test.GET("/user/", func(c *gin.Context) {
				var result map[string]interface{}
				//userId := c.Param("userId")
				cookieUserId := c.MustGet("cookie-user-id").(string)
				fmt.Println("userId")
				fmt.Println(cookieUserId)
				fmt.Println(cookieUserId)
				result = userInfo(cookieUserId)
				c.JSON(http.StatusOK, gin.H{
					"result":  result,
				})
			})
		}
	}

	log.WithFields(logrus.Fields{
		"serverVersion": Version,
	}).Infof("Starting application")
	err = router.Run("0.0.0.0:8083")
	fmt.Println(err)
}
