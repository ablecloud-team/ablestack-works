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

	shell, err := setupShell()
	if err != nil {
		fmt.Errorf("%v", err)
	}
	hostname, err := shell.Exec("hostname")
	if err != nil {
		fmt.Errorf("%v", err)
	}
	hostname = strings.TrimSpace(hostname)
	log.Println(hostname)
	log.Println(Agentconfig.HostName)

	if !strings.EqualFold(hostname, Agentconfig.HostName) && Agentconfig.HostName != "" {
		cmd := fmt.Sprintf("Rename-Computer -NewName %v", Agentconfig.HostName)
		log.Errorf("cmd: %v", cmd)
		out, err := shell.Exec(cmd)
		log.Errorf("out: %v", out)
		log.Errorf("err: %v", err)
		cmd = fmt.Sprintf("shutdown /r /t 10")
		log.Errorf("cmd: %v", cmd)
		out, err = shell.Exec(cmd)
		log.Errorf("out: %v", out)
		log.Errorf("err: %v", err)
	}
	time.Sleep(10 * time.Second)
	interval := Agentconfig.Interval

	//logincheckfnc2()
	//return
	c1 := make(chan string, 1)
	//loginchan := make(chan map[string]string, 1)
	//logoutchan := make(chan map[string]string, 1)
	go healthCheck(c1, interval)
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
			v1.GET("/cmd/", exeShellHandler)
			v1.GET("/app", appListHandler)
			v1.PUT("/ad/:domain", adjoinHandler)
			v1.PATCH("/", bootstrapHandler)
			v1.GET("/ad", adStatusHandler)
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

func ipcheck() []string {
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
	for _, ip := range stdouts {
		if !strings.Contains(ip, "127.0.0.1") {
			ips = append(ips, ip)
		}
	}
	log.Infof("%v", stdouts)

	return ips
}

type shellReturnModel struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

type errorModel struct {
	Msg    string `json:"msg"`
	Target string `json:"target"`
}

// exeShellHandler godoc
// @Summary powershell 명령 처리기
// @Description powershell 명령 처리기
// @Accept  multipart/form-data
// @Produce  application/json
// @Param cmd query string true "명령어"
// @Param arg query string false "인자"
// @Param timeout query int false "시간제한, 기본값"
// @Success 200 {object} shellReturnModel "명령 성공"
// @Failure 401 {object} errorModel "명령 실패"
// @Failure default {object} string
// @Router /cmd/ [get]
func exeShellHandler(c *gin.Context) {

	var (
		pscmd  PSCMD
		stdout string
		stderr string
	)
	if err := c.Bind(&pscmd); err != nil {
		c.JSON(http.StatusBadRequest, errorModel{Msg: "입력 오류", Target: err.Error()})
		return
	}
	if pscmd.CMD == "" {
		c.JSON(http.StatusBadRequest, errorModel{Msg: "Command not found", Target: "powershell"})
		return
	}
	if pscmd.TIMEOUT == 0 {
		pscmd.TIMEOUT = 3
	}

	c1 := make(chan []string, 1)
	go func() {

		shell, _ := setupShell()
		stdout, err := shell.Exec(fmt.Sprintf("%v %v", pscmd.CMD, pscmd.ARG))

		if err == nil {
			c1 <- []string{stdout, ""}
		} else {
			c1 <- []string{stdout, err.Error()}
		}

	}()

	select {
	case res := <-c1:
		stdout = res[0]
		stderr = res[1]
	case <-time.After(time.Duration(pscmd.TIMEOUT) * time.Second):
		stdout = "Timeout Reached"
		stderr = "Timeout Reached"
	}

	log.Infof("cmd: %v, arg: %v, stdout: %v, stderr: %v", pscmd.CMD, pscmd.ARG, stdout, stderr)
	c.JSON(http.StatusOK, shellReturnModel{Stdout: stdout, Stderr: stderr})
	return
}

func logoutcheckfnc() map[string]string {
	shell, err := setupShell()
	stdout, err := shell.Exec("$elog=(Get-EventLog security -source microsoft-windows-security-auditing  | where {($_.instanceID -eq 4634)} | select TimeGenerated,ReplacementStrings,Message -First 1)")

	log.Infof("stdout: %v", stdout)
	if err != nil {
		log.Errorf("logoutcheck error: %v", err)
	}
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
		"vdi":      vdinameout,
		"username": usernameout,
		"time":     logintimeout,
	}
	log.Infof("logout return: %v", ret)
	return ret
}
func logoutcheck(c chan map[string]string, c1 chan string, interval int) {
	for true {
		ret := logoutcheckfnc()
		c <- ret
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
func logincheckfnc2() []map[string]string {
	shell, err := setupShell()
	stdout, err := shell.Exec("Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser")
	stdout, err = shell.Exec("./get-session.ps1")
	//log.Infof("stdout: %v", stdout)
	if err != nil {
		log.Errorf("stderr: %v", err)
	}
	var ret []map[string]string
	err = json.Unmarshal([]byte(stdout), &ret)
	if err != nil {
		log.Errorf("stderr: %v", err)
	}
	//log.Errorf("stderr: %v",err)
	//log.Infof("ret: %v", ret)
	return ret
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
		"vdi":      vdinameout,
		"username": usernameout,
		"time":     logintimeout,
	}
	log.Infof("login return: %v", ret)
	return ret
}

func logincheck(c chan []map[string]string, c1 chan string, interval int) {
	for true {
		ret := logincheckfnc2()
		c <- ret
		time.Sleep(time.Duration(interval) * time.Second)

	}
}
func healthCheck(c1 chan string, interval int) {
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
		li := logincheckfnc2()
		log.Infof("login chan %v", li)
		//lo := <- logout
		//lo := logoutcheckfnc()
		//log.Infof("logout chan %v", lo)
		err := httpReq(li, ips)
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
	Silent      bool   `json:"Silent"`
	Interval    int    `json:"Interval"`
	Status      string `json:"Status"`
	HostName    string `json:"HostName"`
	Domain      string `json:"Domain"`
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

func AgentSave() (err error) {
	data, err := os.Create("conf.json")
	if err != nil {
		log.Fatalf("agentsave: %s", err)
		return err
	}
	byteValue, err := json.MarshalIndent(Agentconfig, "", " ")
	if err != nil {
		log.Fatalf("ADinit: %s", err)
		return err
	}
	log.Errorf("Save %s", byteValue)
	_, err = data.Write(byteValue)
	if err != nil {
		log.Fatalf("ADinit: %s", err)
		return err
	}
	return nil

}

//user password change
func httpReq(li []map[string]string, ips []string) error {
	setLog()
	client := &http.Client{}
	data := url.Values{}
	data.Set("uuid", Agentconfig.UUID)
	data.Set("type", Agentconfig.Type)

	logindata, _ := json.Marshal(li)
	ip, _ := json.Marshal(ips)
	//logoutdata, _ := json.Marshal(lo)

	data.Set("login", string(logindata))
	//data.Set("logout", string(logoutdata))
	data.Set("ip", string(ip))

	//log.Infof("data: %v", data.Encode())
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%v:%v/api/workspaceAgent/%v", Agentconfig.WorksServer, Agentconfig.WorksPort, Agentconfig.UUID), strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	//log.Infof("req: %v", req)
	//log.Infof("Body: %v", req.Body)
	//log.Infof("Header: %v", req.Header)
	//log.Infof("URL: %v", req.URL)
	//log.Infof("Form: %v", req.Form)
	//log.Infof("PostForm: %v", req.PostForm)
	//log.Infof("RequestURI: %v", req.RequestURI)
	br, _ := req.GetBody()
	insert := make([]byte, 2048)
	_, _ = br.Read(insert)
	//log.Infof("BodyRead: %v", string(insert))

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
	if resp.StatusCode >= 400 {
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

func adjoinHandler(c *gin.Context) {
	setLog()
	domain := c.Param("domain")
	domain = Agentconfig.Domain
	shell, err := setupShell()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, errorModel{Msg: err.Error(), Target: "getComputer"})
		return
	}
	cmd := fmt.Sprintf("$username = \"$%v\\administrator\"; ", domain) +
		"$password = \"Ablecloud1!\" | ConvertTo-SecureString -asPlainText -Force; " +
		"$credential = New-Object System.Management.Automation.PSCredential($username,$password); " +
		fmt.Sprintf("Add-Computer -DomainName %v -Credential $credential", domain)
	log.Debugln(cmd)
	output, err := shell.Exec(cmd)
	log.Debugf("output: %v", output)
	log.Debugf("error: %v", err)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, errorModel{Msg: err.Error(), Target: "getComputer"})
		return
	}
	cmd = "shutdown /r /t 10"
	log.Debugln(cmd)
	output, err = shell.Exec(cmd)
	log.Debugf("output: %v", output)
	log.Debugf("error: %v", err)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, errorModel{Msg: err.Error(), Target: "getComputer"})
		return
	}
	Agentconfig.Status = "Joining"
	_ = AgentSave()
	c.JSON(http.StatusOK, output)
	return
}

func bootstrapHandler(c *gin.Context) {
	setLog()

	//computername := c.Param("computername")
	domain := c.PostForm("domain")
	shell, err := setupShell()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, errorModel{Msg: err.Error(), Target: "getComputer"})
		return
	}
	service1 := fmt.Sprintf("c:\\agent\\nssm.exe set \"Ablecloud Works Agent\" objectName %v\\administrator Ablecloud1!", domain)
	output1, err1 := shell.Exec(service1)
	if err1 != nil {

		log.Errorf("err1: %v", service1)
		log.Errorf("err1: %v", err1)
		c.JSON(http.StatusServiceUnavailable, errorModel{Msg: err1.Error(), Target: "getComputer"})
		return
	}
	service2 := fmt.Sprintf("gpupdate /force")
	output2, err2 := shell.Exec(service2)
	if err2 != nil {
		log.Errorf("err1: %v", service2)
		log.Errorf("err2: %v", err2)
		c.JSON(http.StatusServiceUnavailable, errorModel{Msg: err2.Error(), Target: "getComputer"})
		return
	}
	c.JSON(http.StatusOK, map[string]string{"output1": output1, "output2": output2})
	return
}

func adStatusHandler(c *gin.Context) {
	setLog()
	shell, err := setupShell()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, errorModel{Msg: err.Error(), Target: "getComputer"})
		return
	}
	output, err := shell.Exec("get-computerinfo -property csdomain | format-list")
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, errorModel{Msg: err.Error(), Target: "getComputer"})
		return
	}
	domain := strings.TrimSpace(strings.Split(strings.TrimSpace(output), ":")[1])

	if strings.EqualFold(domain, "WORKGROUP") && Agentconfig.Status != "Joining" {
		Agentconfig.Status = "Pending"
		_ = AgentSave()
		c.JSON(http.StatusOK,
			map[string]string{
				"status": Agentconfig.Status,
				"next":   "PUT <vdi>/ad/:domain/",
			})
		return
	} else {
		output, err := shell.Exec("get-computerinfo -property csname | format-list")
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, errorModel{Msg: err.Error(), Target: "getComputer"})
			return
		}
		comname := strings.TrimSpace(strings.Split(strings.TrimSpace(output), ":")[1])
		output, err = shell.Exec(fmt.Sprintf("gpresult /scope computer /r | select-string -pattern cn=%v", comname))
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, errorModel{Msg: err.Error(), Target: "getComputer"})
			return
		}
		if strings.Contains(strings.TrimSpace(output), "OU=") {
			Agentconfig.Status = "Joined"
			_ = AgentSave()
		}
		output, err = shell.Exec("gpresult /scope computer /r | select-string -pattern remote")
		if err != nil {
			c.JSON(http.StatusServiceUnavailable, errorModel{Msg: err.Error(), Target: "getComputer"})
			return
		}
		if strings.Contains(strings.TrimSpace(output), "remotefx") {
			Agentconfig.Status = "Ready"
			_ = AgentSave()
		}
		if Agentconfig.Status == "Joining" {
			c.JSON(http.StatusOK,
				map[string]string{
					"status": Agentconfig.Status,
					"next":   "PUT <dc>/computer/:computername/:groupname",
				})
			return
		} else if Agentconfig.Status == "Joined" {
			c.JSON(http.StatusOK,
				map[string]string{
					"status": Agentconfig.Status,
					"next":   "GET <vdi>/cmd/?timeout=60&cmd=gpupdate /force",
				})
			return
		} else if Agentconfig.Status == "Ready" {
			c.JSON(http.StatusOK,
				map[string]string{
					"status": Agentconfig.Status,
					"next":   "Ready to use",
				})
			return
		}
	}
	c.JSON(http.StatusOK, output)
	return
}
