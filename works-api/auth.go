package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

//login
func login(id string, password string) map[string]interface{} {
	params := url.Values{
		"username": {id},
		"password": {password},
	}
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.PostForm(DCInfo+"/v1/login", params)
	if err != nil {
		log.Fatal(err)
		fmt.Println(resp)
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func userInfo(id string) map[string]interface{} {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	//resp, err := client.PostForm(DCInfo+"/v1/user/", params)
	resp, err := client.Get(DCInfo+"/v1/user/"+id)
	if err != nil {
		log.Fatal(err)
		fmt.Println(resp)
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}
