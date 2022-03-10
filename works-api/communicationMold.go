package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func getTemplate(params []MoldParams) map[string]interface{} {
	var baseurl = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "listTemplates"},
	}
	params = append(params, params1...)
	log.WithFields(logrus.Fields{
		"communicationMold.go": "getTemplate",
	}).Infof("listTemplates params [%v]", params)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getTemplate",
	}).Infof("listTemplates endUrl [%v]", endUrl)

	resp, err := http.Get(endUrl)

	if err != nil {
		log.WithFields(logrus.Fields{
			"communicationMold.go": "getTemplate",
		}).Errorf("Failed to communicate with Mold. (listTemplates) [%v]", err)
	}

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getTemplate",
	}).Debugf("listTemplates result [%v]", resp)

	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getListDesktopMasterVersions(params []MoldParams) map[string]interface{} {
	var baseurl = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "listDesktopMasterVersions"},
	}
	params = append(params, params1...)

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getListDesktopMasterVersions",
	}).Infof("listDesktopMasterVersions params [%v]", params)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getListDesktopMasterVersions",
	}).Infof("listDesktopMasterVersions endUrl [%v]", endUrl)

	resp, err := http.Get(endUrl)

	if err != nil {
		log.WithFields(logrus.Fields{
			"communicationMold.go": "getListDesktopMasterVersions",
		}).Errorf("Failed to communicate with Mold. (listDesktopMasterVersions) [%v]", err)
	}

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getListDesktopMasterVersions",
	}).Debugf("listDesktopMasterVersions result [%v]", resp)

	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getComputeOffering(params []MoldParams) map[string]interface{} {
	var baseurl = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "listServiceOfferings"},
	}
	params = append(params, params1...)

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getComputeOffering",
	}).Infof("listServiceOfferings params [%v]", params)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getComputeOffering",
	}).Infof("listServiceOfferings endUrl [%v]", endUrl)

	resp, err := http.Get(endUrl)
	if err != nil {
		log.WithFields(logrus.Fields{
			"communicationMold.go": "getComputeOffering",
		}).Errorf("Failed to communicate with Mold. (listServiceOfferings) [%v]", err)
	}

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getComputeOffering",
	}).Debugf("listServiceOfferings result [%v]", resp)

	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getNetwork(params []MoldParams) map[string]interface{} {
	var baseurl = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "listNetworks"},
	}
	params = append(params, params1...)

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getNetwork",
	}).Infof("listNetworks params [%v]", params)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getNetwork",
	}).Infof("listNetworks endUrl [%v]", endUrl)

	resp, err := http.Get(endUrl)

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getComputeOffering",
	}).Debugf("listNetworks result [%v]", resp)

	if err != nil {
		log.WithFields(logrus.Fields{
			"communicationMold.go": "getComputeOffering",
		}).Errorf("Failed to communicate with Mold. (listNetworks) [%v]", err)
	}
	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getDeployVirtualMachine(workspaceInfo Workspace, instanceType string) (map[string]interface{}, string) {
	var baseurl = os.Getenv("MoldUrl")

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getDeployVirtualMachine",
	}).Infof("payload workspaceInfo [%v], instanceType [%v]", workspaceInfo, instanceType)

	var displayName string
	if instanceType == WorkspaceString {
		displayName = workspaceInfo.Name + "-TestVM"
	} else {
		displayName = workspaceInfo.Name + "-" + postfixFill(workspaceInfo.Postfix)
	}
	SambaIP := os.Getenv("SambaUrl")
	MyDomain := os.Getenv("SambaDomain")
	WorksIP := os.Getenv("WorksIp")
	WorksPort := os.Getenv("WorksPort")
	VmName := displayName
	instanceUuid := getUuid()
	Type := instanceType
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
		"echo Add-Computer end>> \"c:\\agent\\installed.txt\"\n" +
		"$conf = '{\"WorksServer\": \"" + WorksIP + "\", \"WorksPort\": " + WorksPort + ", \"Type\": \"" + Type + "\", \"UUID\": \"" + instanceUuid + "\",\"HostName\": \"" + VmName + "\",\"Domain\": \"" + MyDomain + "\"}'\n" +
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
		{"name": displayName},
		{"networkids": workspaceInfo.NetworkUuid},
	}

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getDeployVirtualMachine",
	}).Infof("deployVirtualMachine params [%v]", params)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getDeployVirtualMachine",
	}).Infof("deployVirtualMachine endUrl [%v]", endUrl)

	res := map[string]interface{}{}
	resp, err := http.Get(endUrl)
	if err != nil {

		log.WithFields(logrus.Fields{
			"communicationMold.go": "getDeployVirtualMachine",
		}).Errorf("Failed to communicate with Mold. (deployVirtualMachine) [%v]", err)

		res["message"] = "Failed to communicate with Mold. (deployVirtualMachine)"
		res["status"] = BaseErrorCode
	} else {
		if instanceType == WorkspaceString {
			updateWorkspaceTemplateCheck(workspaceInfo.Uuid, AgentCheck)
		}
		err = json.NewDecoder(resp.Body).Decode(&res)
		res["status"] = http.StatusOK
	}
	return res, instanceUuid
}

func getDestroyVirtualMachine(params []MoldParams) map[string]interface{} {
	var baseurl = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "destroyVirtualMachine"},
	}
	params = append(params, params1...)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getDestroyVirtualMachine",
	}).Infof("destroyVirtualMachine endUrl [%v]", endUrl)

	resp, err := http.Get(endUrl)

	if err != nil {
		log.WithFields(logrus.Fields{
			"communicationMold.go": "getDestroyVirtualMachine",
		}).Errorf("Failed to communicate with Mold. (deployVirtualMachine) [%v]", err)
	}
	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getCreateTags(params []MoldParams) map[string]interface{} {
	var baseurl = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "createTags"},
	}
	params = append(params, params1...)

	stringParams := makeStringParams(params)
	log.Infof("stringParams = [%v]", stringParams)
	sig := makeSignature(stringParams)

	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getCreateTags",
	}).Infof("createTags endUrl [%v]", endUrl)

	resp, err := http.Get(endUrl)
	if err != nil {
		log.WithFields(logrus.Fields{
			"communicationMold.go": "getCreateTags",
		}).Errorf("Failed to communicate with Mold. (createTags) [%v]", err)
	}
	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getStartVirtualMachine(params []MoldParams) map[string]interface{} {
	var baseurl = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "startVirtualMachine"},
	}
	params = append(params, params1...)

	stringParams := makeStringParams(params)
	log.Infof("stringParams = [%v]", stringParams)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getStartVirtualMachine",
	}).Infof("startVirtualMachine endUrl [%v]", endUrl)

	resp, err := http.Get(endUrl)
	if err != nil {
		log.WithFields(logrus.Fields{
			"communicationMold.go": "getCreateTags",
		}).Errorf("Failed to communicate with Mold. (startVirtualMachine) [%v]", err)
	}
	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getRebootVirtualMachine(params []MoldParams) map[string]interface{} {
	var baseurl = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "rebootVirtualMachine"},
	}
	params = append(params, params1...)

	stringParams := makeStringParams(params)
	log.Infof("stringParams = [%v]", stringParams)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getStartVirtualMachine",
	}).Infof("startVirtualMachine endUrl [%v]", endUrl)

	resp, err := http.Get(endUrl)
	if err != nil {
		log.WithFields(logrus.Fields{
			"communicationMold.go": "getCreateTags",
		}).Errorf("Failed to communicate with Mold. (startVirtualMachine) [%v]", err)
	}
	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getStopVirtualMachine(params []MoldParams) map[string]interface{} {
	var baseurl = os.Getenv("MoldUrl")
	params1 := []MoldParams{
		{"command": "stopVirtualMachine"},
	}
	params = append(params, params1...)

	stringParams := makeStringParams(params)
	log.Infof("stringParams = [%v]", stringParams)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getStopVirtualMachine",
	}).Infof("stopVirtualMachine endUrl [%v]", endUrl)

	resp, err := http.Get(endUrl)
	if err != nil {
		log.WithFields(logrus.Fields{
			"communicationMold.go": "getCreateTags",
		}).Errorf("Failed to communicate with Mold. (stopVirtualMachine) [%v]", err)
	}
	var res map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getListVirtualMachinesMetrics(params []MoldParams) map[string]interface{} {
	var baseurl = os.Getenv("MoldUrl")
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
	err = json.NewDecoder(resp.Body).Decode(&res)

	log.Debugf("getListVirtualMachinesMetrics result [%v]", res)

	return res
}

func getListVolumesMetrics(params []MoldParams) map[string]interface{} {
	var baseurl = os.Getenv("MoldUrl")
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
	err = json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func getListApis() (*http.Response, error) {
	var baseurl = os.Getenv("MoldUrl")
	params := []MoldParams{
		{"command": "listApis"},
	}
	client := http.Client{
		Timeout: 60 * time.Second,
	}

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getListApis",
	}).Infof("listApis params [%v]", params)

	stringParams := makeStringParams(params)
	sig := makeSignature(stringParams)
	endUrl := baseurl + "?" + stringParams + "&signature=" + sig

	log.WithFields(logrus.Fields{
		"communicationMold.go": "getListApis",
	}).Infof("listNetworks endUrl [%v]", endUrl)

	resp, err := client.Get(endUrl)

	return resp, err
}
