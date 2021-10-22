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
		"moldReference.go": "getTemplate",
	}).Infof("params [%v]", params)
	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig
	log.WithFields(logrus.Fields{
		"moldReference.go": "getTemplate",
	}).Infof("endUrl [%v]", endUrl)
	resp, err := http.Get(endUrl)

	if err != nil {
		log.WithFields(logrus.Fields{
			"moldReference.go": "getTemplate",
		}).Errorf("Mold 와 통신에 실패했습니다.(listTemplates) [%v]", err)
	}

	log.WithFields(logrus.Fields{
		"moldReference.go": "getTemplate",
	}).Debugf("result [%v]", resp)

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

	log.WithFields(logrus.Fields{
		"moldReference.go": "getComputeOffering",
	}).Infof("params [%v]", params)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	log.WithFields(logrus.Fields{
		"moldReference.go": "getComputeOffering",
	}).Infof("endUrl [%v]", endUrl)

	resp, err := http.Get(endUrl)
	if err != nil {
		log.WithFields(logrus.Fields{
			"moldReference.go": "getComputeOffering",
		}).Errorf("Mold 와 통신에 실패했습니다.(listServiceOfferings) [%v]", err)
	}

	log.WithFields(logrus.Fields{
		"moldReference.go": "getComputeOffering",
	}).Debugf("result [%v]", resp)

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

	log.WithFields(logrus.Fields{
		"moldReference.go": "getNetwork",
	}).Infof("params [%v]", params)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	log.WithFields(logrus.Fields{
		"moldReference.go": "getNetwork",
	}).Infof("endUrl [%v]", endUrl)

	resp, err := http.Get(endUrl)

	log.WithFields(logrus.Fields{
		"moldReference.go": "getComputeOffering",
	}).Debugf("result [%v]", resp)

	if err != nil {
		log.WithFields(logrus.Fields{
			"moldReference.go": "getComputeOffering",
		}).Errorf("Mold 와 통신에 실패했습니다.(listNetworks) [%v]", err)
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getDeployVirtualMachine(workspaceUuid string, instanceUuid string, instanceType string) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	workspaceList, _ := selectWorkspaceList(workspaceUuid)
	workspaceInfo := workspaceList[0]

	log.WithFields(logrus.Fields{
		"moldReference.go": "getDeployVirtualMachine",
	}).Infof("payload workspaceUuid [%v], instanceUuid [%v], instanceType [%v]", workspaceUuid, instanceUuid, instanceType)

	var displayName string
	if workspaceInfo.Postfix == 0 {
		displayName = workspaceInfo.Name + "-TestVM"
	} else {
		displayName = workspaceInfo.Name + "-" + postfixFill(workspaceInfo.Postfix)
	}
	SambaIP := os.Getenv("SambaUrl")
	SambaDomain := os.Getenv("SambaDomain") + ".local"
	MyDomain := os.Getenv("SambaDomain")
	WorksIP := os.Getenv("WorksIp")
	WorksPort := os.Getenv("WorksPort")
	VmName := displayName
	InstanceUuid := instanceUuid
	Type := instanceType
	//"Invoke-WebRequest -Uri http://10.1.1.1/latest/user-data >> \"c:\\agent\\userdata.txt\"\n" +
	payload := "<powershell>\n" +
		"date > \"c:\\agent\\installed.txt\"\n" +
		"$dnsip = \"" + SambaIP + "\"\n" +
		"echo $dnsip >> \"c:\\agent\\installed.txt\"\n" +
		"set-DnsClientServerAddress -InterfaceIndex (Get-NetAdapter |Select-Object InterfaceAlias, interfaceindex).interfaceindex -ServerAddresses " + SambaIP + "\n" +
		"$password = \"Ablecloud1!\" | ConvertTo-SecureString -asPlainText -Force\n" +
		"$username = \"$" + MyDomain + "\\administrator\" \n" +
		"$credential = New-Object System.Management.Automation.PSCredential($username,$password)\n" +
		"echo Rename-Computer >> \"c:\\agent\\installed.txt\"\n" +
		"Rename-Computer -NewName " + VmName + "\n" +
		"echo Add-Computer >> \"c:\\agent\\installed.txt\"\n" +
		"Add-Computer -DomainName " + SambaDomain + " -Credential $credential -NewName " + VmName + "\n" +
		"echo Add-Computer end>> \"c:\\agent\\installed.txt\"\n" +
		"$conf = '{\"WorksServer\": \"" + WorksIP + "\", \"WorksPort\": " + WorksPort + ", \"Type\": \"" + Type + "\", \"UUID\": \"" + InstanceUuid + "\"}'\n" +
		"echo $conf| Out-File -Encoding ascii \"c:\\agent\\conf.json\"\n" +
		"echo $conf\n" +
		"echo $conf >> \"c:\\agent\\installed.txt\"\n" +
		"C:\\agent\\nssm.exe set \"Ablecloud Works Agent\" start SERVICE_DELAYED_AUTO_START\n" +
		"C:\\agent\\nssm.exe restart \"Ablecloud Works Agent\"\n" +
		"date >> \"c:\\agent\\installed.txt\"\n" +
		"</powershell>"
	params := []MoldParams{
		{"command": "deployVirtualMachine"},
		{"templateid": workspaceInfo.TemplateUuid},
		{"displayname": displayName},
		{"serviceofferingid": workspaceInfo.ComputeOfferingUuid},
		{"zoneid": selectZoneId()},
		{"userdata": baseEncoding(payload)},
	}

	log.WithFields(logrus.Fields{
		"moldReference.go": "getDeployVirtualMachine",
	}).Infof("params params [%v]", params)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	log.WithFields(logrus.Fields{
		"moldReference.go": "getDeployVirtualMachine",
	}).Infof("endUrl [%v]", endUrl)

	res := map[string]interface{}{}
	resp, err := http.Get(endUrl)
	if err != nil {

		log.WithFields(logrus.Fields{
			"moldReference.go": "getDeployVirtualMachine",
		}).Errorf("Mold 와 통신에 실패했습니다.(getDeployVirtualMachine) [%v]", err)

		res["message"] = "Mold 와 통신에 실패했습니다.(getDeployVirtualMachine)"
		res["status"] = BaseErrorCode
	} else {
		if instanceType == WorkspaceString {
			updateWorkspaceTemplateCheck(workspaceInfo.Uuid, AgentCheck)
		}
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

func getStartVirtualMachine(params []MoldParams) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "startVirtualMachine"},
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

func getStopVirtualMachine(params []MoldParams) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "stopVirtualMachine"},
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

func getListVirtualMachinesMetrics(params []MoldParams) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "listVirtualMachinesMetrics"},
	}
	params = append(params, params1...)

	stringParams := makeStringParams(params)
	log.Infof("stringParams = [%v]", stringParams)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig
	log.Infof("getListVirtualMachinesMetrics URL [%v]", endUrl)
	resp, err := http.Get(endUrl)
	if err != nil {
		log.Error("Mold 와 통신에 실패했습니다.(listVirtualMachinesMetrics)")
		log.Error(err.Error())
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	log.Debugf("getListVirtualMachinesMetrics result [%v]", res)

	return res
}

func getlistVolumesMetrics(params []MoldParams) map[string]interface{} {
	var baseurl string = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "listVolumesMetrics"},
	}
	params = append(params, params1...)

	stringParams := makeStringParams(params)
	log.Infof("stringParams = [%v]", stringParams)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig
	log.Info("Mold 통신 URL [" + endUrl + "]")
	resp, err := http.Get(endUrl)
	if err != nil {
		log.Error("Mold 와 통신에 실패했습니다.(listVirtualMachinesMetrics)")
		log.Error(err.Error())
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)

	return res
}
