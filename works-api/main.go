package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"os"
	"os/exec"
	//_ "works-api/docs"
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

func main() {
	var (
		err error
	)
	DBSetting()          //DB 접속정보 셋팅
	MoldSetting()        //Mold 정보 셋팅
	DCSetting()          //DC 정보 셋팅
	WorksSetting()       //Works-API 정보 셋팅
	SambaSetting()       //SAMBA 정보 셋팅
	GuacamoleSetting()   //guacamole 정보 셋팅
	ClusterNameSetting() //clusterName 정보 셋팅
	logSetting()
	RDPPortSetting()
	//dcBootstrap()

	router := gin.Default()
	router.Use(SetHeader)
	//router.LoadHTMLGlob("templates/*")
	//router.Use(static.Serve("/", static.LocalFile("./app/dist/", true)))
	router.Use(static.Serve("/swagger/", static.LocalFile("./docs", true)))
	api := router.Group("/api")
	{
		api.POST("/login", getLogin)
		api.POST("/workspaceAgent/:instanceUuid", putWorkspacesAgent)
		api.POST("/workspaceAgent", putWorkspacesAgent)

		api.GET("/serverCheck", getServerCheck)

		api.GET("/version", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"version": Version})
		})
		v1 := api.Group("/v1")
		v1.Use(checkToken)
		//v1.Use(updateInstanceChecked0)
		{
			v1.GET("/logout", getLogout)

			v1.GET("/token", getUserToken)

			v1.GET("/dashboard", getDashboard)

			v1.GET("/workspace", getWorkspaces)
			v1.GET("/workspace/:workspaceUuid", getWorkspacesDetail)
			v1.POST("/workspace", putWorkspaces)
			v1.DELETE("/workspace/:workspaceUuid", deleteWorkspaces)

			v1.GET("/offering", getOffering)

			v1.GET("/instance/:instanceUuid", getInstances)
			v1.GET("/instance/detail/:instanceUuid", getInstancesDetail)
			v1.PUT("/instance", putInstances)
			//v1.POST("/instance", postInstances) // VDI 에 유저 할당
			v1.PATCH("/instance/:action/:instanceUuid", patchInstances)

			v1.POST("/connection/:instanceUuid/:username/:connection", putAppConnection)
			v1.PUT("/connection/:instanceUuid/:username", putConnection)
			v1.DELETE("/connection/:instanceUuid", deleteConnection)

			v1.GET("/configuration", getConfiguration)
			v1.PATCH("/configuration/:id", patchConfiguration)

			v1.PATCH("/handshake/:instanceUuid/:instanceType", patchHandshake)

			v1.GET("/user", getUser)
			v1.GET("/user/:userName", getUserDetail)
			v1.PUT("/user", putUser)
			v1.DELETE("/user/:userName", deleteUser)

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
	go updateInstanceChecked()
	url := ginSwagger.URL("/swagger/swagger.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	err = router.Run("0.0.0.0:8080")
	fmt.Println(err)
}
