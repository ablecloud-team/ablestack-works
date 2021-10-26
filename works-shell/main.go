package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	WorksDB = "works-db"
	WorksApi = "works-api"
	WorksSamba = "works-samba"
	Guacd = "guacd"
	Guacamole = "guacamole"
)
type Config struct {
	ClusterDefaultName       string `json:"cluster.default.name"`
	DcDefaultUrl             string `json:"dc.default.url"`
	MoldDefaultApiKey        string `json:"mold.default.api.key"`
	MoldDefaultDomainAccount string `json:"mold.default.domain.account"`
	MoldDefaultDomainId      string `json:"mold.default.domain.id"`
	MoldDefaultPort          string `json:"mold.default.port"`
	MoldDefaultProtocol      string `json:"mold.default.protocol"`
	MoldDefaultSecretKey     string `json:"mold.default.secret.key"`
	MoldDefaultUrl           string `json:"mold.default.url"`
	MoldNetworkUuid          string `json:"mold.network.uuid"`
	MoldZoneId               string `json:"mold.zone.id"`
	SambaDefaultDomain       string `json:"samba.default.domain"`
	SambaDefaultUrl          string `json:"samba.default.url"`
	WorksDefaultUrl          string `json:"works.default.url"`
	Bootstrapped          string `json:"bootstrapped"`
}

func containerChecker(containerName string) bool {
	cmdPodmanPs, _ := exec.Command("podman", "ps").Output()
	containerExecBool := strings.Contains(string(cmdPodmanPs), containerName)
	log.Infof("[%v] container exec [%v]", containerName, containerExecBool)

	return containerExecBool
}

func main () {
	var writers []io.Writer
	writers = append(writers, os.Stdout)
	file, err := os.OpenFile("/works/works_bootstrap.log", os.O_CREATE|os.O_RDWR|os.O_SYNC, 0666)
	if err == nil {
		writers = append(writers, file)
	} else {
		log.Infof("Failed to log to file, using default stderr: %v", err)
	}
	cmdPwd, _ := exec.Command("pwd").Output()
	log.Infof("pwd [%v]", string(cmdPwd))
	w := io.MultiWriter(writers...)
	log.SetOutput(w)
	log.SetReportCaller(true)
	// bootstrap Log 설정 끝

	log.WithFields(log.Fields{
		"serverVersion": "Bronto-Works",
	}).Infof("=========Bootstrap exec=========")
	configMap := map[string]string{}
	configStruct := Config{}
	for  {
		configJson, err1 := os.Open("/works/works_userdata.json")
		defer file.Close()
		if err1 != nil {
			log.Errorf("works_userdata.json file cannot be read. \n Wait 10 seconds and check again.[%v]", err1)
			time.Sleep(10 * time.Second)
		} else {
			log.Infof("works_userdata.json file is fine. [%v]", configJson)
			log.Infof("works_userdata.json data read START.")
			decoder := json.NewDecoder(configJson)
			errConfigMap := decoder.Decode(&configMap)
			configJson2, _ := json.Marshal(configMap)
			errConfigStruct := json.Unmarshal(configJson2, &configStruct)
			if errConfigMap != nil && errConfigStruct != nil {
				log.Errorf("errConfigMap [%v], errConfigStruct [%v]", errConfigMap, errConfigStruct)
			}
			log.Infof("configMap [%v] \nconfigStruct [%v]", configMap, configStruct)
			log.Infof("works_userdata.json data read END.")
			break
		}
	}


	if configStruct.Bootstrapped != "true" {
		containers1 := []string{
			WorksDB, Guacd,
		}
		for _, container := range containers1{
			for {
				if !containerChecker(container) {
					cmdWorksDBStart, _ := exec.Command("podman", "start", container).Output()
					log.Infof("========== podman start %v ========== \n %v", container, string(cmdWorksDBStart))
					time.Sleep(10 * time.Second)
				} else {
					break
				}
			}
		}
		cmdSamba, errSamba := exec.Command(
			"podman", "run",
			"--name", "works-samba",
			"-d", "-t",
			"--net", "host",
			"localhost/works-ad:v3",
			"config", "8.8.8.8",
			configStruct.SambaDefaultDomain,
			"Ablecloud1!", "ad",
		).Output()
		log.Infof("SAMBA run exec [%v], error [%v]",string(cmdSamba), errSamba)
		time.Sleep(10 * time.Second)

		cmdGuacamole, errGuacamole := exec.Command(
			"podman", "run",
			"--name", "guacamole",
			"--net", "works",
			"--ip", "10.88.2.10",
			"-v", "share:/share",
			"-e", "LDAP_HOSTNAME="+configStruct.SambaDefaultUrl,
			"-e", "LDAP_USER_BASE_DN=CN=Users,DC="+configStruct.SambaDefaultDomain+",DC=local",
			"-e", "LDAP_USERNAME_ATTRIBUTE=cn",
			"-e", "LDAP_CONFIG_BASE_DN=CN=Users,DC="+configStruct.SambaDefaultDomain+",DC=local",
			"-e", "GUACD_HOSTNAME='10.88.2.14'",
			"-d", "-p", "8080:8080",
			"guacamole:v0.1",
		).Output()
		log.Infof("guacamole run exec [%v], error [%v]",string(cmdGuacamole), errGuacamole)
		time.Sleep(10 * time.Second)

		for key,value := range(configMap) {
			log.Infof("key [%v], value [%v]", key, value)
			cmdDB, errDB := exec.Command(
				"podman", "exec",
				"-it", "works-db",
				"mysql", "-uroot",
				"-pAblecloud1!", "works",
				"-e",
				"UPDATE configuration SET value='"+value+"' WHERE name='"+key+"'",
			).Output()
			log.Infof("mysql update exec [%v], error [%v]",string(cmdDB), errDB)
			fmt.Println(string(cmdDB))
			time.Sleep(2 * time.Second)
		}
		containers2 := []string{
			WorksApi,
		}
		for _, container := range containers2{
			for {
				if !containerChecker(container) {
					cmdWorksDBStart, _ := exec.Command("podman", "start", container).Output()
					log.Infof("========== podman start %v ========== \n %v", container, string(cmdWorksDBStart))
					time.Sleep(10 * time.Second)
				} else {
					break
				}
			}
		}
		configStruct.Bootstrapped = "true"
		jsonDoc, _ := json.Marshal(configStruct)
		errWriteFile := ioutil.WriteFile("/works/works_userdata.json", jsonDoc, os.FileMode(0755))
		if errWriteFile != nil {
			log.Errorf("works_userdata.json write file error [%v]",errWriteFile)
		}
	} else {
		containers := []string{
			WorksSamba, WorksDB, Guacd, Guacamole, WorksApi,
		}
		for _, container := range containers{
			for {
				if !containerChecker(container) {
					cmdWorksDBStart, _ := exec.Command("podman", "start", container).Output()
					log.Infof("========== podman start %v ========== \n %v", container, string(cmdWorksDBStart))
				} else {
					break
				}
			}
		}
	}
}

//func main() {
//		// bootstrap Log 설정 시작
//	var writers []io.Writer
//	writers = append(writers, os.Stdout)
//	file, err := os.OpenFile("/works/works_bootstrap.log", os.O_CREATE|os.O_RDWR|os.O_SYNC, 0666)
//	if err == nil {
//		writers = append(writers, file)
//	} else {
//		log.Infof("Failed to log to file, using default stderr: %v", err)
//	}
//	cmdPwd, _ := exec.Command("pwd").Output()
//	log.Infof("pwd [%v]", cmdPwd)
//	w := io.MultiWriter(writers...)
//	log.SetOutput(w)
//	log.SetReportCaller(true)
//	// bootstrap Log 설정 끝
//
//	log.WithFields(log.Fields{
//		"serverVersion": "Bronto-Works",
//	}).Infof("=========Bootstrap exec=========")
//	for {
//		time.Sleep(30 * time.Second)
//		cmd1, _ := exec.Command("cloud-init", "status").Output()
//		log.Infof("cloud-init status [%v]",string(cmd1))
//		log.Infof("[%v]",strings.Contains(string(cmd1), "works-db"))
//		if !strings.Contains(string(cmd1), "done") {
//			log.Infof("cloud-init status [%v] ",cmd1)
//		} else {
//			log.Infof("cloud-init status [%v] ",cmd1)
//			break
//		}
//	}
//	for {
//		//cmdCrontab, _ := exec.Command("crontab", "-r").Output()
//		//log.Infof("crontab remove [%v] ",cmdCrontab)
//
//		configJson, err1 := os.Open("/works/works_userdata.json")
//		defer file.Close()
//		if err1 != nil {
//			log.Errorf("works_userdata.json 파일을 읽을수가 없습니다. 10초 대기 후 다시 확인합니다.[%v]", err1)
//			time.Sleep(10 * time.Second)
//		} else {
//			log.Infof("works_userdata.json 파일이 정상적으로 있습니다. [%v]", configJson)
//			log.Infof("works_userdata.json 데이터 Read START.")
//			decoder := json.NewDecoder(configJson)
//			configMap := map[string]string{}
//			configStruct := Config{}
//			errConfigMap := decoder.Decode(&configMap)
//			configJson2,_ := json.Marshal(configMap)
//			errConfigStruct := json.Unmarshal(configJson2,&configStruct)
//			if errConfigMap != nil && errConfigStruct != nil {
//				log.Errorf("errConfigMap [%v], errConfigStruct [%v]", errConfigMap, errConfigStruct)
//			}
//			log.Infof("configMap [%v], configStruct [%v]",configMap, configStruct)
//			log.Infof("works_userdata.json 데이터 Read END.")
//
//			log.Infof("Guacd Setting START. [%v]", configStruct)
//			log.Infof("Guacd Container start.")
//			cmdGuacd, errGuacd := exec.Command("podman", "start", "guacd").Output()
//			log.Infof("guacd run exec [%v], error [%v]",string(cmdGuacd), errGuacd)
//			fmt.Println(string(cmdGuacd))
//			log.Infof("Guacad Setting END. [%v]", configJson)
//			time.Sleep(60 * time.Second)
//
//			log.Infof("Guacamole Setting START. [%v]", configStruct)
//			log.Infof("Guacamole params samba url [%v], domain [%v]", configStruct.SambaDefaultUrl, configStruct.SambaDefaultDomain)
//			cmdGuacamole, errGuacamole := exec.Command(
//				"podman", "run",
//				"--name", "guacamole",
//				"--net", "works",
//				"--ip", "10.88.2.10",
//				"-v", "share:/share",
//				"-e", "LDAP_HOSTNAME="+configStruct.SambaDefaultUrl,
//				"-e", "LDAP_USER_BASE_DN=CN=Users,DC="+configStruct.SambaDefaultDomain+",DC=local",
//				"-e", "LDAP_USERNAME_ATTRIBUTE=cn",
//				"-e", "LDAP_CONFIG_BASE_DN=CN=Users,DC="+configStruct.SambaDefaultDomain+",DC=local",
//				"-e", "GUACD_HOSTNAME='10.88.2.14'",
//				"-d", "-p", "8080:8080",
//				"guacamole/guacamole:ablestack",
//			).Output()
//			log.Infof("guacamole run exec [%v], error [%v]",string(cmdGuacamole), errGuacamole)
//			fmt.Println(string(cmdGuacamole))
//			log.Infof("Guacamole Setting END. [%v]", configJson)
//			time.Sleep(60 * time.Second)
//
//			log.Infof("SAMBA Setting START. [%v]", configStruct)
//			log.Infof("SAMBA params samba [%v]", configStruct.SambaDefaultDomain)
//			cmdSamba, errSamba := exec.Command(
//				"podman", "run",
//				"--name", "works-samba",
//				"-d", "-t",
//				"--net", "host",
//				"localhost/works-ad:v3",
//				"config", "8.8.8.8",
//				configStruct.SambaDefaultDomain,
//				"Ablecloud1!", "ad",
//			).Output()
//			log.Infof("SAMBA run exec [%v], error [%v]",string(cmdSamba), errSamba)
//			fmt.Println(string(cmdSamba))
//			log.Infof("SAMBA Setting END. [%v]", configJson)
//			time.Sleep(60 * time.Second)
//
//			log.Infof("DB Setting START. [%v]", configStruct)
//			log.Infof("DB params samba [%v]", configStruct)
//			log.Infof("Guacd Container start.")
//			cmdMysql, errMysql := exec.Command("podman", "start", "works-db").Output()
//			time.Sleep(60 * time.Second)
//			log.Infof("60 second Waiting")
//			log.Infof("works-db run exec [%v], error [%v]",string(cmdMysql), errMysql)
//			fmt.Println(string(cmdMysql))
//			log.Infof("works-db update start")
//			for key,value := range(configMap) {
//				log.Infof("key [%v], value [%v]", key, value)
//				cmdDB, errDB := exec.Command(
//					"podman", "exec",
//					"-it", "works-db",
//					"mysql", "-uroot",
//					"-pAblecloud1!", "works",
//					"-e",
//					"UPDATE configuration SET value='"+value+"' WHERE name='"+key+"'",
//				).Output()
//				log.Infof("guacamole run exec [%v], error [%v]",string(cmdDB), errDB)
//				fmt.Println(string(cmdDB))
//			}
//			log.Infof("works-db update END")
//			log.Infof("DB Setting END.")
//			time.Sleep(60 * time.Second)
//
//			log.Infof("works-api Setting START. [%v]", configStruct)
//			log.Infof("works-api Container start.")
//			cmdWorksApi, errWorksApi := exec.Command("podman", "start", "works-api").Output()
//			log.Infof("works-api run exec [%v], error [%v]",string(cmdWorksApi), errWorksApi)
//			fmt.Println(string(cmdWorksApi))
//			log.Infof("works-api Setting END. [%v]", configJson)
//			time.Sleep(60 * time.Second)
//
//			log.Infof("works-ui Setting START. [%v]", configStruct)
//			log.Infof("works-ui Container start.")
//			cmdWorksUi, errWorksUi := exec.Command("podman", "start", "works-api").Output()
//			log.Infof("works-ui run exec [%v], error [%v]",string(cmdWorksUi), errWorksUi)
//			fmt.Println(string(cmdWorksUi))
//			log.Infof("works-ui Setting END. [%v]", configJson)
//			time.Sleep(60 * time.Second)
//
//			break
//		}
//	}
//}



// guacd     IP : 10.88.2.14
// gucamole  IP : 10.88.2.10 , port forwarding : 8080 -> 8080
// SAMBA     IP : 가상머신 IP
// mysql     IP : 10.88.2.12 , port forwarding : 3306 -> 3306
// works-api IP : 10.88.2.13 , port forwarding : 8080 -> 8082
// works-ui  IP : 10.88.2.11 , port forwarding : 8080 -> 8081

// podman run --name works-db --net works --ip 10.88.2.12 -p 3306:3306 -v works-db:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=Ablecloud1! -d mysql:latest