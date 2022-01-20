package main

import (
	_ "domain-controller/docs"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
)

var log = logrus.New() //.WithField("who", "Main")[
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
		NoColors: true,
		//FieldsOrder: []string{"component", "category"},
		CallerFirst: false,
		CustomCallerFormatter: func(f *runtime.Frame) string {
			s := strings.Split(f.Function, ".")
			funcName := s[len(s)-1]
			return fmt.Sprintf(" [%s:%d][%s()]", path.Base(f.File), f.Line, funcName)
		},
	})

	log.SetLevel(logrus.TraceLevel)

	var writers []io.Writer
	if !ADconfig.Silent {
		writers = append(writers, os.Stdout)
	}
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_APPEND|os.O_RDWR|os.O_SYNC, 0666)
	if err == nil {
		writers = append(writers, file)
	} else {
		log.Infof("Failed to log to file, using default stderr: %v", err)
	}
	w := io.MultiWriter(writers...)

	log.SetOutput(w)

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

// @host ad-api.ablecloud.io
// @BasePath /api/v1
// @query.collection.format multi

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	var (
		err error
	)

	//l, err := setupLdap()
	setup()

	//ADsave()
	err = ADinit()
	if err != nil {
		log.Fatalf("ADinit: %s", err)
	}
	router := gin.Default()
	router.Use(CORSMiddleware())


	checkStatus()
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
				c.JSON(http.StatusOK, gin.H{"version": Version, "Bootstraped":ADconfig.BootStraped})
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

			//사용자에게 vm 할당
			//v1.POST("/user/:username", addConnectionHandler)
			v1.POST("/connection/:conname", addConnectionHandler)
			v1.DELETE("/connection/:conname", deleteConnectionHandler)

			v1.POST("/group", addGroupHandler)
			v1.GET("/group", listGroupHandler)
			v1.GET("/group/:groupname", getGroupHandler)
			v1.PUT("/group/:groupname", setGroupHandler)
			v1.PUT("/group/:groupname/:username", addUserToGroupHandler)
			v1.DELETE("/group/:groupname/:username", deleteUserFromGroupHandler)
			v1.DELETE("/group/:groupname", deleteGroupHandler)

			v1.GET("/group/:groupname/policy", getGroupPolicyHandler)
			v1.PUT("/group/:groupname/policy", attachPolicyHandler)
			v1.DELETE("/group/:groupname/policy", detachPolicyHandler)

			v1.PUT("/computer/:computername/:groupname", addComputerToGroupHandler)
			v1.DELETE("/computer/:computername/:groupname", delComputerFromGroupHandler)
			v1.GET("/computer/", listComputerHandler)
			v1.GET("/computer/:computername", getComputerHandler)
			/*
			backup-gpo usb_block
			import-gpo -BackupGpoName usb_block -TargetName usb_block_ -Path 'C:\reports\' -CreateIfNeeded

			New-GPLink -name usb_block -Target "ou=dev3,dc=dc1,dc=local"
			Remove-GPLink  -name usb_block -target "ou=dev3,dc=dc1,dc=local"
			Get-GPInheritance -Target "ou=dev3,dc=dc1,dc=local"

			clipboard_block:{2020CB72-2E18-420A-97FC-887557FB916A}
			usb_block:{4ACB1953-468E-4AA1-9203-CF8417197981}
			share_block:{05df62ba-64f1-4b73-be1c-b63e76bf3075}
			remotefx:{0f35f16e-9eaf-4b1a-9507-fa9dfe0e0028}
			install_block:{91ad635c-2486-4dd4-ba0c-7a361ba947d8}
			update_block:{142cae05-9f35-4a33-8da5-ce00af5d2af0}
			 */

			//초기화
			v1.PATCH("/", bootStrapHandler)
			//초기화
			v1.PATCH("/policy", updatePolicyHandler)

			//목록
			v1.GET("/policy", listPolicyHandler)

			//정보
			v1.GET("/policy/:policyname", getPolicyHandler)

			//그룹에 적용
			v1.PUT("/policy/:policyname/:groupname", attachPolicyHandler)

			//그룹에 해제
			v1.DELETE("/policy/:policyname/:groupname", detachPolicyHandler)


			v1.GET("/status", statusHandler)
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

