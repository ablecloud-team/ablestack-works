package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func destroyWorkspaces(workspaceInfo Workspace) {
	ResultDeleteGroup, errDeleteGroup := deleteGroup(workspaceInfo.Name)
	ResultDeleteWorkspace := deleteWorkspace(workspaceInfo.Uuid)
	log.Infof("%v%v%v", ResultDeleteGroup, errDeleteGroup, ResultDeleteWorkspace)
}

func selectWorkspaceInfo(workspaceUuid string) Workspace {

	workspaceList, _ := selectWorkspaceList(workspaceUuid)

	workspaceInfo := workspaceList[0]

	workspacePolicyList, _ := selectWorkspacePolicyList(workspaceUuid)

	workspaceInfo.Policy.RdpPort = workspacePolicyList[0].Policy.RdpPort
	workspaceInfo.Policy.RdpAccessAllow = workspacePolicyList[0].Policy.RdpAccessAllow

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

	paramsCreatePortForwardingRule := []MoldParams{
		{"ipaddressid": listPublicIpqAddressesResponse.Publicipaddress[0].Id},
		{"networkid": listNetworksResponse.Network[0].Id},
		{"privateport": strconv.Itoa(workspaceInfo.Policy.RdpPort)},
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
