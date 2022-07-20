package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func selectWorkspaceInfo(workspaceUuid string) Workspace {

	workspaceList, _ := selectWorkspaceList(workspaceUuid)

	workspaceInfo := workspaceList[0]

	return workspaceInfo

}

func selectPublicPort(instanceInfo Instance, workspaceInfo Workspace) int {

	publicPort := selectPortForwardingNumber()

	var res map[string]interface{}

	paramsNetwork := []MoldParams{
		{"id": workspaceInfo.NetworkUuid},
	}
	networkResult := getNetwork(paramsNetwork)
	log.Infof("networkResult [%v]", networkResult)

	listNetworksResponse := ListNetworksResponse{}
	networksResponseInfo, _ := json.Marshal(networkResult["listnetworksresponse"])
	json.Unmarshal([]byte(networksResponseInfo), &listNetworksResponse)

	log.Infof("listNetworksResponse [%v]", listNetworksResponse.Network[0].Id)

	paramsListPublicIpAddresses := []MoldParams{
		{"associatednetworkid": listNetworksResponse.Network[0].Id},
	}

	listPublicIpqAddressesResponse := ListPublicIpAddressesResponse{}
	resultListPublicIpAddress, _ := getListPublicIpAddresses(paramsListPublicIpAddresses)
	json.NewDecoder(resultListPublicIpAddress.Body).Decode(&res)

	listPublicIpqAddressesResponseByte, _ := json.Marshal(res["listpublicipaddressesresponse"])
	json.Unmarshal([]byte(listPublicIpqAddressesResponseByte), &listPublicIpqAddressesResponse)

	log.Infof("listPublicIpqAddressesResponse [%v]", listPublicIpqAddressesResponse)

	//portForwardingNumber := selectPortForwardingNumber()

	policyList := policyList(workspaceInfo.Uuid)

	var rdpPort string

	for _, policyInfo := range policyList {
		if policyInfo.Name == "rdp_port" {
			rdpPort = policyInfo.Value
		}
	}

	paramsCreatePortForwardingRule := []MoldParams{
		{"ipaddressid": listPublicIpqAddressesResponse.Publicipaddress[0].Id},
		{"networkid": listNetworksResponse.Network[0].Id},
		{"privateport": rdpPort},
		{"protocol": "tcp"},
		{"publicport": strconv.Itoa(publicPort)},
		{"virtualmachineid": instanceInfo.MoldUuid},
		{"openfirewall": "true"},
	}

	resultgetCreatePortForwardingRule, _ := getCreatePortForwardingRule(paramsCreatePortForwardingRule)

	json.NewDecoder(resultgetCreatePortForwardingRule.Body).Decode(&res)
	//resultData["startvirtualmachineresponse"].(map[string]interface{})["jobid"].(string)
	//aaa := res["createportforwardingruleresponse"].(map[string]interface{})["id"].(string)
	portForwardingRuleId := res["createportforwardingruleresponse"].(map[string]interface{})["id"].(string)
	log.Infof("resultgetCreatePortForwardingRule [%v]", portForwardingRuleId)

	updatePortForwardingInstanceUuid(instanceInfo.Uuid, publicPort, portForwardingRuleId)

	//paramsCreateFirewallRule := []MoldParams{
	//	{"ipaddressid": listPublicIpqAddressesResponse.Publicipaddress[0].Ipaddress},
	//	{"protocol": "TCP"},
	//	{"startport": strconv.Itoa(publicPort)},
	//}
	//getCreateFirewallRule(paramsCreateFirewallRule)

	return publicPort
}

func deletePortForwading(instance Instance) {
	log.WithFields(logrus.Fields{
		"workspacesService": "deletePortForwading",
	}).Infof("params instance [%v]", instance)

	portForwardingRuleId := selectPortForwardingRuleId(instance)
	paramsDeletePortForwardingRule := []MoldParams{
		{"id": portForwardingRuleId},
	}

	getDeletePortForwardingRule(paramsDeletePortForwardingRule)
}

func checkRdpConnect(instanceInfo Instance) {
	log.WithFields(logrus.Fields{
		"workspacesService": "checkRdpConnect",
	}).Infof("params instance [%v]", instanceInfo)

	log.WithFields(logrus.Fields{
		"workspacesService": "checkRdpConnect",
	}).Infof("checkRdpConnect [%v]", "start")

	countNumber := 0
	for {
		log.WithFields(logrus.Fields{
			"workspacesService": "checkRdpConnect",
		}).Infof("checkRdpConnect [%v]", "for문 입장")
		time.Sleep(30 * time.Second)
		countRdpConnectInstances, _ := selectCountRdpConnectInstances(instanceInfo)
		log.WithFields(logrus.Fields{
			"workspacesService": "checkRdpConnect",
		}).Infof("countRdpConnectInstances [%v], countNumber[%v]", countRdpConnectInstances, countNumber)

		if countRdpConnectInstances == 0 {
			log.WithFields(logrus.Fields{
				"workspacesService": "checkRdpConnect",
			}).Infof("checkRdpConnect [%v]", "close")
			break
		} else {
			if countNumber == 10 {
				log.WithFields(logrus.Fields{
					"workspacesService": "checkRdpConnect",
				}).Infof("checkRdpConnect [%v]", "close")
				deletePortForwading(instanceInfo)
				updatePortForwardingInstanceUuidDelete(instanceInfo.Uuid)
				updateRdpConnected(instanceInfo, 0)
				break
			} else {
				countNumber += 1
			}
		}

	}
}
