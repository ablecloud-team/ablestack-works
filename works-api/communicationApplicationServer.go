package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Application struct {
	id                   int      `json:"id"`
	Uuid                 string   `json:"uuid"`
	WorkspaceUuid        string   `json:"workspace_uuid"`
	Name                 string   `json:"name"`
	Path                 string   `json:"path"`
	Description          string   `json:"description"`
	WorkingDirectory     string   `json:"working_directory"`
	ApplicationServer    []string `json:"application_server"`
	RegApplicationServer string   `json:"reg_application_server"`
	CreateDate           string   `json:"create_date"`
}

// Application List 를 조회하는 func
func getApplicationList(url string) ([]Application, error) {
	var appServer = "http://" + url + ":8083/api"
	log.WithFields(logrus.Fields{
		"communicationApplicationServer": "getApplicationList",
	}).Infof("payload. [%v]", url)
	//client := http.Client{
	//	Timeout: 60 * time.Second,
	//}
	//resp, err := client.PostForm(DCInfo+"/v1/user/", params)
	resp, err := http.Get(appServer + "/v1/app")
	var appInfoList []Application
	if err != nil {
		log.WithFields(logrus.Fields{
			"communicationApplicationServer": "getApplicationList",
		}).Errorf("ApplicationServer communication failed. [%v], error [%v]", resp, err)
	} else {
		var res map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&res)
		appList, _ := json.Marshal(res["app"])
		json.Unmarshal(appList, &appInfoList)
		log.WithFields(logrus.Fields{
			"communicationApplicationServer": "getApplicationList",
		}).Infof("Application Server communication result. resp.StatusCode [%v], appInfoList [%v], res [%v]", resp.StatusCode, appInfoList, res["app"])
		//json.NewDecoder(resp.Body).Decode(&res)
	}

	return appInfoList, err
}
