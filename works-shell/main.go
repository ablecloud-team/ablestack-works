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
	WorksDB    = "works-db"
	WorksApi   = "works-api"
	WorksUi   = "works-ui"
	WorksSamba = "works-samba"
	Guacd      = "guacd"
	Guacamole  = "guacamole"
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
	Bootstrapped             string `json:"bootstrapped"`
}

func containerChecker(containerName string) bool {
	cmdPodmanPs, _ := exec.Command("podman", "ps").Output()
	containerExecBool := strings.Contains(string(cmdPodmanPs), containerName)
	log.Infof("[%v] container exec [%v]", containerName, containerExecBool)

	return containerExecBool
}

func main() {
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
	for {
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
			WorksDB, Guacd, WorksUi,
		}
		for _, container := range containers1 {
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
		log.Infof("SAMBA run exec [%v], error [%v]", string(cmdSamba), errSamba)
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
			"-e", "GUACD_HOSTNAME=10.88.2.14",
			"-d", "-p", "8080:8080",
			"guacamole:v0.1",
		).Output()
		log.Infof("guacamole run exec [%v], error [%v]", string(cmdGuacamole), errGuacamole)
		time.Sleep(10 * time.Second)

		for key, value := range configMap {
			log.Infof("key [%v], value [%v]", key, value)
			cmdDB, errDB := exec.Command(
				"podman", "exec",
				"-it", "works-db",
				"mysql", "-uroot",
				"-pAblecloud1!", "works",
				"-e",
				"UPDATE configuration SET value='"+value+"' WHERE name='"+key+"'",
			).Output()
			log.Infof("mysql update exec [%v], error [%v]", string(cmdDB), errDB)
			fmt.Println(string(cmdDB))
			time.Sleep(2 * time.Second)
		}
		containers2 := []string{
			WorksApi,
		}
		for _, container := range containers2 {
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
			log.Errorf("works_userdata.json write file error [%v]", errWriteFile)
		}
	} else {
		containers := []string{
			WorksSamba, WorksDB, Guacd, Guacamole, WorksApi, WorksUi,
		}
		for _, container := range containers {
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

//podman run --name guacamole --net works --ip 10.88.2.10 -v share:/share -e LDAP_HOSTNAME=10.1.1.155 -e LDAP_USER_BASE_DN=CN=Users,DC=cl,DC=local -e LDAP_USERNAME_ATTRIBUTE=cn -e LDAP_CONFIG_BASE_DN=CN=Users,DC=cl,DC=local -e GUACD_HOSTNAME=10.88.2.14 -d -p 8080:8080 guacamole:v0.1

// guacd     IP : 10.88.2.14
// gucamole  IP : 10.88.2.10 , port forwarding : 8080 -> 8080
// SAMBA     IP : 가상머신 IP
// mysql     IP : 10.88.2.12 , port forwarding : 3306 -> 3306
// works-api IP : 10.88.2.13 , port forwarding : 8080 -> 8082
// works-ui  IP : 10.88.2.11 , port forwarding : 8080 -> 8081

// podman run --name works-db --net works --ip 10.88.2.12 -p 3306:3306 -v works-db:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=Ablecloud1! -d mysql:latest
