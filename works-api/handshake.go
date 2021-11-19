package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	Pending = "Pending"
	Failed  = "Failed"
	Joining = "Joining"
	Joined  = "Joined"
	Ready   = "Ready"
)

//listVirtualMachinesMetrics.Virtualmachine[0].Ipaddress

func handshakeVdi(instanceInfo Instance, vdiType string) {
	vdiUrl := "http://" + instanceInfo.Ipaddress + ":8083/api"
	vdiName := instanceInfo.Name
	var DCInfo = os.Getenv("DCUrl")
	MyDomain := os.Getenv("SambaDomain")
	client := http.Client{
		Timeout: 300 * time.Second,
	}
	for i := 0; i <= 120; i++ {
		// 중간에 인터벌을 안두면 for 문이 120번이 순간적으로 돌면서 handshake 가 종료됨
		time.Sleep(time.Second * 30)
		statusString, err := getVdiAdStatus(vdiUrl)

		if statusString == Failed {
			log.WithFields(logrus.Fields{
				"VDI_IP": vdiUrl,
			}).Errorf("VDI communication failed. [%v], status [%v]", err, statusString)

		} else if statusString == Pending {
			updateInstanceHandshakeStatus(instanceInfo.Uuid, statusString)
			if vdiType == WorkspaceString {
				updateWorkspaceTemplateCheck(instanceInfo.WorkspaceUuid, statusString)
			}
			log.WithFields(logrus.Fields{
				"VDI_IP": vdiUrl,
			}).Infof("VDI AD join status. [%v]", statusString)
			reqPutAd, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%v/v1/ad/%v", vdiUrl, MyDomain), nil)
			if err != nil {
				log.Errorf("An error occurred while setting the URL. [%v]", err)
			}
			reqPutAd.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			resp, err1 := client.Do(reqPutAd)
			if err1 != nil {
			} else if resp.Status == OK200 {
				log.WithFields(logrus.Fields{
					"VDI_IP": vdiUrl,
				}).Infof("VDI AD join Success. [%v]", statusString)
			}

		} else if statusString == Joining {
			updateInstanceHandshakeStatus(instanceInfo.Uuid, statusString)
			if vdiType == WorkspaceString {
				updateWorkspaceTemplateCheck(instanceInfo.WorkspaceUuid, statusString)
			}
			log.WithFields(logrus.Fields{
				"VDI_IP": vdiUrl,
			}).Infof("VDI AD join status. [%v]", statusString)
			reqPutComputer, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%v/v1/computer/%v/%v", DCInfo, vdiName, instanceInfo.WorkspaceName), nil)
			if err != nil {
				log.Errorf("An error occurred while setting the URL. [%v]", err)
			}
			reqPutComputer.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			resp, err1 := client.Do(reqPutComputer)
			if err1 != nil {
				log.Errorf("VDI add workspaces group error [%v]", err1)
			} else if resp.StatusCode == http.StatusOK {
				log.WithFields(logrus.Fields{
					"VDI_IP": vdiUrl,
				}).Infof("VDI add workspaces group success. [%v]", statusString)
			}

		} else if statusString == Joined {
			updateInstanceHandshakeStatus(instanceInfo.Uuid, statusString)
			if vdiType == WorkspaceString {
				updateWorkspaceTemplateCheck(instanceInfo.WorkspaceUuid, statusString)
			}
			log.WithFields(logrus.Fields{
				"VDI_IP": vdiUrl,
			}).Infof("VDI AD join status. [%v]", statusString)
			reqGetGpupdate, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/v1/cmd/?timeout=60&cmd=%v", vdiUrl, url.QueryEscape("gpupdate /force")), nil)
			if err != nil {
				log.Errorf("An error occurred while setting the URL. [%v]", err)
			}
			reqGetGpupdate.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			resp, err1 := client.Do(reqGetGpupdate)
			log.Warnf("resp [%v], err1 [%v]", resp.StatusCode, err1)
			if err1 != nil {
				log.WithFields(logrus.Fields{
					"VDI_IP": vdiUrl,
				}).Errorf("VDI gpupdate failed. [%v], err1 [%v]", statusString, err1)
			} else if resp.StatusCode == http.StatusOK {
				log.WithFields(logrus.Fields{
					"VDI_IP": vdiUrl,
				}).Infof("VDI gpupdate success. [%v], resp [%v]", statusString, resp)
			} else {
				log.WithFields(logrus.Fields{
					"VDI_IP": vdiUrl,
				}).Errorf("VDI gpupdate . [%v], resp [%v], err1 [%v]", statusString, resp, err1)
			}

		} else if statusString == Ready {
			updateInstanceHandshakeStatus(instanceInfo.Uuid, statusString)
			log.WithFields(logrus.Fields{
				"VDI_IP": vdiUrl,
			}).Infof("VDI AD join status. [%v]", statusString)
			if vdiType == WorkspaceString {
				workspaceTemplateCheck := updateWorkspaceTemplateCheck(instanceInfo.WorkspaceUuid, statusString)
				asyncJob := AsyncJob{}
				asyncJob.Id = getUuid()
				asyncJob.Name = VMDestroy
				asyncJob.ExecUuid = instanceInfo.Uuid
				asyncJob.Ready = 1
				resultInsertAsyncJob := insertAsyncJob(asyncJob)
				log.Infof("AsyncJob Insert Result [%v]", resultInsertAsyncJob)
				updateWorkspacePostfix(instanceInfo.WorkspaceUuid, 0)
				if workspaceTemplateCheck["status"] == http.StatusOK {
				}
			}
			break
		}
	}
}

func getVdiAdStatus(vdiUrl string) (string, error) {
	client := http.Client{
		Timeout: 300 * time.Second,
	}
	returnString := Failed
	res := map[string]interface{}{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/v1/ad/", vdiUrl), nil)
	if err != nil {
		log.Errorf("An error occurred while setting the URL. [%v]", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err1 := client.Do(req)

	if err1 != nil {

	} else if resp.Status == OK200 {

		respBody, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			log.Errorf("ioutil.ReadAll error [%v]", err2)
		}
		json.Unmarshal(respBody, &res)
		returnString = res["status"].(string)
	}

	return returnString, err1
}
