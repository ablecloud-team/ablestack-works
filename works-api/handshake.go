package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
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

func handshakeVdi(listVirtualMachinesMetrics ListVirtualMachinesMetrics, instance Instance) {
	vdiUrl := "http://" + listVirtualMachinesMetrics.Virtualmachine[0].Ipaddress + ":8083/api"
	vdiName := listVirtualMachinesMetrics.Virtualmachine[0].Displayname
	var DCInfo = os.Getenv("DCUrl")
	MyDomain := os.Getenv("SambaDomain")
	client := http.Client{
		Timeout: 60 * time.Second,
	}
	for i := 0; i <= 120; i++ {
		statusString, err := getVdiAdStatus(vdiUrl)

		if statusString == Failed {
			log.WithFields(logrus.Fields{
				"VDI IP": vdiUrl,
			}).Errorf("VDI communication failed. [%v]", err)
			log.Infof("VDI AD join status. [%v]", statusString)

		} else if statusString == Pending {
			updateInstanceHandshakeStatus(statusString, instance.Uuid)
			log.WithFields(logrus.Fields{
				"VDI IP": vdiUrl,
			}).Infof("VDI AD join status. [%v]", statusString)
			reqPutAd, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%v/v1/ad/%v", vdiUrl, MyDomain), nil)
			if err != nil {
				log.Errorf("An error occurred while setting the URL. [%v]", err)
			}
			reqPutAd.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			client.Do(reqPutAd)

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
			client.Do(reqPutComputer)

		} else if statusString == Joined {
			updateInstanceHandshakeStatus(statusString, instance.Uuid)
			log.WithFields(logrus.Fields{
				"VDI IP": vdiUrl,
			}).Infof("VDI AD join status. [%v]", statusString)
			reqGetGpupdate, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v/v1/cmd/?timeout=60&cmd=gpupdate /force", vdiUrl), nil)
			if err != nil {
				log.Errorf("An error occurred while setting the URL. [%v]", err)
			}
			reqGetGpupdate.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			client.Do(reqGetGpupdate)

		} else if statusString == Ready {
			updateInstanceHandshakeStatus(statusString, instance.Uuid)
			log.WithFields(logrus.Fields{
				"VDI IP": vdiUrl,
			}).Infof("VDI AD join status. [%v]", statusString)
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
