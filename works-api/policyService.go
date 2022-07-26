package main

import "github.com/sirupsen/logrus"

func policyList(workspaceUuid string) []Policy {

	policyList, _ := selectDefaultPolicyList()
	workspacesPolicyList, _ := selectWorkspacesPolicyList(workspaceUuid)

	log.WithFields(logrus.Fields{
		"policyService": "policyList",
	}).Infof("policyList [%v], workspacesPolicyList [%v]", policyList, workspacesPolicyList)

	for _, workspacesPolicyInfo := range workspacesPolicyList {
		for i, policyInfo := range policyList {
			if workspacesPolicyInfo.Name == policyInfo.Name {
				policyList[i].Value = workspacesPolicyInfo.Value
			}
		}
	}

	return policyList
}

func updatePolicy(workspace Workspace, policy Policy) map[string]interface{} {

	deleteWorkspacesPolicy(policy)
	updateResult, _ := insertPolicy(policy)
	defaultPolicyInfo, _ := selectDefaultPolicyInfo(policy.Name)
	log.Debugf("workspace.Name [%v], policy.Value [%v]", workspace.Name, policy.Value)
	if policy.Name == "" {
		//TODO 정책 기간별 설정 필요
	}
	if defaultPolicyInfo.Value == policy.Value {
		deletePolicyDC(workspace.Name, defaultPolicyInfo)
	} else {
		insertPolicyDC(workspace.Name, defaultPolicyInfo)
	}

	return updateResult
}
