package main

import (
	"encoding/json"
	"net/http"
)

func selectApplicationServerListService(workspaceUuid string) []Application {
	workspaceList, err := selectWorkspaceList(workspaceUuid)
	if err != nil {

	}

	workspaceInfo := workspaceList[0]

	instanceList, err := selectInstanceList(workspaceInfo.Uuid, WorkspaceString)
	appList := []Application{}
	//serverAppInfo := map[string]interface{}{}
	instanceAppList, _ := getApplicationList(instanceList[0].Ipaddress)
	log.Errorf("instanceAppList [%v]", instanceAppList)
	for _, instanceAppItem := range instanceAppList {
		instanceAppItem.ApplicationServer = append(instanceAppItem.ApplicationServer, instanceList[0].Uuid)
		appList = append(appList, instanceAppItem)
	}
	if len(instanceList) >= 2 {
		for _, instanceInfo := range instanceList[1:] {
			instanceAppList, _ := getApplicationList(instanceInfo.Ipaddress)
			log.Errorf("instanceAppList [%v]", instanceAppList)
			for _, instanceAppItem := range instanceAppList {
				appFlag := true
				for key, appItem := range appList {

					if instanceAppItem.Path == appItem.Path {
						appList[key].ApplicationServer = append(appItem.ApplicationServer, instanceInfo.Uuid)
						appFlag = false
					}
				}
				if appFlag {
					instanceAppItem.ApplicationServer = append(instanceAppItem.ApplicationServer, instanceInfo.Uuid)
					appList = append(appList, instanceAppItem)
				}
			}
		}
	}

	//} else if appList[key2].Path == instanceAppItem.Path {
	//instanceAppItem.ApplicationServer = append(instanceAppItem.ApplicationServer, instanceInfo.Uuid)

	appListMarshal, _ := json.Marshal(appList)
	log.Warnf("applicationList [%s]", appListMarshal)

	return appList
}

func insertApplicationService(application Application) map[string]interface{} {
	resultInsertApplication := insertApplicationDao(application)
	returnData := map[string]interface{}{}
	if resultInsertApplication["status"] == http.StatusOK {
		returnData["status"] = resultInsertApplication["status"]
		returnData["message"] = resultInsertApplication["message"]
	}
	return returnData
}

func deleteApplicationService(application Application) map[string]interface{} {
	resultInsertApplication := insertApplicationDao(application)
	returnData := map[string]interface{}{}
	if resultInsertApplication["status"] == http.StatusOK {
		returnData["status"] = resultInsertApplication["status"]
		returnData["message"] = resultInsertApplication["message"]
	}
	return returnData
}

func selectApplicationListService(workspaceUuid string) []Application {
	//resultInsertApplication := insertApplicationDao(appInfo)
	applicationList, _ := selectApplicationDao(workspaceUuid)
	return applicationList
}
