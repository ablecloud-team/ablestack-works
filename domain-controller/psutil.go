package main

import (
	"fmt"
	"github.com/ycyun/go-powershell"
	"os"
	"strings"
)

type PSCMD struct {
	CMD     string `form:"cmd" example:"echo"`
	ARG     string `form:"arg" example:"Hello World!"`
	TIMEOUT int    `form:"timeout" example:"3"`
}

type APPVAL struct {
	Name string `json:"name" example:"notepad.exe"`
	Path string `json:"path" example:"C:\Windows\System32\notepad.exe"`
	Desc string `json:"desc" example:"메모장"`
}

func setupShell() (shell *powershell.Shell, err error) {
	shell, err = powershell.New()
	shell.Exec("$pshost = get-host;$pswindow = $pshost.ui.rawui;$newsize = $pswindow.buffersize;$newsize.width = 400;$pswindow.buffersize = $newsize;")
	if err != nil {
		panic(err)
	}

	return shell, err
}

func PowershellOut2map(stdout string) (output []map[string]string) {
	setLog()
	std3 := strings.TrimSpace(stdout)
	std2 := strings.Split(std3, "\n")
	log.Infof("3\n%v\n", std3)
	dict := map[string]string{}
	log.Infof("2\n%v\n", std2[0])
	defer log.Errorf("error\n")
	for _, j := range std2 {

		if strings.TrimSpace(j) != "" {
			splited := strings.SplitN(j, ":", 2)
			key := strings.TrimSpace(splited[0])
			val := strings.TrimSpace(splited[1])
			dict[key] = val
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
