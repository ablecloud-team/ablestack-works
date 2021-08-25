package main

import (
	_ "domain-controller/docs"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
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




// @title Ablecloud Works Domain-Controller API
// @version 1.0
// @description 이 API는 Ablecloud Works의 Domain Controller(DC)를 제어하는 역할을 합니다.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 10.10.1.25:8083
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	var (
		err error
	)

	//l, err := setupLdap()

	setup()
	err = sentry.Init(sentry.ClientOptions{
		TracesSampleRate: 1,
		Dsn: "https://f0f2fdd530cc4bfe86f60ba333bc3d6f@o961393.ingest.sentry.io/5909805",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2 * time.Second)


	router := gin.Default()
	router.Use(CORSMiddleware())
	router.Use(sentrygin.New(sentrygin.Options{}))

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
	router.Use(static.Serve("/swagger/", static.LocalFile("./swagger", true)))

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
			v1.PUT("/group/:groupname/:username", addUserToGroupHandler)
			v1.DELETE("/group/:groupname/:username", deleteUserFromGroupHandler)
			v1.DELETE("/group/:groupname", deleteGroupHandler)

			//생성
			v1.POST("/policy", addPolicyHandler)

			//목록
			v1.GET("/policy", listPolicyHandler)

			//정보
			v1.GET("/policy/:policyname", getPolicyHandler)

			//그룹에 적용
			v1.PUT("/policy/:policyname/:groupname", attachPolicyHandler)

			//그룹에 해제
			v1.DELETE("/policy/:policyname/:groupname", detachPolicyHandler)

			/* TODO:
			사용자를 그룹에 추가(or그룹에 사용자를 추가)
			https://mpain.tistory.com/tag/group%20policy
			GPO 생성(?그냥 그룹이랑 같이?)
			GPO 읽기(적용된 정책 확인)
			그룹에 정책 적용
			그룹에 정책 제거
			적용 가능한 정책 목록
			https://github.com/ablecloud-team/ablestack-desktop/issues/84
			 사용자 그룹 정보 변경
			 사용자 그룹에서 사용자 목록 조회
			*/

		}
	}

	log.WithFields(logrus.Fields{
		"serverVersion": Version,
	}).Infof("Starting application")
	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	err = router.Run("0.0.0.0:8083")
	fmt.Println(err)
}

func detachPolicyHandler(c *gin.Context) {
	c.Param("policyname")
	c.Param("groupname")
}

func attachPolicyHandler(c *gin.Context) {
	c.Param("policyname")
	c.Param("groupname")

}

func getPolicyHandler(c *gin.Context) {
	c.Param("policyname")
	shell, err := setupShell()
	if err != nil{
		log.Errorf("%v", err)
	}
	stdout, err := shell.Exec("Get-GPInheritance -Target \"ou=dev3,dc=dc1,dc=local\"")
	log.Infof("%v", stdout)
	log.Infof("%v", err)
}

func listPolicyHandler(c *gin.Context) {

}

func addPolicyHandler(c *gin.Context) {
	c.Param("policyname")
	shell, err := setupShell()
	if err != nil{
		log.Errorf("%v", err)
	}
	shell.Exec("import-gpo -BackupGpoName usb_block -TargetName usb_block_ -Path 'C:\\reports\\' -CreateIfNeeded")

}
