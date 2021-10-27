package main

import (
	"encoding/json"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/url"
	"strconv"
	"time"

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
	if !Agentconfig.Silent {
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




// @title Ablecloud Works windows agent API
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
	err = AgentInit()
	if err != nil {
		log.Fatalf("ADinit: %s", err)
	}
	interval := Agentconfig.Interval
	c1 := make(chan string, 1)
	loginchan := make(chan map[string]string, 1)
	logoutchan := make(chan map[string]string, 1)
	go healthCheck(loginchan, logoutchan, c1, interval)
	log.Infof("healthCehck thread started")

	router := gin.Default()

	router.Use(static.Serve("/swagger/", static.LocalFile("./swagger", true)))

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/version", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"version": Version})
			})
			v1.GET("/app", appListHandler)
		}
	}

	log.WithFields(logrus.Fields{
		"serverVersion": Version,
	}).Infof("Starting application")
	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	err = router.Run("0.0.0.0:8083")
	res := <-c1
	log.Errorf("error: %v", res)

	fmt.Println(err)
}

func appListHandler(c *gin.Context) {
	setLog()
	shell, _ := setupShell()
	apps := getApps(shell)
	log.Infof("finish")
	c.JSON(http.StatusOK, gin.H{"app": apps})
	return
}

func ipcheck() []string{
	shell, err := setupShell()
	if err != nil {
		log.Errorf("setupShell error: %v", err)
	}
	stdout, err := shell.Exec(" Get-NetIPAddress -AddressFamily ipv4| Format-table -Property ipaddress")
	if err != nil {
		log.Errorf("Get-NetIPAddress error: %v", err)
	}
	stdouts := strings.Split(strings.Replace(strings.TrimSpace(stdout), "\r\n", "\n", -1), "\n")[2:]
	ips := []string{}
	for _, ip := range stdouts{
		if ! strings.Contains(ip, "127.0.0.1"){
			ips = append(ips, ip)
		}
	}
	log.Infof("%v", stdouts)

	return ips
}

func logoutcheckfnc() map[string]string {
	shell, err := setupShell()
	stdout, err := shell.Exec("$elog=(Get-EventLog security -source microsoft-windows-security-auditing  | where {($_.instanceID -eq 4634)} | select TimeGenerated,ReplacementStrings,Message -First 1)")

	log.Infof("stdout: %v", stdout)
	if err != nil {
		log.Errorf("logoutcheck error: %v", err)
	}
	/*
		count, err := shell.Exec("$elog.ReplacementStrings.count")
		log.Infof("count: %v", count)
		if err != nil {
			log.Errorf("vdi-name error: %v", err)
		} //vdi-name
		lens, err:=strconv.Atoi(strings.TrimSpace(count))
		if err != nil {
			log.Errorf("atoi error: %v", err)
		} //vdi-name
		log.Infof("lens: %v", lens)
		for  i := 0; i<lens; i++ {
			out_, err := shell.Exec(fmt.Sprintf("$elog.ReplacementStrings[%v]", i))
			out := strings.TrimSpace(out_)
			log.Infof("i=%v: %v", i, out)
			if err != nil {
				log.Errorf("vdi-name error: %v", err)
			}
		}

	*/
	vdinameout_, err := shell.Exec("$elog.ReplacementStrings[1]")
	if err != nil {
		log.Errorf("vdi-name error: %v", err)
	} //vdi-name
	usernameout_, err := shell.Exec("$elog.ReplacementStrings[5]")
	if err != nil {
		log.Errorf("user-name error: %v", err)
	} //user-name


	logintimeout_, err := shell.Exec("get-date -Date $elog.TimeGenerated -UFormat \"%Y/%m/%d %T\"")
	if err != nil {
		log.Errorf("logintime error: %v", err)
	} //시간
	vdinameout := strings.TrimSpace(vdinameout_)
	usernameout := strings.TrimSpace(usernameout_)
	logintimeout := strings.TrimSpace(logintimeout_)
	//log.Infof("%v, %v, %v, ", vdinameout, usernameout, logintimeout)
	ret := map[string]string{
		"vdi":vdinameout,
		"username":usernameout,
		"time":logintimeout,
	}
	log.Infof("logout return: %v", ret)
	return ret
}
func logoutcheck(c chan map[string]string, c1 chan string, interval int) {
	for true{
		ret := logoutcheckfnc()
		c<-ret
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func logincheckfnc() map[string]string {
	shell, err := setupShell()
	stdout, err := shell.Exec("$elog=(Get-EventLog security -source microsoft-windows-security-auditing  | where {($_.instanceID -eq 4624)} | select TimeGenerated,ReplacementStrings,Message -First 1)")
	log.Infof("stdout: %v", stdout)
	if err != nil {
		log.Errorf("logincheck error: %v", err)
	}
	/*
		count, err := shell.Exec("$elog.ReplacementStrings.count")
		log.Infof("count: %v", count)
		if err != nil {
			log.Errorf("vdi-name error: %v", err)
		} //vdi-name


			lens, err:=strconv.Atoi(strings.TrimSpace(count))
			if err != nil {
				log.Errorf("atoi error: %v", err)
			} //vdi-name
			log.Infof("lens: %v", lens)
			for  i := 0; i<lens; i++ {
				out_, err := shell.Exec(fmt.Sprintf("$elog.ReplacementStrings[%v]", i))
				out := strings.TrimSpace(out_)
				log.Infof("i=%v: %v", i, out)
				if err != nil {
					log.Errorf("vdi-name error: %v", err)
				}
			}
	*/

	vdinameout_, err := shell.Exec("$elog.ReplacementStrings[1]")
	if err != nil {
		log.Errorf("vdi-name error: %v", err)
	} //vdi-name
	usernameout_, err := shell.Exec("$elog.ReplacementStrings[5]")
	if err != nil {
		log.Errorf("user-name error: %v", err)
	} //user-name


	logintimeout_, err := shell.Exec("get-date -Date $elog.TimeGenerated -UFormat \"%Y/%m/%d %T\"")
	if err != nil {
		log.Errorf("logintime error: %v", err)
	} //시간
	vdinameout := strings.TrimSpace(vdinameout_)
	usernameout := strings.TrimSpace(usernameout_)
	logintimeout := strings.TrimSpace(logintimeout_)
	//log.Infof("%v, %v, %v, ", vdinameout, usernameout, logintimeout)
	ret := map[string]string{
		"vdi":vdinameout,
		"username":usernameout,
		"time":logintimeout,
	}
	log.Infof("login return: %v", ret)
	return ret
}

func logincheck(c chan map[string]string, c1 chan string, interval int) {
	for true{
		ret := logincheckfnc()
		c<-ret
		time.Sleep(time.Duration(interval) * time.Second)

	}
}
func healthCheck(login chan map[string]string, logout chan map[string]string, c1 chan string, interval int) {
	//var (
	//	li map[string]string
	//	lo map[string]string
	//)
	err := AgentInit()
	if err != nil {
		log.Fatalf("ADinit: %s", err)
	}
	ips := ipcheck()
	for true {
		//li := <- login
		li := logincheckfnc()
		log.Infof("login chan %v", li)
		//lo := <- logout
		lo := logoutcheckfnc()
		log.Infof("logout chan %v", lo)
		err := httpReq(li, lo, ips[0])
		if err != nil {
			log.Errorf("httpReq: %v", err)
			c1 <- err.Error()
			return
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}


type AGENTCONFIG struct {
	WorksServer string `json:"WorksServer"`
	WorksPort   int    `json:"WorksPort"`
	Type        string `json:"Type"`
	UUID        string `json:"UUID"`
	Silent		bool   `json:"Silent"`
	Interval	int	   `json:"Interval"`
}
var Agentconfig = AGENTCONFIG{Silent: false, Interval: 10}

func AgentInit() (err error) {
	data, err := os.Open("conf.json")
	if err != nil {
		log.Fatalf("ADinit: %s", err)
		return err
	}

	byteValue, err := ioutil.ReadAll(data)
	if err != nil {
		log.Fatalf("ADinit: %s", err)
		return err
	}

	err = json.Unmarshal(byteValue, &Agentconfig)
	if err != nil {
		log.Fatalf("ADinit: %s", err)
		return err
	}
	log.Infof("Agentconfig: %v", Agentconfig)
	return nil

}
//user password change
func httpReq(li map[string]string, lo map[string]string, ip string) error {
	setLog()
	client := &http.Client{}
	data := url.Values{}
	data.Set("uuid", Agentconfig.UUID)
	data.Set("type", Agentconfig.Type)

	logindata, _ := json.Marshal(li)
	logoutdata, _ := json.Marshal(lo)

	data.Set("login", string(logindata))
	data.Set("logout", string(logoutdata))
	data.Set("ip", string(ip))


	log.Infof("data: %v", data.Encode())
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%v:%v/api/workspaceAgent/%v", Agentconfig.WorksServer, Agentconfig.WorksPort, Agentconfig.UUID), strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, err := client.Do(req)
	if err != nil {
		log.Errorln(err)
		return err
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Errorf("response cannot close, %v", err)
		}
	}()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(err)
		log.Errorln(respBody)
		return err
	}
	log.Infof("response Status : %v", resp.Status)
	if resp.StatusCode>=400{
		log.Infof("%v", resp.Status)
	} else {
		var responseData map[string]interface{}
		err = json.Unmarshal(respBody, &responseData)
		if err != nil {
			log.Errorln(err)
			log.Errorln(respBody)
			return err
		}
		log.Infof("response Body_parsed : %v", responseData)
	}
	//log.Println("response Headers : ", resp.Header)
	//log.Println("response Body : ", string(respBody))
	//log.Infof("response stderr %v: %v",  responseData["stderr"] != "", responseData["stderr"])
	//if responseData["stderr"] != "" {
	//
	//	errstr:=fmt.Sprint(responseData["stderr"])
	//	err = errors.New(errstr)
	//	log.Infof("stderr %v",  err)
	//	return err
	//}
	//cmd=string(respBody)
	return nil
}