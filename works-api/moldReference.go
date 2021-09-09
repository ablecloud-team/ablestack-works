package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func getTemplate(params []MoldParams) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "listTemplates"},
	}
	params = append(params, params1...)
	log.WithFields(logrus.Fields{
		"params": "params",
	}).Debugf("%v", params)
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
	log.Infof("Template 조회 결과값 [%v]", resp)
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
	log.Infof("ComputeOffering 조회 결과값 [%v]", resp)
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

func getDeployVirtualMachine(workspaceUuid string) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	workspaceInfo := selectWorkspaceDetail(workspaceUuid)
	displayName := workspaceInfo.Name + "-" + postfixFill(workspaceInfo.Postfix)
	params := []MoldParams{
		{"command": "deployVirtualMachine"},
		{"templateid": workspaceInfo.TemplateUuid},
		{"displayname": displayName},
		{"networkids": workspaceInfo.NetworkUuid},
		{"serviceofferingid": workspaceInfo.ComputeOfferingUuid},
		{"zoneid": selectZoneId()},
	}
	log.Info("09909090909090909090909090")
	log.Info(workspaceInfo)
	log.Info(params)
	log.Info("09909090909090909090909090")
	//params = append(params, params1...)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig
	log.Info("Mold stringParams [" + stringParams + "]")
	log.Info("Mold sig [" + sig + "]")
	log.Info("Mold 통신 URL [" + endUrl + "]")
	log.Info("endUrl1 [" + baseurl + "?" + stringParams + "&signature=" + sig + "]")
	log.Info("endUrl2 [" + endUrl + "]")
	res := map[string]interface{}{}
	resp, err := http.Get(baseurl + "?" + stringParams + "&signature=" + sig)
	if err != nil {
		res["message"] = "Mold 와 통신에 실패했습니다.(deployVirtualMachine)"
		res["status"] = BaseErrorCode
	} else {
		json.NewDecoder(resp.Body).Decode(&res)
		res["status"] = http.StatusOK
	}
	return res
}

func getDestroyVirtualMachine(params []MoldParams) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "destroyVirtualMachine"},
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

func getCreateTags(params []MoldParams) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "createTags"},
		{"resourcetype": "UserVm"},
	}
	params = append(params, params1...)

	stringParams := makeStringParams(params)
	log.Infof("stringParams = [%v]", stringParams)
	sig := makeSignature(stringParams)
	//endUrl := baseurl + "?" + stringParams + "&signature=" + sig
	log.Info("Mold 통신 URL {" + baseurl + "?" + stringParams + "&signature=" + sig + "}")
	resp, err := http.Get(baseurl + "?" + stringParams + "&signature=" + sig)
	if err != nil {
		log.Error("Mold 와 통신에 실패했습니다.(deployVirtualMachine)")
		log.Error(err.Error())
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getQueryAsyncJobResult(params []MoldParams) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "queryAsyncJobResult"},
	}
	params = append(params, params1...)

	stringParams := makeStringParams(params)
	log.Infof("stringParams = [%v]", stringParams)
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
