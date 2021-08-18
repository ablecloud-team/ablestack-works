package main

import (
	"fmt"
	"net/http"
	"os"
)

func getTemplate(params map[string]string) *http.Response {
	var secretKey string = os.Getenv("MoldSecretKey")
	var apiKey string = os.Getenv("MoldApiKey")
	var baseurl string = os.Getenv("MoldUrl")
	params["apikey"] = apiKey
	params["command"] = "listTemplates"
	params["templatefilter"] = "all"
	params["response"] = "json"
	sig := makeSignature(params)
	stringParams := makeStringParams(params)
	fmt.Println("123123123123")
	fmt.Println(apiKey)
	fmt.Println(secretKey)
	fmt.Println(baseurl)
	fmt.Println(stringParams)
	fmt.Println(sig)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig
	fmt.Println("endUrl")
	fmt.Println(endUrl)
	resp, err := http.Get(endUrl)
	if err != nil {

	}
	fmt.Println("resp")
	fmt.Println(resp)
	return resp
}
