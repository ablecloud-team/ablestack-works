package main

import (
	"encoding/json"
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
		res["status"] = ""
	} else {
		json.NewDecoder(resp.Body).Decode(&res)
	}

	return res
}

func selectUserDetail(id string) map[string]interface{} {
	var DCInfo = os.Getenv("DCUrl")
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	//resp, err := client.PostForm(DCInfo+"/v1/user/", params)
	resp, err := client.Get(DCInfo + "/v1/user/" + id)
	var res map[string]interface{}
	if err != nil {
		log.Error(err)
		log.Error(resp)
	} else {
		json.NewDecoder(resp.Body).Decode(&res)
	}

	return res
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

func putDCUser(userInfo UserInfo) map[string]interface{} {

	var DCInfo = os.Getenv("DCUrl")
	params := url.Values{
		"username":  {userInfo.Username},
		"password":  {userInfo.Password},
		"phone":     {userInfo.Phone},
		"email":     {userInfo.Email},
		"sn":        {userInfo.Sn},
		"givenName": {userInfo.GivenName},
		"title":     {userInfo.Title},
	}
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	var res map[string]interface{}
	resp, err := client.PostForm(DCInfo+"/v1/user", params)
	if err != nil {
		log.Error(err)
		log.Error(resp)
	} else {
		json.NewDecoder(resp.Body).Decode(&res)
	}

	return res
}
