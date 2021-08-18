package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func makeStringParams(params map[string]string) string {
	var result string
	for key, value := range params {
		result = result + key + "=" + value + "&"
	}
	result = strings.Replace(result, "+", "%20", -1)
	result = strings.TrimRight(result, "&")
	//result = strings.TrimRight(strings.ToLower(result),"&")
	return result
}

func makeSignature(params map[string]string) string {
	//baseurl := "https://mold.ablecloud.io/client/api?"
	secretkey := os.Getenv("MoldSecretKey")
	//strurl := strings.ToLower(makeStringParams(params))
	strurl := makeStringParams(params)

	fmt.Println("secretkey = " + secretkey)
	fmt.Println("params =" + strurl)

	secret := []byte(secretkey)
	message := []byte(strurl)
	hash := hmac.New(sha1.New, secret)
	hash.Write(message)
	strHash := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	fmt.Println("strhash =")
	fmt.Println(strHash)

	return strHash
}
