package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

type Configuration struct {
	Id, Name, Category, Value, DefaultValue, Description string
	UpdateDate                                           time.Time
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

func MoldSetting() {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	log.Debug("||||||||||||||||||||||||||||||")
	log.Debug(os.Getenv("MysqlType"))
	log.Debug(os.Getenv("DbInfo"))
	if err != nil {
		fmt.Println("DB connect error")
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println("DB connect success")

	rows, err := db.Query("SELECT * FROM configuration WHERE name LIKE 'mold.default%'")
	if err != nil {
		fmt.Println("worksInit Mold Setting Query Failed")
		fmt.Println(err.Error())
	}
	defer rows.Close()

	result, err := rowsToString(rows)
	if err != nil {
		fmt.Println(err)
	}

	jsonUnmarshal := []Configuration{}
	err = json.Unmarshal([]byte(result), &jsonUnmarshal)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
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
	fmt.Println(moldValue)
	fmt.Println("Set Mold URL = " + protocol + url + port + postfix)
	os.Setenv("MoldUrl", protocol+url+port+postfix)
	os.Setenv("MoldApiKey", apiKey)
	os.Setenv("MoldSecretKey", secretKey)
}

func DCSetting() {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		fmt.Println("DB connect error")
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println("DB connect success")

	rows, err := db.Query("SELECT * FROM configuration WHERE name LIKE 'dc.default%'")
	if err != nil {
		fmt.Println("worksInit Mold Setting Query Failed")
		fmt.Println(err.Error())
	}
	defer rows.Close()

	result, err := rowsToString(rows)
	if err != nil {
		fmt.Println(err)
	}

	jsonUnmarshal := []Configuration{}
	err = json.Unmarshal([]byte(result), &jsonUnmarshal)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}

	dcValue := map[string]interface{}{}
	for _, v := range jsonUnmarshal {
		dcValue[v.Name] = v.Value
	}
	fmt.Println(dcValue)
	protocol := dcValue["dc.default.protocol"].(string)
	url := dcValue["dc.default.url"].(string)
	port := dcValue["dc.default.port"].(string)
	if port == "80" {
		port = ""
	} else {
		port = ":" + port
	}
	postfix := dcValue["dc.default.url.postfix"].(string)

	fmt.Println("Set Domain Controller URL = " + protocol + url + port + postfix)
	os.Setenv("DCUrl", protocol+url+port+postfix)
}
