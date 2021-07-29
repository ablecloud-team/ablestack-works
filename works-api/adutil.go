package main

import (
	"fmt"
	ldap "github.com/go-ldap/ldap/v3"
	auth "github.com/korylprince/go-ad-auth/v3"
	//"github.com/sirupsen/logrus"

	//"github.com/sirupsen/logrus"
	"strings"
)

//var log = logrus.New().WithField("who", "AD")

var ADusername = "Administrator"
var ADpassword = "Ablecloud1!"
var domain = "dc1.local"
var ADserver = "dc1.local"
var ADport = 636
var ADbasedn = "DC=dc1,DC=local"


var username = "user"
var password = "Ablecloud1!"

var config = &auth.Config{
Server:   ADserver,
Port:     ADport,
BaseDN:   ADbasedn,
Security: auth.SecurityInsecureTLS,
}

func ConnectAD() (conn *auth.Conn, status bool, err error) {
	//connect
	conn, err = config.Connect()
	if err != nil {
		log.Errorln(err)
		return conn, false, err
	}

	//bind
	upn, err := config.UPN(ADusername)
	if err != nil {
		log.Errorln(err)
		upn = ADusername
	}
	status, err = conn.Bind(upn, ADpassword)
	if err != nil {
		log.Errorln(err)
		return conn, false, err
	}
	if !status {
		return conn, false, nil
	}
	return conn, status, err
}

func Auth(conn *auth.Conn, username string, password string) (status bool, err error){
	status, err = auth.Authenticate(config, username, password)
	return status, err
}
func listGroups(conn*auth.Conn, username string) (status bool, entry *ldap.Entry, userGroups []string, err error){
	//conn, status, err := ConnectAD()

	upn, err := config.UPN(username)
	if err != nil {
		log.Errorln(err)
		upn = username
	}
	log.Errorln(upn)
	entry, err = conn.GetAttributes("userPrincipalName", upn, []string{"cn"})
	if err != nil {
		log.Errorln(entry)
		log.Errorln(err)
		return false, nil, nil, err
	}
	log.Infof("entry")
	for _, attrs := range entry.Attributes{
		attrs.PrettyPrint(4)
	}
	log.Errorln(err)
	const LDAPMatchingRuleInChain = "1.2.840.113556.1.4.1941"

	foundGroups, err := conn.Search(fmt.Sprintf("(member:%s:=%s)", LDAPMatchingRuleInChain, ldap.EscapeFilter(entry.DN)), []string{""}, 1000)
	if err != nil {
		return false, nil, nil, err
	}


	log.Infoln("foundGroups")
	for _, group := range foundGroups {
		userGroups=append(userGroups, group.DN)
	}
	log.Errorln(err)
	return status, entry, userGroups, nil
}
func inGroup(conn *auth.Conn, username string, groupname string) (bool, error) {
	_, _, groups, err := listGroups(conn, username)
	if err != nil {
		return false, err
	}
	for _, group := range groups{
		if strings.Compare(strings.TrimSpace(strings.ToLower(groupname)), strings.TrimSpace(strings.ToLower(group))) == 0{
			return true, nil
		}
	}

	return false, nil
}
/*
func InGroup(username string, groupname string){
	status, entry, groups, err := auth.AuthenticateExtended(config, fmt.Sprintf("%v", username), password, []string{"cn"}, []string{groupname})

	if err != nil {
		fmt.Println("handle error")
		fmt.Println(err)
	}

	if !status {
		fmt.Println("auth failed")
		return
	}

	fmt.Println(groups)
	if len(groups) == 0 {
		//handle user not being in any groups
		return
	}

	//get attributes
	cn := entry.GetAttributeValue("cn")

	fmt.Println(cn)
}
*/