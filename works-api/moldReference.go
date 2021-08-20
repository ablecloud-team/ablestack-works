package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

//var baseurl = os.Getenv("MoldUrl")

func getTemplate(params []MoldParams) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"templatefilter": "executable"},
		{"command": "listTemplates"},
		{"id": "5905386f-f644-4cb0-b5c4-97c7c53e6f9b"},
	}
	params = append(params, params1...)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig
	resp, err := http.Get(endUrl)
	log.WithFields(logrus.Fields{
		"MoldEndURL": endUrl,
	}).Infof("Starting application")
	if err != nil {

	}
	fmt.Println(resp)
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getComputeOffering(params []MoldParams) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "listServiceOfferings"},
	}
	params = append(params, params1...)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig
	resp, err := http.Get(endUrl)
	if err != nil {

	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getNetwork(params []MoldParams) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "listNetworks"},
	}
	params = append(params, params1...)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig
	resp, err := http.Get(endUrl)
	if err != nil {

	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getDiskOffering(params []MoldParams) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "listDiskOfferings"},
	}
	params = append(params, params1...)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig
	resp, err := http.Get(endUrl)
	if err != nil {

	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}
