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
		log.Error("Mold 와 통신에 실패했습니다.(listTemplates)")
		log.Error(err.Error())
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
		log.Error("Mold 와 통신에 실패했습니다.(listServiceOfferings)")
		log.Error(err.Error())
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
	log.Info("Mold 통신 URL [" + endUrl + "]")
	resp, err := http.Get(endUrl)
	if err != nil {
		log.Error("Mold 와 통신에 실패했습니다.(listNetworks)")
		log.Error(err.Error())
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
	log.Info("Mold 통신 URL [" + endUrl + "]")
	resp, err := http.Get(endUrl)
	if err != nil {
		log.Error("Mold 와 통신에 실패했습니다.(listDiskOfferings)")
		log.Error(err.Error())
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getDeployVirtualMachine(params []MoldParams) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "deployVirtualMachine"},
	}
	params = append(params, params1...)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig
	log.Info("Mold 통신 URL [" + endUrl + "]")
	resp, err := http.Get(endUrl)
	if err != nil {
		log.Error("Mold 와 통신에 실패했습니다.(deployVirtualMachine)")
		log.Error(err.Error())
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}
