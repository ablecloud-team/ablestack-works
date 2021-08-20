package main

import (
	"context"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	logrus "github.com/sirupsen/logrus"
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
		HideKeys:    false,
		CallerFirst: false,
	})
	log.SetReportCaller(true)
}
func main() {
	var (
		err error
	)
	setup()
	DBSetting()   //DB 접속정보 셋팅
	MoldSetting() //Mold 정보 셋팅
	DCSetting()   //DC 정보 셋팅

	router := gin.Default()
	router.Use(SetHeader)
	//router.LoadHTMLGlob("templates/*")
	router.Use(static.Serve("/", static.LocalFile("./app/dist/", true)))
	api := router.Group("/api")
	{
		api.POST("/login", loginController)
		api.GET("/workspace", getWorkspaces)
		api.PUT("/workspace", putWorkspaces)
		api.GET("/offering", getOffering)
		v1 := api.Group("/v1")
		v1.Use(checkToken)
		{
			v1.GET("/version", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"version": Version})
			})
			v1.GET("/logout", logoutController)
			v1.GET("/user", userDetailController)
		}
		test := api.Group("/test")
		{
			test.GET("/test", testFunc)
		}
	}

	log.WithFields(logrus.Fields{
		"serverVersion": Version,
	}).Infof("Starting application")
	err = router.Run("0.0.0.0:8083")
	fmt.Println(err)
}
