package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
)

func makeSignature(params map[string]string ) string {
	baseurl := "https://mold.ablecloud.io/client/api?"
	secretkey := "64YHQCLtc0Ad_k-0NmeupaAwpsOJsirmVz52YSfYd-5RQs4F2deFen4O4SncOQ8I_M4Peew-JfMP2A2ZmKEnXg"

	var strurl string
	for key, value := range params{
		strurl = strurl + key+"="+value+"&"
	}
	strurl = strings.Replace(strurl, "+", "%20", -1)
	strurl = strings.TrimRight(strings.ToLower(strurl),"&")
	fmt.Println("baseurl = " + baseurl)
	fmt.Println("secretkey = " + secretkey)
	fmt.Println("params ="+strurl)

	secret := []byte(secretkey)
	message := []byte(strurl)
	hash := hmac.New(sha1.New, secret)
	hash.Write(message)
	strHash := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	fmt.Println("strhash =")
	fmt.Println(strHash)

	return strHash
}
