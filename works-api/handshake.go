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

func handshakeVdi(instance Instance, vdiType string) {
	vdiUrl := "http://" + instance.Ipaddress + ":8083/api"
	vdiName := instance.Name
	var DCInfo = os.Getenv("DCUrl")
	MyDomain := os.Getenv("SambaDomain")
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	for i := 0; i <= 120; i++ {
		statusString, err := getVdiAdStatus(vdiUrl)

		if statusString == Failed {
			time.Sleep(time.Second * 10)
			log.WithFields(logrus.Fields{
				"VDI IP": vdiUrl,
			}).Errorf("VDI communication failed. [%v], status [%v]", err, statusString)

		} else if statusString == Pending {
			updateInstanceHandshakeStatus(statusString, instance.Uuid)
			log.WithFields(logrus.Fields{
				"VDI IP": vdiUrl,
			}).Infof("VDI AD join status. [%v]", statusString)
			time.Sleep(time.Second * 30)
			reqPutAd, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%v/v1/ad/%v", vdiUrl, MyDomain), nil)
			if err != nil {
				log.Errorf("An error occurred while setting the URL. [%v]", err)
			}
			reqPutAd.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			resp, err1 := client.Do(reqPutAd)
			if err1 != nil {
				time.Sleep(time.Second * 30)
			} else if resp.Status == OK200 {
				log.WithFields(logrus.Fields{
					"VDI IP": vdiUrl,
				}).Infof("VDI AD join Success. [%v]", statusString)
				time.Sleep(time.Second * 30)
			}

		} else if statusString == Joining {
			updateInstanceHandshakeStatus(statusString, instance.Uuid)
			log.WithFields(logrus.Fields{
				"VDI IP": vdiUrl,
			}).Infof("VDI AD join status. [%v]", statusString)
			reqPutComputer, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%v/v1/computer/%v/%v", DCInfo, vdiName, instance.WorkspaceName), nil)
			if err != nil {
				log.Errorf("An error occurred while setting the URL. [%v]", err)
			}
			reqPutComputer.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			resp, err1 := client.Do(reqPutComputer)
			if err1 != nil {
				log.Errorf("VDI add workspaces group error [%v]", err1)
				time.Sleep(time.Second * 30)
			} else if resp.Status == OK200 {
				time.Sleep(time.Second * 30)
				log.WithFields(logrus.Fields{
					"VDI IP": vdiUrl,
				}).Infof("VDI add workspaces group success. [%v]", statusString)
			}

		} else if statusString == Joined {
			updateInstanceHandshakeStatus(statusString, instance.Uuid)
			log.WithFields(logrus.Fields{
				"VDI IP": vdiUrl,
			}).Infof("VDI AD join status. [%v]", statusString)
			reqGetGpupdate, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/v1/cmd/?timeout=60&cmd=%v", vdiUrl, url.QueryEscape("gpupdate /force")), nil)
			if err != nil {
				log.Errorf("An error occurred while setting the URL. [%v]", err)
			}
			reqGetGpupdate.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			resp, err1 := client.Do(reqGetGpupdate)
			log.Warnf("resp [%v], err1 [%v]", resp, err1)
			if err1 != nil {
				log.WithFields(logrus.Fields{
					"VDI IP": vdiUrl,
				}).Errorf("VDI gpupdate failed. [%v], err1 [%v]", statusString, err1)
				time.Sleep(time.Second * 30)
			} else if resp.Status == OK200 {
				log.WithFields(logrus.Fields{
					"VDI IP": vdiUrl,
				}).Infof("VDI gpupdate success. [%v], resp [%v]", statusString, resp)
				time.Sleep(time.Second * 30)
			} else {
				log.WithFields(logrus.Fields{
					"VDI IP": vdiUrl,
				}).Errorf("VDI gpupdate . [%v], resp [%v], err1 [%v]", statusString, resp, err1)
			}

		} else if statusString == Ready {
			updateInstanceHandshakeStatus(statusString, instance.Uuid)
			log.WithFields(logrus.Fields{
				"VDI IP": vdiUrl,
			}).Infof("VDI AD join status. [%v]", statusString)
			if vdiType == WorkspaceString {
				workspaceTemplateCheck := updateWorkspaceTemplateCheck(instance.WorkspaceUuid, AgentOK)
				if workspaceTemplateCheck["status"] == http.StatusOK {
					asyncJob := AsyncJob{}
					asyncJob.Id = getUuid()
					asyncJob.Name = VMDestroy
					asyncJob.ExecUuid = instance.Uuid
					asyncJob.Ready = 1
					resultInsertAsyncJob := insertAsyncJob(asyncJob)
					log.Infof("AsyncJob Insert Result [%v]", resultInsertAsyncJob)
					updateWorkspacePostfix(instance.WorkspaceUuid, 0)
				}
			}
			break
		}
	}
}

func getVdiAdStatus(vdiUrl string) (string, error) {
	client := http.Client{
		Timeout: 60 * time.Second,
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
