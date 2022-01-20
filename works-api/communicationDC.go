package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"os"
	"time"
)

type UserInfo struct {
	C                 string `json:"c"`                 // 국가코드 영문
	Cn                string `json:"cn"`                // 유저아이디
	Co                string `json:"co"`                // 국가코드 영문
	CountryCode       string `json:"countryCode"`       // 국가코드 숫자
	DistinguishedName string `json:"distinguishedName"` // 고유이름
	GivenName         string `json:"givenName"`         // 사용자 이름
	Mail              string `json:"mail"`              // e-mail
	MemberOf          string `json:"memberOf"`          // 소속 그룹
	Name              string `json:"name"`              // 사용자 계정
	SAMAccountName    string `json:"sAMAccountName"`    // 사용자 계정
	Sn                string `json:"sn"`                // 사용자 성
	TelephoneNumber   string `json:"telephoneNumber"`   // 전화번호
	Title             string `json:"title"`             // 직책
	UserPrincipalName string `json:"userPrincipalName"` // 사용자 도메인 계정정보
	Password          string `json:"password"`          // 사용자 비밀번호
	Department        string `json:"department"`        // 사용자 부서
}

//login
func postLogin(id string, password string) (*http.Response, error) {
	var DCInfo = os.Getenv("DCUrl")
	params := url.Values{
		"username": {id},
		"password": {password},
	}
	client := http.Client{}
	resp, err := client.PostForm(DCInfo+"/v1/login", params)
	//if err != nil {
	//	log.Error(err)
	//	log.Errorf("%v\n", resp)
	//} else {
	//	json.NewDecoder(resp.Body).Decode(&res)
	//}

	return resp, err
}

//************************** User 관련 시작 **************************
//유저를 DC 에 추가하는 func
func postDCUser(userInfo UserInfo) (*http.Response, error) {
	var DCInfo = os.Getenv("DCUrl")
	params := url.Values{
		"username":  {userInfo.Cn},
		"password":  {userInfo.Password},
		"sn":        {userInfo.Sn},
		"givenName": {userInfo.GivenName},
		"email":     {userInfo.Mail},
		"phone":     {userInfo.TelephoneNumber},
		"title":     {userInfo.Title},
	}
	log.Infof("paramsInfo = [%v]", params)
	log.Infof("userInfo = [%v]", userInfo)
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	resp, err := client.PostForm(DCInfo+"/v1/user", params)

	return resp, err
}

// 유저 리스트를 DC 에서 조회하는 func
func getUserList() ([]UserInfo, error) {
	var DCInfo = os.Getenv("DCUrl")
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	//resp, err := client.PostForm(DCInfo+"/v1/user/", params)
	resp, err := client.Get(DCInfo + "/v1/user/")
	var res []map[string]interface{}
	var userInfoList []UserInfo
	if err != nil {
		log.WithFields(logrus.Fields{
			"authDAO": "getUserList",
		}).Errorf("DC communication failed. [%v], error [%v]", resp, err)
	} else {
		json.NewDecoder(resp.Body).Decode(&res)
		jsonMarshal, _ := json.Marshal(res)
		json.Unmarshal([]byte(jsonMarshal), &userInfoList)
		log.WithFields(logrus.Fields{
			"authDAO": "getUserList",
		}).Infof("DC communication result. resp.StatusCode [%v], userInfoList [%v]", resp.StatusCode, userInfoList)
		//json.NewDecoder(resp.Body).Decode(&res)
	}

	return userInfoList, err
}

// 유저 정보를 DC 에서 조회하는 func
func getUserInfo(id string) *http.Response {
	var DCInfo = os.Getenv("DCUrl")
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	//resp, err := client.PostForm(DCInfo+"/v1/user/", params)
	resp, err := client.Get(DCInfo + "/v1/user/" + id)
	if err != nil {
		log.Error(err)
		log.Error(resp)
	}

	return resp
}

func deleteDCUser(userName string) (*http.Response, error) {
	var DCInfo = os.Getenv("DCUrl")
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%v/v1/user/%v", DCInfo, userName), nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("유저 삭제중 에러가 발생했습니다. [%v]", err)
	}
	return resp, err
}

//************************** User 관련 끝 **************************

func insertGroup(groupName string) (*http.Response, error) {
	var DCInfo = os.Getenv("DCUrl")
	params := url.Values{
		"groupname": {groupName},
	}
	log.Infof("paramsInfo = [%v]", params)
	client := http.Client{}
	resp, err := client.PostForm(DCInfo+"/v1/group", params)

	log.Infof("%v %v", resp, err)
	return resp, err
}

func insertPolicyRemotefx(groupName string) (*http.Response, error) {
	var DCInfo = os.Getenv("DCUrl")
	client := http.Client{
		Timeout: 20 * time.Second,
	}
	params := url.Values{
		"groupname": {groupName},
	}
	log.Infof("paramsInfo = [%v]", params)

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%v/v1/policy/remotefx/%v", DCInfo, groupName), nil)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	log.WithFields(logrus.Fields{
		"dcDAO": "insertPolicyRemotefx",
	}).Infof("workspace group add remotrfx status. resp [%v], err [%v]", resp, err)
	return resp, err
}

func selectWorkspacePolicyList(workspace Workspace) (*http.Response, error) {
	var DCInfo = os.Getenv("DCUrl")

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(DCInfo + "/v1/group/" + workspace.Name + "/policy")

	return resp, err
}

func selectGroupList() (*http.Response, error) {
	var DCInfo = os.Getenv("DCUrl")

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(DCInfo + "/v1/group/")

	return resp, err
}

func selectGroupDetail(groupName string) (*http.Response, error) {
	var DCInfo = os.Getenv("DCUrl")

	client := http.Client{
		Timeout: 60 * time.Second,
	}
	resp, err := client.Get(DCInfo + "/v1/group/" + groupName)

	return resp, err
}

func deleteGroup(groupName string) (*http.Response, error) {
	var DCInfo = os.Getenv("DCUrl")

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	//resp, err := client.Get(DCInfo + "/v1/group/" + groupName)
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%v/v1/group/%v", DCInfo, groupName), nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	log.WithFields(logrus.Fields{
		"authDAO": "deleteGroup",
	}).Infof("workspace group delete status. resp [%v], err [%v]", resp, err)

	return resp, err
}

func insertAddUserToGroup(groupName string, userName string) (*http.Response, error) {
	var DCInfo = os.Getenv("DCUrl")
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%v/v1/group/%v/%v", DCInfo, groupName, userName), nil)
	if err != nil {
		log.Errorf("유저 삭제중 에러가 발생했습니다. [%v]", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err1 := client.Do(req)
	return resp, err1
}

func deleteAddUserToGroup(groupName string, userName string) (*http.Response, error) {
	var DCInfo = os.Getenv("DCUrl")
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%v/v1/group/%v/%v", DCInfo, groupName, userName), nil)
	if err != nil {
		log.Errorf("유저 삭제중 에러가 발생했습니다. [%v]", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err1 := client.Do(req)
	return resp, err1
}

func insertUserDB(userInfo UserInfo) map[string]interface{} {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	resultReturn := map[string]interface{}{}
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
		resultReturn["message"] = "DB connect error"
	}
	defer db.Close()
	result, err := db.Exec("INSERT INTO users(uuid, user_name, password, create_date) VALUES (?, ?, ?, now())", getUuid(), userInfo.Cn, userInfo.Password)
	if err != nil {
		log.Errorf("유저를 DB 등록중에러가 발생했습니다. [%v]", err)
		resultReturn["message"] = "An error occurred while registering Async Job."
		resultReturn["status"] = SQLQueryError
	}
	n, _ := result.RowsAffected()
	if n == 1 {
		log.Info("유저가 정상적으로 등록되었습니다.")
		resultReturn["status"] = http.StatusOK
		resultReturn["message"] = "The async job has been successfully registered."
	}
	return resultReturn
}

func insertUserAllocatedInstance1(username string, instanceUuid string) *http.Response {
	var DCInfo = os.Getenv("DCUrl")
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resultInstance, _ := selectInstanceList(instanceUuid, InstanceString)
	instanceInfo := resultInstance[0]
	resultUser := selectUserDBDetail(username)
	paramsMold := []MoldParams{
		{"id": instanceInfo.MoldUuid},
	}
	resultMoldInstance := getListVirtualMachinesMetrics(paramsMold)
	listVirtualMachinesMetrics := ListVirtualMachinesMetrics{}
	virtualMachineInfo, _ := json.Marshal(resultMoldInstance["listvirtualmachinesmetricsresponse"])
	json.Unmarshal([]byte(virtualMachineInfo), &listVirtualMachinesMetrics)
	log.Infof("%v", listVirtualMachinesMetrics)
	params := url.Values{
		"connection": {instanceInfo.Name},
		"parameter":  {"hostname=" + listVirtualMachinesMetrics.Virtualmachine[0].Nic[0].Ipaddress + ",port=3389,ignore-cert=true,username=" + username + ",password=" + resultUser.Password + ",domain=testdomain"},
	}
	resp, err := client.PostForm(DCInfo+"/v1/user/"+username, params)
	//req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%v/v1/user/%v", DCInfo, username), params)
	if err != nil {
		log.Errorf("유저 삭제중 에러가 발생했습니다. [%v]", err)
	}

	log.Infof(resp.Status)
	return resp
}

func insertConnection(username string, connectName string, parameter string) *http.Response {
	var DCInfo = os.Getenv("DCUrl")
	client := http.Client{}

	params := url.Values{
		"username":  {username},
		"parameter": {parameter},
	}
	//resp, err := client.PostForm(DCInfo+"/v1/connection/"+connectName, params)
	resp, err := client.PostForm(fmt.Sprintf("%v/v1/connection/%v", DCInfo, connectName), params)
	if err != nil {
		log.Errorf("An error occurred while creating the connection information. [%v]", err)
	}

	log.Infof(resp.Status)
	return resp
}

func delConnection(connectName string) *http.Response {
	var DCInfo = os.Getenv("DCUrl")
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("%v/v1/connection/%v", DCInfo, connectName), nil)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.WithFields(logrus.Fields{
			"communicationMold.go": "delConnection",
		}).Errorf("An error occurred while deleting the connection. resp [%v], err [%v]", resp, err)
	}

	return resp
}
