package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type settingInfo struct {
	Database struct{
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
	Mold struct {
		Protocol string `json:"protocol"`
		Url      string `json:"url"`
		Port     string `json:"port"`
		Postfix   string `json:"postfix"`
		User    struct {
			Account   string `json:"account"`
			ApiKey    string `json:"apiKey"`
			SecretKey string `json:"secretKey"`
		} `json:"user"`
	} `json:"mold"`
	DomainController struct {
		Protocol string `json:"protocol"`
		Url      string `json:"url"`
		Port     string `json:"port"`
		Postfix   string `json:"postfix"`
	} `json:"domainController"`
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