package main

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus"
	"github.com/ycyun/go-powershell"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
)

func main() {
	//SambaDomain := "js.local"
	MyDomain := "js"
	VmName := "testVM"
	SambaIP := "10.1.1.59"
	WorksIP := "10.1.1.59"
	WorksPort := "8083"
	Type := "server"
	InstanceUuid := "3713d4c3-48f7-4b83-b892-d9bb2335d15d"

	payload := "<powershell>\n" +
		"date > \"c:\\agent\\installed.txt\"\n" +
		"$dnsip = \"" + SambaIP + "\"\n" +
		"echo $dnsip >> \"c:\\agent\\installed.txt\"\n" +
		"set-DnsClientServerAddress -InterfaceIndex (Get-NetAdapter |Select-Object InterfaceAlias, interfaceindex).interfaceindex -ServerAddresses " + SambaIP + "\n" +
		"$password = \"Ablecloud1!\" | ConvertTo-SecureString -asPlainText -Force\n" +
		"$username = \"$" + MyDomain + "\\administrator\" \n" +
		"$credential = New-Object System.Management.Automation.PSCredential($username,$password)\n" +
		"echo Rename-Computer >> \"c:\\agent\\installed.txt\"\n" +
		"Rename-Computer -NewName " + VmName + "\n" +
		"echo Add-Computer >> \"c:\\agent\\installed.txt\"\n" +
		//"Add-Computer -DomainName " + SambaDomain + " -Credential $credential -NewName " + VmName + "\n" +
		"echo Add-Computer end>> \"c:\\agent\\installed.txt\"\n" +
		"$conf = '{\"WorksServer\": \"" + WorksIP + "\", \"WorksPort\": " + WorksPort + ", \"Type\": \"" + Type + "\", \"UUID\": \"" + InstanceUuid + "\", \"Domain\":\""+MyDomain+"\", \"HostName\":\""+VmName+"\"}'\n" +
		"echo $conf| Out-File -Encoding ascii \"c:\\agent\\conf.json\"\n" +
		"echo $conf\n" +
		"echo $conf >> \"c:\\agent\\installed.txt\"\n" +
		"C:\\agent\\nssm.exe set \"Ablecloud Works Agent\" start SERVICE_DELAYED_AUTO_START\n" +
		"C:\\agent\\nssm.exe restart \"Ablecloud Works Agent\"\n" +
		"date >> \"c:\\agent\\installed.txt\"\n" +
		"</powershell>"
	fmt.Println(payload)

	shell, err := setupShell()
	if err != nil{
		fmt.Errorf("%v", err)
	}
	out, err := shell.Exec("hostname")
	if err != nil{
		fmt.Errorf("%v", err)
	}
	out = strings.TrimSpace(out)
	fmt.Println(out)
	fmt.Println("DEV")
	fmt.Println(out=="DEV")
}


func setLog(msg ...string) {
	if len(msg) == 0 {
		msg = append(msg, "Starting")
	}
	startlogger := logrus.New()
	startlogger.SetFormatter(&nested.Formatter{
		HideKeys: false,
		//FieldsOrder: []string{"component", "category"},
		CallerFirst: false,
		NoColors: true,
		CustomCallerFormatter: func(f *runtime.Frame) string {

			File := make([]string, 0)
			Line := make([]int, 0)
			Func := make([]string, 0)
			pc := make([]uintptr, 10)
			n := runtime.Callers(11, pc)
			if n == 0 {
				return "[]"
			}
			pc = pc[:10]
			frames := runtime.CallersFrames(pc)
			names := make([]string, 2)
			for {
				frame, more := frames.Next()
				names = append(names, frame.Function)
				if !more {
					break
				}
				File = append(File, frame.File)
				Line = append(Line, frame.Line)
				Func = append(Func, frame.Function)
				//fmt.Println(File, Line, Func)
			}
			return fmt.Sprintf(" [%s:%d][%s()] from [%s:%d][%s()] ", path.Base(File[0]), Line[0], Func[0], path.Base(File[1]), Line[1], Func[1])
			//startlogger.Infof("- more:%v | %s", more, frame.Function)
			//return fmt.Sprintf("start [%s]",names)
		},
	})

	startlogger.SetLevel(logrus.TraceLevel)

	var writers []io.Writer

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_APPEND|os.O_RDWR|os.O_SYNC, 0666)
	if err == nil {
		writers = append(writers, file)
	} else {
		startlogger.Infof("Failed to log to file, using default stderr: %v", err)
	}
	w := io.MultiWriter(writers...)
	startlogger.SetOutput(w)
	startlogger.SetReportCaller(true)
	startlogger.Infof("%v", msg)
}



type PSCMD struct {
	CMD     string `form:"cmd"`
	ARG     string `form:"arg"`
	TIMEOUT int    `form:"timeout"`
}

type APPVAL struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Desc string `json:"desc"`
}

func setupShell() (shell *powershell.Shell, err error) {
	shell, err = powershell.New()
	if err != nil {
		return nil, err
	}
	exec, err := shell.Exec("$pshost = get-host;$pswindow = $pshost.ui.rawui;$newsize = $pswindow.buffersize;$newsize.width = 400;$pswindow.buffersize = $newsize;")
	if err != nil {
		return nil, err
	}
	log.Infof("setup shell: %v",exec)

	return shell, err
}

func PowershellOut2map(stdout string) (output []map[string]string){
	setLog()
	std3:=strings.TrimSpace(stdout)
	std2:=strings.Split(std3, "\n")
	log.Infof("3\n%v\n", std3)
	dict := map[string]string{}
	log.Infof("2\n%v\n", std2[0])
	defer log.Errorf("error\n")
	for _, j := range std2{

		if strings.TrimSpace(j) != "" {
			splited := strings.SplitN(j, ":", 2)
			key := strings.TrimSpace(splited[0])
			val := strings.TrimSpace(splited[1])
			dict[key]=val
			//fmt.Println(dict)
			//fmt.Printf("%v, %v: %v, %v\n", i, j, key, val)
		} else {
			//fmt.Printf("%v, %v\n", i, j)
			output = append(output, dict)
			dict = map[string]string{}
		}
	}
	output = append(output, dict)
	log.Infof("o\n%v\n", output)
	return output
}

func getApps(shell *powershell.Shell) (apps []*APPVAL) {
	setLog()
	stdout, err := shell.Exec("$WScript = New-Object -ComObject WScript.Shell")
	if err != nil {
		panic(err)
	}
	stdout, err = shell.Exec("Get-ChildItem -Path \"C:\\ProgramData\\Microsoft\\Windows\\Start Menu\\Programs\\**\\*.lnk\" | ForEach-Object {$WScript.CreateShortcut($_.FullName).TargetPath} | select-string -pattern exe | sort-object -unique")
	if err != nil {
		panic(err)
	}

	applist := strings.Split(stdout, "\r\n")
	for _, app := range applist {

		if _, err := os.Stat(app); err == nil {
			cmd := fmt.Sprintf("Get-ItemPropertyValue \"%v\" -Name versionInfo | format-list", app)
			stdout, err = shell.Exec(cmd)
			if err != nil {
				panic(err)
			}
			appdetails := strings.Split(strings.TrimSpace(stdout), "\r\n")
			appdetaild := map[string]string{}
			for _, appdetail := range appdetails {
				kv := strings.Split(appdetail, ":")
				key := strings.TrimSpace(kv[0])
				if len(kv) >= 2 {
					val := strings.TrimSpace(kv[1])
					appdetaild[key] = val
				} else {
					appdetaild[key] = ""
				}
			}
			desc := app
			name := app
			val, exests := appdetaild["FileDescription"]
			if exests {
				desc = val
			}

			val, exests = appdetaild["ProductName"]
			if exests {
				name = val
			}

			fmt.Println(stdout)
			a := &APPVAL{
				Name: name,
				Path: app,
				Desc: desc,
			}
			apps = append(apps, a)

		} else if os.IsNotExist(err) {
			fmt.Println("notExist")
			continue
		} else {
			fmt.Println("other")
			continue
		}
	}
	return apps
}
