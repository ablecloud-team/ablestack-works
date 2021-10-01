package main

import (
	"context"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"os"
	"os/exec"
	_ "works-api/docs"
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
	log.SetLevel(logrus.DebugLevel)
}
func main() {
	var (
		err error
	)
	setup()
	DBSetting()        //DB 접속정보 셋팅
	MoldSetting()      //Mold 정보 셋팅
	DCSetting()        //DC 정보 셋팅
	WorksSetting()     //Works-API 정보 셋팅
	SambaSetting()     //SAMBA 정보 셋팅
	GuacamoleSetting() //guacamole 정보 셋팅

	router := gin.Default()
	router.Use(SetHeader)
	//router.LoadHTMLGlob("templates/*")
	router.Use(static.Serve("/", static.LocalFile("./app/dist/", true)))
	router.Use(static.Serve("/swagger/", static.LocalFile("./swagger", true)))
	api := router.Group("/api")
	{
		api.POST("/login", getLogin)
		api.POST("/workspaceAgent", putWorkspacesAgent)
		v1 := api.Group("/v1")
		v1.Use(checkToken)
		{
			v1.GET("/token", getUserToken)

			v1.GET("/workspace", getWorkspaces)
			v1.GET("/workspace/:workspaceUuid", getWorkspacesDetail)
			v1.PUT("/workspace", putWorkspaces)

			v1.GET("/offering", getOffering)

			v1.GET("/instance/:instanceUuid", getInstances)
			v1.GET("/instance/detail/:instanceUuid", getInstancesDetail)
			v1.PUT("/instance", putInstances)
			v1.POST("/instance", postInstances)
			v1.PATCH("/instance/:action/:instanceUuid", patchInstances)

			v1.GET("/version", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"version": Version})
			})

			v1.GET("/logout", getLogout)

			v1.GET("/user/:username", getUserDetail)
			v1.GET("/user", getUser)
			v1.PUT("/user", putUser)

			v1.GET("/group", getGroup)
			v1.GET("/group/:groupName", getGroupDetail)
			v1.DELETE("/group/:groupName", delGroupDetail)
			v1.DELETE("/group/:groupName/:userName", delDeleteUserToGroup)
			v1.PUT("/group/:groupName/:userName", putAddUserToGroup)
		}
		test := api.Group("/test")
		{
			test.POST("/test", testFunc)
		}
	}

	log.WithFields(logrus.Fields{
		"serverVersion": Version,
	}).Infof("Starting application")
	go asyncJobMonitoring()
	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	err = router.Run("0.0.0.0:8083")
	fmt.Println(err)
}
