package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func getTemplate(params []MoldParams) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"templatefilter": "all"},
		{"command": "listTemplates"},
	}

	stringParams := makeStringParams(params1)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig
	resp, err := http.Get(endUrl)
	if err != nil {

	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getComputeOffering(params []MoldParams) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "listServiceOfferings"},
	}

	stringParams := makeStringParams(params1)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig
	resp, err := http.Get(endUrl)
	if err != nil {

	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}
