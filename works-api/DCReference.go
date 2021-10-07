package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

type UserInfo struct {
	Username, Password, Phone, Email string
	GivenName, Sn, Title             string
}

//login
func login(id string, password string) map[string]interface{} {
	var DCInfo = os.Getenv("DCUrl")
	params := url.Values{
		"username": {id},
		"password": {password},
	}
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	var res map[string]interface{}
	resp, err := client.PostForm(DCInfo+"/v1/login", params)
	if err != nil {
		log.Error(err)
		log.Errorf("%v\n", resp)
	} else {
		json.NewDecoder(resp.Body).Decode(&res)
	}

	return res
}

func selectUserDetail(id string) *http.Response {
	var DCInfo = os.Getenv("DCUrl")
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	//resp, err := client.PostForm(DCInfo+"/v1/user/", params)
	resp, err := client.Get(DCInfo + "/v1/user/" + id)
	if err != nil {
		log.Error(err)
		log.Error(resp)
	}

	return resp
}

func selectUserList() []map[string]interface{} {
	var DCInfo = os.Getenv("DCUrl")
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	//resp, err := client.PostForm(DCInfo+"/v1/user/", params)
	resp, err := client.Get(DCInfo + "/v1/user/")
	var res []map[string]interface{}

	if err != nil {
		log.Error(err)
		log.Error(resp)
	} else {
		json.NewDecoder(resp.Body).Decode(&res)
	}

	return res
}

func insertDCUser(userInfo UserInfo) (*http.Response, error) {
	var DCInfo = os.Getenv("DCUrl")
	params := url.Values{
		"username":  {userInfo.Username},
		"password":  {userInfo.Password},
		"email":     {userInfo.Email},
		"sn":        {userInfo.Sn},
		"givenName": {userInfo.GivenName},
		"phone":     {userInfo.Phone},
		"title":     {userInfo.Title},
	}
	log.Infof("paramsInfo = [%v]", params)
	log.Infof("userInfo = [%v]", userInfo)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.PostForm(DCInfo+"/v1/user", params)

	return resp, err
}

func deleteDCUser(username string) (*http.Response, error) {
	var DCInfo = os.Getenv("DCUrl")
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%v/v1/user/%v", DCInfo, username), nil)
	if err != nil {
		log.Errorf("유저 삭제중 에러가 발생했습니다. [%v]", err)
	}
	resp, err1 := client.Do(req)
	return resp, err1
}

func insertGroup(groupName string) (*http.Response, error) {
	var DCInfo = os.Getenv("DCUrl")
	params := url.Values{
		"groupname": {groupName},
	}
	log.Infof("paramsInfo = [%v]", params)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.PostForm(DCInfo+"/v1/group", params)

	log.Infof("%v %v", resp, err)
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
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(DCInfo + "/v1/group/" + groupName)

	return resp, err
}

func deleteGroupDetail(groupName string) (*http.Response, error) {
	var DCInfo = os.Getenv("DCUrl")

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(DCInfo + "/v1/group/" + groupName)

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
	result, err := db.Exec("INSERT INTO users(uuid, user_name, password, create_date) VALUES (?, ?, ?, now())", getUuid(), userInfo.Username, userInfo.Password)
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

func insertUserAllocatedInstance(username string, connectName string, parameter string) *http.Response {
	var DCInfo = os.Getenv("DCUrl")
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	params := url.Values{
		"username":  {username},
		"parameter": {parameter},
	}
	resp, err := client.PostForm(DCInfo+"/v1/connection/"+connectName, params)
	//req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%v/v1/user/%v", DCInfo, username), params)
	if err != nil {
		log.Errorf("유저 삭제중 에러가 발생했습니다. [%v]", err)
	}

	log.Infof(resp.Status)
	return resp
}
