package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func dcBootstrap() {
	for {
		log.Infof("1st bootstrap start")
		time.Sleep(10 * time.Second)
		var DCInfo = os.Getenv("DCUrl")
		log.Infof("DCInfo [%v]", DCInfo)
		client := http.Client{
			Timeout: 10 * time.Second,
		}
		resp, err := client.Get(DCInfo + "/v1/version")
		res := map[string]interface{}{}

		if err != nil {
			log.Errorf("Failed to communicate with DC Server. [%v]", err)
		} else if resp.Status == OK200 {
			json.NewDecoder(resp.Body).Decode(&res)
			if res["Bootstraped"] == false {
				log.Infof("1st bootstrap end")
				r, _ := http.NewRequest("PATCH", DCInfo+"/v1", nil) // URL-encoded payload

				r.Header.Add("Content-Type", "application/json")
				resp, _ := client.Do(r)
				log.Infof("%v", resp)
				time.Sleep(60 * time.Second)
				break
			} else {
				log.Infof("1st bootstrap end & bootstraped true")
				break
			}
			break
		}
	}

	for {
		client := http.Client{
			Timeout: 10 * time.Second,
		}
		var DCInfo = os.Getenv("DCUrl")
		r1, _ := http.NewRequest("GET", DCInfo+"/v1/policy", nil) // URL-encoded payload
		resp1, err := client.Do(r1)
		if err != nil {
			log.Errorf("%v/v1/policy err [%v]",DCInfo, err)
		} else {
			respBody1, err1 := ioutil.ReadAll(resp1.Body)
			if err1 != nil {
				log.Errorf("[%v] [%v]",respBody1, err1)
			}else if string(respBody1) == "[]" {
				log.Errorf("%v/v1/policy err [%v]",DCInfo, err)
				r, _ := http.NewRequest("PATCH", DCInfo+"/v1/policy", nil) // URL-encoded payload

				r.Header.Add("Content-Type", "application/json")
				resp, err := client.Do(r)
				log.Infof("%v, %v",resp, err)
			} else {
				log.Errorf("%v/v1/policy respBody [%v]",DCInfo, respBody1)
				log.Infof("4th bootstrap end")
				break
			}
		}
	}
	//for {
	//	log.Infof("4th bootstrap start")
	//	time.Sleep(10 * time.Second)
	//	var DCInfo = os.Getenv("DCUrl")
	//
	//	client := http.Client{
	//		Timeout: 60 * time.Second,
	//	}
	//	r, _ := http.NewRequest("PATCH", DCInfo+"/v1/policy", nil) // URL-encoded payload
	//
	//	r.Header.Add("Content-Type", "application/json")
	//	res := map[string]interface{}{}
	//	resp, err := client.Do(r)
	//	defer resp.Body.Close()
	//	log.Infof("resp.Status [%v], resp.Body [%v]", resp.Status, resp.Body)
	//	if err != nil {
	//		log.Errorf("err [%v]", err)
	//	} else if resp.Status == OK200 {
	//		respBody, err := ioutil.ReadAll(resp.Body)
	//		if err != nil {
	//			log.Errorf("resp1 [%v], err [%v]", resp.Body, err)
	//		}
	//		json.Unmarshal(respBody, &res)
	//		r1, _ := http.NewRequest("GET", DCInfo+"/v1/policy", nil) // URL-encoded payload
	//		resp1, err := client.Do(r1)
	//		respBody1, _ := ioutil.ReadAll(resp1.Body)
	//		if string(respBody1) == "null" {
	//			log.Errorf("%v/v1/policy err [%v]",DCInfo, err)
	//		} else {
	//			log.Errorf("%v/v1/policy respBody [%v]",DCInfo, respBody1)
	//			log.Infof("4th bootstrap end")
	//			break
	//		}
	//	}
	//}
}
