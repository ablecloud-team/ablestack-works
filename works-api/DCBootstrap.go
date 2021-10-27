package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func dcBootstrap() {
	oneshot := true

	for oneshot {
		log.Infof("1st bootstrap start")
		time.Sleep(10 * time.Second)
		var DCInfo = os.Getenv("DCUrl")
		log.Infof("DCInfo [%v]", DCInfo)
		client := http.Client{
			Timeout: 5 * time.Second,
		}
		resp, err := client.Get(DCInfo + "/v1/version")
		res := map[string]interface{}{}

		if err != nil {
			log.Errorf("Failed to communicate with DC Server. [%v]", err)
		} else if resp.Status == OK200 {
			json.NewDecoder(resp.Body).Decode(&res)
			if res["Bootstraped"] == false {
				break
				log.Infof("1st bootstrap end")
			} else {
				oneshot = false
				log.Infof("1st bootstrap end & bootstraped true")
				break
			}
			break
		}
	}

	for oneshot {
		log.Infof("2nd bootstrap start")
		time.Sleep(10 * time.Second)
		var DCInfo = os.Getenv("DCUrl")

		client := http.Client{
			Timeout: 5 * time.Second,
		}
		r, _ := http.NewRequest("PATCH", DCInfo+"/v1", nil) // URL-encoded payload

		r.Header.Add("Content-Type", "application/json")
		res := map[string]interface{}{}
		resp, err := client.Do(r)
		defer resp.Body.Close()
		if err != nil {
			log.Errorf("err [%v]", err)
		} else if resp.Status == OK200 {
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Errorf("resp1 [%v], err [%v]", resp.Body, err)
			}
			json.Unmarshal(respBody, &res)
			log.Infof("2nd bootstrap end")
			time.Sleep(60 * time.Second)
			break
		}
	}

	for oneshot {
		log.Infof("3nd bootstrap start")
		time.Sleep(10 * time.Second)
		var DCInfo = os.Getenv("DCUrl")
		log.Infof("DCInfo [%v]", DCInfo)
		client := http.Client{
			Timeout: 5 * time.Second,
		}
		resp, err := client.Get(DCInfo + "/v1/version")
		if err != nil {
			log.Errorf("Failed to communicate with DC Server. [%v]", err)
		} else if resp.Status == OK200 {
			log.Infof("3nd bootstrap end")
			break
		}
	}

	for oneshot {
		log.Infof("4th bootstrap start")
		time.Sleep(10 * time.Second)
		var DCInfo = os.Getenv("DCUrl")

		client := http.Client{
			Timeout: 60 * time.Second,
		}
		r, _ := http.NewRequest("PATCH", DCInfo+"/v1/policy", nil) // URL-encoded payload

		r.Header.Add("Content-Type", "application/json")
		res := map[string]interface{}{}
		resp, err := client.Do(r)
		defer resp.Body.Close()
		log.Infof("resp.Status [%v], resp.Body [%v]", resp.Status, resp.Body)
		if err != nil {
			log.Errorf("err [%v]", err)
		} else if resp.Status == OK200 {
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Errorf("resp1 [%v], err [%v]", resp.Body, err)
			}
			json.Unmarshal(respBody, &res)
			log.Infof("4th bootstrap end")
			break
		}
	}
}
