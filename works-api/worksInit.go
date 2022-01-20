package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"os"
)

type settingInfo struct {
	Database struct {
		TYPE string `json:"type"`
		User struct {
			ID       string `json:"id"`
			Password string `json:"password"`
		} `json:"user"`
		Host struct {
			Address  string `json:"address"`
			Port     string `json:"port"`
			Protocol string `json:"protocol"`
		} `json:"host"`
		DB string `json:"db"`
	} `json:"database"`
}

//WorksConfig
var WorksConfig settingInfo

func getWorksInfo() settingInfo {
	file, err := os.Open("properties.json")
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&WorksConfig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(WorksConfig)
	return WorksConfig
}

func DBSetting() {
	os.Setenv("MysqlType", getWorksInfo().Database.TYPE)
	os.Setenv("MysqlUser", getWorksInfo().Database.User.ID)
	os.Setenv("MysqlPassword", getWorksInfo().Database.User.Password)
	os.Setenv("MysqlProtocol", getWorksInfo().Database.Host.Protocol)
	os.Setenv("MysqlServerIp", getWorksInfo().Database.Host.Address)
	os.Setenv("MysqlServerPort", getWorksInfo().Database.Host.Port)
	os.Setenv("MysqlDb", getWorksInfo().Database.DB)
	os.Setenv("DbInfo", getWorksInfo().Database.User.ID+":"+getWorksInfo().Database.User.Password+"@"+getWorksInfo().Database.Host.Protocol+"("+getWorksInfo().Database.Host.Address+":"+getWorksInfo().Database.Host.Port+")/"+getWorksInfo().Database.DB)
}

func Setup() {
	i := 0
	for i == 1 {
		db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
		if err != nil {
			log.WithFields(logrus.Fields{
				"worksInit": "Setup",
			}).Errorf("DB connect error[%v]", err)
		} else {
			DBSetting()        //DB 접속정보 셋팅
			MoldSetting()      //Mold 정보 셋팅
			DCSetting()        //DC 정보 셋팅
			WorksSetting()     //Works-API 정보 셋팅
			SambaSetting()     //SAMBA 정보 셋팅
			GuacamoleSetting() //guacamole 정보 셋팅
			i = 1
		}
		defer db.Close()
	}
}

func logSetting() {
	log.SetFormatter(&nested.Formatter{
		HideKeys:    false,
		CallerFirst: false,
	})
	log.SetReportCaller(true)
	log.SetLevel(logrus.DebugLevel)
}

func MoldSetting() {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "MoldSetting",
		}).Errorf("DB connect error[%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"worksInit": "MoldSetting",
	}).Infof("DB connect success")

	rows, err := db.Query("SELECT * FROM configuration WHERE name LIKE 'mold.default%'")
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "MoldSetting",
		}).Errorf("worksInit Mold Setting Query Failed [%v]", err)
	}
	defer rows.Close()

	result, err := rowsToString(rows)
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "MoldSetting",
		}).Errorf("Row to String conversion error [%v]", err)
	}

	jsonUnmarshal := []Configuration{}
	err = json.Unmarshal([]byte(result), &jsonUnmarshal)
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "MoldSetting",
		}).Errorf("String to JSON conversion error [%v]", err)
	}

	moldValue := map[string]interface{}{}
	for _, v := range jsonUnmarshal {
		moldValue[v.Name] = v.Value
	}
	protocol := moldValue["mold.default.protocol"].(string)
	url := moldValue["mold.default.url"].(string)
	port := moldValue["mold.default.port"].(string)
	if port == "80" {
		port = ""
	} else {
		port = ":" + port
	}
	postfix := moldValue["mold.default.url.postfix"].(string)
	apiKey := moldValue["mold.default.api.key"].(string)
	secretKey := moldValue["mold.default.secret.key"].(string)
	domainId := moldValue["mold.default.domain.id"].(string)
	log.WithFields(logrus.Fields{
		"worksInit": "MoldSetting",
	}).Infof("Set Mold URL [%v%v%v%v]", protocol, url, port, postfix)
	os.Setenv("MoldUrl", protocol+url+port+postfix)
	os.Setenv("MoldApiKey", apiKey)
	os.Setenv("MoldSecretKey", secretKey)
	os.Setenv("MoldDomainId", domainId)
}

func DCSetting() {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "DCSetting",
		}).Errorf("DB connect error[%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"worksInit": "DCSetting",
	}).Infof("DB connect success")

	rows, err := db.Query("SELECT * FROM configuration WHERE name LIKE 'dc.default%'")
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "DCSetting",
		}).Errorf("worksInit DC Setting Query Failed [%v]", err)
	}
	defer rows.Close()

	result, err := rowsToString(rows)
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "DCSetting",
		}).Errorf("Row to String conversion error [%v]", err)
	}

	jsonUnmarshal := []Configuration{}
	err = json.Unmarshal([]byte(result), &jsonUnmarshal)
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "DCSetting",
		}).Errorf("String to JSON conversion error [%v]", err)
	}

	dcValue := map[string]interface{}{}
	for _, v := range jsonUnmarshal {
		dcValue[v.Name] = v.Value
	}

	protocol := dcValue["dc.default.protocol"].(string)
	url := dcValue["dc.default.url"].(string)
	port := dcValue["dc.default.port"].(string)
	if port == "80" {
		port = ""
	} else {
		port = ":" + port
	}
	postfix := dcValue["dc.default.url.postfix"].(string)
	log.WithFields(logrus.Fields{
		"worksInit": "DCSetting",
	}).Infof("Set Domain Controller URL [%v%v%v%v]", protocol, url, port, postfix)
	os.Setenv("DCUrl", protocol+url+port+postfix)
	os.Setenv("DCIp", url)
	os.Setenv("DCPort", port)
}

func WorksSetting() {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "WorksSetting",
		}).Errorf("DB connect error[%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"worksInit": "WorksSetting",
	}).Infof("DB connect success")

	rows, err := db.Query("SELECT * FROM configuration WHERE name LIKE 'works.default%'")
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "WorksSetting",
		}).Errorf("worksInit Works Setting Query Failed[%v]", err)
	}
	defer rows.Close()

	result, err := rowsToString(rows)
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "WorksSetting",
		}).Errorf("Row to String conversion error [%v]", err)
	}

	jsonUnmarshal := []Configuration{}
	err = json.Unmarshal([]byte(result), &jsonUnmarshal)
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "WorksSetting",
		}).Errorf("String to JSON conversion error [%v]", err)
	}

	worksValue := map[string]interface{}{}
	for _, v := range jsonUnmarshal {
		worksValue[v.Name] = v.Value
	}
	url := worksValue["works.default.url"].(string)
	port := worksValue["works.default.port"].(string)

	log.WithFields(logrus.Fields{
		"worksInit": "WorksSetting",
	}).Infof("Works URL [%v][%v]", url, port)
	os.Setenv("WorksUrl", url+port)
	os.Setenv("WorksIp", url)
	os.Setenv("WorksPort", port)
}

func SambaSetting() {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "SambaSetting",
		}).Errorf("DB connect error[%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"worksInit": "SambaSetting",
	}).Infof("DB connect success")

	rows, err := db.Query("SELECT * FROM configuration WHERE name LIKE 'samba.default%'")
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "SambaSetting",
		}).Errorf("worksInit Samba Setting Query Failed[%v]", err)
	}
	defer rows.Close()

	result, err := rowsToString(rows)
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "SambaSetting",
		}).Errorf("Row to String conversion error [%v]", err)
	}

	jsonUnmarshal := []Configuration{}
	err = json.Unmarshal([]byte(result), &jsonUnmarshal)
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "SambaSetting",
		}).Errorf("String to JSON conversion error [%v]", err)
	}

	sambaValue := map[string]interface{}{}
	for _, v := range jsonUnmarshal {
		sambaValue[v.Name] = v.Value
	}
	url := sambaValue["samba.default.url"].(string)
	domain := sambaValue["samba.default.domain"].(string)

	log.WithFields(logrus.Fields{
		"worksInit": "WorksSetting",
	}).Infof("Works URL [%v] SAMBA Domain [%v]", url, domain)
	os.Setenv("SambaUrl", url)
	os.Setenv("SambaDomain", domain)
}

func GuacamoleSetting() {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "GuacamoleSetting",
		}).Errorf("DB connect error[%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"worksInit": "GuacamoleSetting",
	}).Infof("DB connect success")

	rows, err := db.Query("SELECT * FROM configuration WHERE name LIKE 'guacamole.default%'")
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "GuacamoleSetting",
		}).Errorf("worksInit Guacamole Setting Query Failed[%v]", err)
	}
	defer rows.Close()

	result, err := rowsToString(rows)
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "GuacamoleSetting",
		}).Errorf("Row to String conversion error [%v]", err)
	}

	jsonUnmarshal := []Configuration{}
	err = json.Unmarshal([]byte(result), &jsonUnmarshal)
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "GuacamoleSetting",
		}).Errorf("String to JSON conversion error [%v]", err)
	}

	guacamoleValue := map[string]interface{}{}
	for _, v := range jsonUnmarshal {
		guacamoleValue[v.Name] = v.Value
	}
	protocol := guacamoleValue["guacamole.default.protocol"].(string)
	url := guacamoleValue["guacamole.default.url"].(string)
	port := guacamoleValue["guacamole.default.port"].(string)
	postfix := guacamoleValue["guacamole.default.url.postfix"].(string)
	username := guacamoleValue["guacamole.default.username"].(string)

	log.WithFields(logrus.Fields{
		"worksInit": "WorksSetting",
	}).Infof("guacamole porotocol [%v], guacamole url [%v], guacamole port [%v], guacamole postfix [%v], guacamole username [%v]", protocol, url, port, postfix, username)
	os.Setenv("GuacamoleUrl", protocol+url+":"+port)
	os.Setenv("GuacamoleIp", url)
	os.Setenv("GuacamolePort", port)
	os.Setenv("GuacamoleUsername", username)
}

func ClusterNameSetting() {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "ClusterNameSetting",
		}).Errorf("DB connect error[%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"worksInit": "ClusterNameSetting",
	}).Infof("DB connect success")

	rows, err := db.Query("SELECT * FROM configuration WHERE name LIKE 'cluster.default%'")
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "ClusterNameSetting",
		}).Errorf("worksInit Guacamole Setting Query Failed[%v]", err)
	}
	defer rows.Close()

	result, err := rowsToString(rows)
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "ClusterNameSetting",
		}).Errorf("Row to String conversion error [%v]", err)
	}

	jsonUnmarshal := []Configuration{}
	err = json.Unmarshal([]byte(result), &jsonUnmarshal)
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "ClusterNameSetting",
		}).Errorf("String to JSON conversion error [%v]", err)
	}

	clusterValue := map[string]interface{}{}
	for _, v := range jsonUnmarshal {
		clusterValue[v.Name] = v.Value
	}
	clusterName := clusterValue["cluster.default.name"].(string)

	log.WithFields(logrus.Fields{
		"worksInit": "WorksSetting",
	}).Infof("clutster Name [%v]", clusterName)
	os.Setenv("ClusterName", clusterName)
}

func RDPPortSetting() {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "RDPPortSetting",
		}).Errorf("DB connect error[%v]", err)
	}
	defer db.Close()
	log.WithFields(logrus.Fields{
		"worksInit": "RDPPortSetting",
	}).Infof("DB connect success")

	rows, err := db.Query("SELECT * FROM configuration WHERE name LIKE 'rdp.default%'")
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "RDPPortSetting",
		}).Errorf("worksInit RDP port Setting Query Failed[%v]", err)
	}
	defer rows.Close()

	result, err := rowsToString(rows)
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "RDPPortSetting",
		}).Errorf("Row to String conversion error [%v]", err)
	}

	jsonUnmarshal := []Configuration{}
	err = json.Unmarshal([]byte(result), &jsonUnmarshal)
	if err != nil {
		log.WithFields(logrus.Fields{
			"worksInit": "RDPPortSetting",
		}).Errorf("String to JSON conversion error [%v]", err)
	}

	clusterValue := map[string]interface{}{}
	for _, v := range jsonUnmarshal {
		clusterValue[v.Name] = v.Value
	}
	portForRDP := clusterValue["rdp.default.port"].(string)

	log.WithFields(logrus.Fields{
		"worksInit": "RDPPortSetting",
	}).Infof("RDP Port [%v]", portForRDP)
	os.Setenv("PortForRDP", portForRDP)
}
