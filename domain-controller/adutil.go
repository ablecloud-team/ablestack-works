package main

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	auth "github.com/korylprince/go-ad-auth/v3"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	//"github.com/sirupsen/logrus"

	//"github.com/sirupsen/logrus"
	"strings"
)

//var log = logrus.New().WithField("who", "AD")

type ADConfig struct {
	ADusername string
	ADpassword string
	ADdomain   string
	ADserver   string
	ADport     int
	ADbasedn   string
	Silent     bool
	PolicyPATH string
	PolicyLIST string
	UpdatePolicy bool
	BootStraped bool
}

var ADconfig = ADConfig{}
var UserAttributes = []string{
	"dn",
	"distinguishedName",
	"cn",
	"sn",
	"name",
	"accountname",
	"givenName",
	"initials",
	"userPrincipalName",
	"username",
	"sAMAccountName",
	"description",
	"info",
	"title",
	"o",
	"company",
	"postOfficeBox",
	"physicalDeliveryOfficeName",
	"streetAddress",
	"I",
	"st",
	"department",
	"mail",
	"telephoneNumber",
	"pager",
	"mobile",
	"facsimileTelephoneNumber",
	"homePhone",
	"ipPhone",
	"postalCode",
	"manager",
	"wWWHomePage",
	"countryCode",
	"c",
	"co",
	"memberOf",
	"member",
}
var GroupAttributes = []string{
	"dn",
	"cn",
	"distinguishedName",
	"member",
}

//
//var username = "user"
//var password = "Ablecloud1!"

var authconfig = &auth.Config{}

func ADsave() (err error) {

	//ADconfig.ADusername = "Administrator"
	//ADconfig.ADpassword = "Ablecloud1!"
	//ADconfig.ADdomain = "dc1.local"
	//ADconfig.ADserver = "dc1.local"
	//ADconfig.ADport = 636
	//ADconfig.ADbasedn = "DC=dc1,DC=local"
	//ADconfig.PolicyPATH = "./grouppolicy"
	//ADconfig.PolicyLIST = "policylist.json"
	//ADconfig.UpdatePolicy = false
	//ADconfig.Silent = false

	byteValue, err := json.MarshalIndent(ADconfig, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("conf.json", byteValue, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ADinit() (err error) {
	conffile:="conf.json"
	data, err := os.Open(conffile)
	if err != nil {
		log.Fatalf("Can not find %v file", conffile)
		return err
	}

	byteValue, err := ioutil.ReadAll(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &ADconfig)
	if err != nil {
		return err
	}

	authconfig = &auth.Config{
		Server:   ADconfig.ADserver,
		Port:     ADconfig.ADport,
		BaseDN:   ADconfig.ADbasedn,
		Security: auth.SecurityInsecureTLS,
	}


	err2 := ""
	currentWorkingDirectory, err := os.Getwd()
	policyfile := fmt.Sprintf("%v/%v/%v",currentWorkingDirectory, ADconfig.PolicyPATH, ADconfig.PolicyLIST)
	data, err = os.Open(policyfile)
	if err != nil {
		log.Fatalf("Can not find %v file, %v", policyfile, err)
		return err
	}

	byteValue, err = ioutil.ReadAll(data)
	if err != nil {
		return err
	}
	var policyList []map[string]string
	err = json.Unmarshal(byteValue, &policyList)
	if err != nil {
		return err
	}

	if ADconfig.UpdatePolicy != ADconfig.UpdatePolicy {
		for _, policyItem := range policyList {
			shell, err := setupShell()
			if err != nil {
				log.Errorf("%v", err)
				return err
			}
			policy := policyItem["name"]
			description := policyItem["description"]
			stdout, err := shell.Exec(fmt.Sprintf("c:\\Works-DC\\psexec.exe -accepteula -u %v\\%v -p %v \"powershell import-gpo -BackupGpoName %v -TargetName %v -Path '%v/%v' -CreateIfNeeded\"", ADconfig.ADdomain, ADconfig.ADusername, ADconfig.ADpassword, policy, policy, currentWorkingDirectory, ADconfig.PolicyPATH))
			log.Infof("stdout: %v, \nstderr: %v\n", stdout, err)
			if err != nil {
				err2 = fmt.Sprintf("%v, %v", err2, err)
			}
			shell.Exec(fmt.Sprintf("$policy = Get-Gpo -Name '%v'", policy))
			shell.Exec(fmt.Sprintf("$policy.Description = '%v' ", description))

		}
	}
	if err2 == ""{
		return nil
	}
	return errors.New(err2)

}
func ConnectAD() (conn *auth.Conn, status bool, err error) {
	setLog()
	//connect
	setLog(fmt.Sprintf("ADconfig: %v", ADconfig))
	conn, err = authconfig.Connect()
	if err != nil {
		//log.Errorln(err)
		return conn, false, err
	}

	setLog(fmt.Sprintf("ADconfig: %v", ADconfig))
	//bind
	upn, err := authconfig.UPN(ADconfig.ADusername)
	if err != nil {
		//log.Errorln(err)
		upn = ADconfig.ADusername
	}
	status, err = conn.Bind(upn, ADconfig.ADpassword)
	if err != nil {
		log.Errorln(err)
		return conn, false, err
	}
	if !status {
		log.Errorf("Connection Failed with %v %v", upn, ADconfig.ADpassword)
		return conn, false, nil
	}
	return conn, status, err
}

func Auth(conn *auth.Conn, username string, password string) (status bool, err error) {
	status, err = auth.Authenticate(authconfig, username, password)
	conn = conn
	return status, err
}
func listUserGroups(conn *auth.Conn, username string) (status bool, entry *ldap.Entry, userGroups []string, err error) {
	setLog()
	upn, err := authconfig.UPN(username)
	if err != nil {
		//log.Errorln(err)
		upn = username
	}
	entry, err = conn.GetAttributes("userPrincipalName", upn, []string{"cn"})
	if err != nil {
		//log.Errorln(entry)
		//log.Errorln(err)
		entry, err = conn.GetAttributes("cn", username, []string{"cn"})
		if err != nil {
			//log.Errorln(entry)
			//log.Errorln(err)
			return false, nil, nil, err
		}
	}
	for _, attrs := range entry.Attributes {
		attrs.PrettyPrint(4)
	}
	const LDAPMatchingRuleInChain = "1.2.840.113556.1.4.1941"

	foundGroups, err := conn.Search(fmt.Sprintf("(member:%s:=%s)", LDAPMatchingRuleInChain, ldap.EscapeFilter(entry.DN)), []string{""}, 1000)
	if err != nil {
		return false, nil, nil, err
	}

	//log.Infoln("foundGroups")
	for _, group := range foundGroups {
		userGroups = append(userGroups, group.DN)
	}
	//log.Infoln(userGroups)
	return status, entry, userGroups, nil
}
func inGroup(conn *auth.Conn, username string, groupname string) (bool, error) {
	setLog()
	_, _, groups, err := listUserGroups(conn, username)
	if err != nil {
		return false, err
	}
	for _, group := range groups {
		//log.Infof("%v, %v\n",strings.TrimSpace(strings.ToLower(groupname)), strings.TrimSpace(strings.ToLower(group)))
		if strings.Contains(strings.TrimSpace(strings.ToLower(group)), strings.TrimSpace(strings.ToLower(groupname))) {
			return true, nil
		}
	}

	return false, nil
}

func setupLdap() (l *ldap.Conn, err error) {
	ldapsServer := fmt.Sprintf("ldaps://%v:%v", ADconfig.ADserver, ADconfig.ADport)
	//ldapServer := fmt.Sprintf("ldap://%v:%v", ADconfig.ADserver, ADconfig.ADport)
	l, err = ldap.DialURL(ldapsServer, ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	if err != nil {
		//log.Fatal(err)
	}

	upn, err := authconfig.UPN(ADconfig.ADusername)
	if err != nil {
		//log.Errorln(err)
		upn = ADconfig.ADusername
	}
	err = l.Bind(upn, ADconfig.ADpassword)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	return
}

func testLDAP() {
	setLog()
	// https://cybernetist.com/2020/05/18/getting-started-with-go-ldap/
	var filter = []string{
		"(cn=user)",
		"(cn=user2)",
		"(&(owner=*)(cn=user))",
		"(&(objectclass=rfc822mailgroup)(cn=*Computer*))",
		"(&(objectclass=rfc822mailgroup)(cn=*Mathematics*))"}
	var attributes = []string{
		"dn",
		"cn",
		"sn",
		"givenName",
		"initials",
		"userPrincipalName",
		"username",
		"sAMAccountName",
		"description",
		"info",
		"title",
		"o",
		"company",
		"postOfficeBox",
		"physicalDeliveryOfficeName",
		"streetAddress",
		"I",
		"st",
		"department",
		"mail",
		"telephoneNumber",
		"pager",
		"mobile",
		"facsimileTelephoneNumber",
		"homePhone",
		"ipPhone",
		"postalCode",
		"manager",
		"wWWHomePage",
		"countryCode",
		"c",
		"co",
	}

	l, _ := setupLdap()

	searchRequest := ldap.NewSearchRequest(
		ADconfig.ADbasedn,
		ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false,
		filter[0],
		attributes,
		nil)

	sr, err := l.Search(searchRequest)
	if err != nil {
		//log.Fatal(err)
	}
	//log.Infof("TestSearch: %s -> num of entries = %d", searchRequest.Filter, len(sr.Entries))
	sr.PrettyPrint(4)

}

//add group
func addGroup(l *ldap.Conn, groupname string) (err_ret error) {

	existOU := false
	existCN := false
	var (
		err  error
		err2 error
	)
	//ou add
	addReq := ldap.NewAddRequest(fmt.Sprintf("ou=%v,%v", groupname, ADconfig.ADbasedn), []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "organizationalUnit"})
	addReq.Attribute("name", []string{groupname})
	addReq.Attribute("ou", []string{groupname})
	addReq.Attribute("instanceType", []string{fmt.Sprintf("%d", 0x00000004)})
	if err = l.Add(addReq); err != nil {
		log.Error("error adding OU:", addReq, err)

		if strings.Contains(err.Error(), "Already Exists") {
			existOU = true
		}
	}
	//group add
	addReq = ldap.NewAddRequest(fmt.Sprintf("cn=%v,ou=%v,%v", groupname, groupname, ADconfig.ADbasedn), []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "group"})
	addReq.Attribute("name", []string{groupname})
	addReq.Attribute("sAMAccountName", []string{groupname})
	addReq.Attribute("instanceType", []string{fmt.Sprintf("%d", 0x00000004)})
	addReq.Attribute("groupType", []string{fmt.Sprintf("%d", 0x00000004|0x80000000)})
	if err2 = l.Add(addReq); err2 != nil {
		log.Error("error adding group:", addReq, err2)
		if strings.Contains(err2.Error(), "Already Exists") {
			existCN = true
		}
	}
	if existOU {
		if existCN {
			return errors.New("Entry Already Exists, OU, CN")
		} else {
			if err2 != nil {
				return errors.New(fmt.Sprintf("Entry Already Exists, OU and error %v", err2))
			} else {
				return errors.New("Entry Already Exists, OU")
			}
		}
	} else {
		if existCN {
			return errors.New("Entry Already Exists, CN")
		} else {
			if err != nil {
				return err
			}
			return nil
		}
	}
}


//add Connection
func addConnection(l *ldap.Conn, user ADUser, connection string, guacparameter []string) (err_ret error) {

	var (
		err  error
	)
	//guacparameter = []string{"hostname=10.1.1.18", "port=3389", "ignore-cert=true", "enable-drive=true ", "create-drive-path=true ", "username=Administrator", "password=Ablecloud1!", "domain=TEST", "drive-name=G", "drive-path=/share"}
	retuser, err := getUser(l, &USER{Username: user.username})
	//ou add
	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%v,cn=Users,%v", connection, ADconfig.ADbasedn), []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "guacConfigGroup", "groupOfNames"})
	addReq.Attribute("name", []string{connection})
	addReq.Attribute("guacConfigProtocol", []string{"rdp"})
	addReq.Attribute("guacConfigParameter", guacparameter)
	addReq.Attribute("member", []string{retuser["distinguishedName"].(string)})

	if err = l.Add(addReq); err != nil {
		log.Errorf("error adding connection:%v [%v]", addReq, err)

		if strings.Contains(err.Error(), "Already Exists") {
			log.Errorf("[%v]", "Already Exists")
		}
	}

	return err
}

//add Connection
func delConnection(l *ldap.Conn, connection string) (err_ret error) {

	var (
		err  error
	)
	//guacparameter = []string{"hostname=10.1.1.18", "port=3389", "ignore-cert=true", "enable-drive=true ", "create-drive-path=true ", "username=Administrator", "password=Ablecloud1!", "domain=TEST", "drive-name=G", "drive-path=/share"}
	delReq := ldap.NewDelRequest(fmt.Sprintf("cn=%v,cn=Users,%v", connection, ADconfig.ADbasedn), []ldap.Control{})


	if err = l.Del(delReq); err != nil {
		log.Errorf("error Deleting connection:%v [%v]", delReq, err)

		if strings.Contains(err.Error(), "Already Exists") {
			log.Errorf("[%v]", "Already Exists")
		}
	}

	return err
}

//get group info
func getGroup(l *ldap.Conn, group *GROUP) (retgroup ADGROUP, err error) {
	setLog()
	// https://cybernetist.com/2020/05/18/getting-started-with-go-ldap/

	l, _ = setupLdap()

	searchRequest := ldap.NewSearchRequest(
		ADconfig.ADbasedn,
		ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false,
		fmt.Sprintf("(&(cn=%v)(objectclass=group))", group.Groupname),
		GroupAttributes,
		nil)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return nil, err
	}
	if len(sr.Entries) >= 1 {
		groupEntry := sr.Entries[0]
		log.Infof("TestSearch: %s -> num of entries = %d", searchRequest.Filter, len(sr.Entries))

		adgroup := ADGROUP{} //NewADUser()
		log.Infoln(adgroup)
		for i := range groupEntry.Attributes {
			log.Infoln(i)
			log.Infoln(groupEntry.Attributes[i].Name)
			log.Infoln(groupEntry.Attributes[i].Values)
			if groupEntry.Attributes[i].Name == "member"{
				adgroup[groupEntry.Attributes[i].Name] = groupEntry.Attributes[i].Values
			} else {
				if len(groupEntry.Attributes[i].Values) >= 2 {
					adgroup[groupEntry.Attributes[i].Name] = groupEntry.Attributes[i].Values
				} else {
					adgroup[groupEntry.Attributes[i].Name] = groupEntry.Attributes[i].Values[0]
				}
			}
		}
		ret, err := json.Marshal(adgroup)
		log.Infoln(string(ret))
		return adgroup, err
	} else {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}
}

//list group
func searchGroup(l *ldap.Conn) (retusers []ABLEGROUP) {
	setLog()
	// https://cybernetist.com/2020/05/18/getting-started-with-go-ldap/

	l, _ = setupLdap()

	searchGroupRequest := ldap.NewSearchRequest(
		ADconfig.ADbasedn,
		ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false,
		"(&(objectCategory=group)(objectClass=group))",
		UserAttributes,
		nil)

	searchGroupResult, err := l.Search(searchGroupRequest)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("TestSearch: %s -> num of entries = %d", searchGroupRequest.Filter, len(searchGroupResult.Entries))
	var adGroups = []ADGROUP{}
	for _, groupEntry := range searchGroupResult.Entries {
		adGroup := ADGROUP{}
		for i := range groupEntry.Attributes {
			if len(groupEntry.Attributes[i].Values) >= 2 {
				adGroup[groupEntry.Attributes[i].Name] = groupEntry.Attributes[i].Values
			} else {
				adGroup[groupEntry.Attributes[i].Name] = groupEntry.Attributes[i].Values[0]
			}
		}
		adGroups = append(adGroups, adGroup)
	}

	searchOURequest := ldap.NewSearchRequest(
		ADconfig.ADbasedn,
		ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false,
		//"(&(objectCategory=Organizational-Unit)(objectClass=organizationalUnit))",
		"(objectClass=organizationalUnit)",
		UserAttributes,
		nil)

	searchOUResult, err := l.Search(searchOURequest)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("TestSearch: %s -> num of entries = %d", searchOURequest.Filter, len(searchOUResult.Entries))
	var adOUs = []ADOU{}
	for _, OUEntry := range searchOUResult.Entries {
		adOU := ADOU{}
		for i := range OUEntry.Attributes {
			if len(OUEntry.Attributes[i].Values) >= 2 {
				adOU[OUEntry.Attributes[i].Name] = OUEntry.Attributes[i].Values
			} else {
				adOU[OUEntry.Attributes[i].Name] = OUEntry.Attributes[i].Values[0]
			}
		}
		adOUs = append(adOUs, adOU)
	}

	var retResults []ABLEGROUP
	for _, valueG := range adGroups {
		for _, valueO := range adOUs {
			if valueO["name"] == valueG["name"] {
				retResult := ABLEGROUP{}
				for key, value := range valueO {
					retResult[key] = value
				}
				for key, value := range valueG {
					retResult[key] = value
				}
				retResults = append(retResults, retResult)
			}
		}
	}
	//Organizational-Unit
	ret, err := json.Marshal(adGroups)
	log.Infoln(string(ret))
	ret, err = json.Marshal(adOUs)
	log.Infoln(string(ret))
	return retResults
}

//user mod
func modGroup(l *ldap.Conn, user ADGROUP) (group_ ADGROUP, err error) {
	if val, ok := user["groupname"]; !ok || val == "" {
		return group_, errors.New("no group name")
	}
	modReq := ldap.NewModifyRequest(fmt.Sprintf("cn=%v,OU=%v,%v", user["groupname"], user["groupname"], ADconfig.ADbasedn), []ldap.Control{})

	group_ = NewADGroup(user).ToMap()

	for key, val := range group_ {
		switch val.(type) {
		case []string:
			if key != "groupname" && key != "memberOf" && key != "distinguishedName" {
				log.Infof("%v:%v", key, len(val.([]string)))
				modReq.Replace(key, val.([]string))
			}
		case int:
			i := val.(int)
			modReq.Replace(key, []string{strconv.Itoa(i)})
		case string:
			if key != "groupname" && val.(string) != "" && key != "distinguishedName" {
				modReq.Replace(key, []string{val.(string)})
			}
		default:
			log.Errorf("%v:%T, %v", key, val, val)
		}
	}
	err = l.Modify(modReq)
	if err != nil {
		log.Errorf("error moding group: %v, %v", modReq, err)
	}
	return group_, err
}

//delete group
func delGroup(l *ldap.Conn, groupname string) (err error) {
	log.Debugf("groupname : %v", groupname)
	if groupname == "" {
		return errors.New("no groupname name")
	}

	log.Debugf("groupname : %v", groupname)
	delreq := ldap.NewDelRequest(fmt.Sprintf("cn=%v,ou=%v,%v", groupname, groupname, ADconfig.ADbasedn), []ldap.Control{})

	if err := l.Del(delreq); err != nil {
		log.Error("error deleting groupname:", delreq, err)
		return err
	}

	delreq = ldap.NewDelRequest(fmt.Sprintf("ou=%v,%v", groupname, ADconfig.ADbasedn), []ldap.Control{})

	if err := l.Del(delreq); err != nil {
		log.Error("error deleting groupname:", delreq, err)
		return err
	}
	return err
}

//add user to group
func addUserToGroup(l *ldap.Conn, user ADUSER, group ADGROUP) (group_ ADGROUP, err error) {
	//group.member=userDN
	if val, ok := group["groupname"]; !ok || val == "" {
		return group_, errors.New("no group name")
	}

	if val, ok := user["username"]; !ok || val == "" {
		return group_, errors.New("no user name")
	} else {
		var user_ = &USER{Username: user["username"].(string)}
		user, err = getUser(l, user_)
		if err != nil {
			return nil, err
		}
	}

	modReq := ldap.NewModifyRequest(fmt.Sprintf("cn=%v,OU=%v,%v", group["groupname"], group["groupname"], ADconfig.ADbasedn), []ldap.Control{})

	group_ = NewADGroup(group).ToMap()

	modReq.Add("member", []string{user["distinguishedName"].(string)})
	err = l.Modify(modReq)
	if err != nil {
		log.Errorf("error moding group: %v, %v", modReq, err)
	}
	retgroup,_ := getGroup(l, &GROUP{Groupname: group["groupname"].(string)})
	return retgroup, err
}

//delete user from group
func deleteUserFromGroup(l *ldap.Conn, user ADUSER, group ADGROUP) (group_ ADGROUP, err error) {
	//group.member=userDN
	if val, ok := group["groupname"]; !ok || val == "" {
		return group_, errors.New("no group name")
	}

	if val, ok := user["username"]; !ok || val == "" {
		return group_, errors.New("no user name")
	} else {
		var user_ = &USER{Username: user["username"].(string)}
		user, err = getUser(l, user_)
		if err != nil {
			return nil, err
		}
	}

	modReq := ldap.NewModifyRequest(fmt.Sprintf("cn=%v,OU=%v,%v", group["groupname"], group["groupname"], ADconfig.ADbasedn), []ldap.Control{})

	group_ = NewADGroup(group).ToMap()

	modReq.Delete("member", []string{user["distinguishedName"].(string)})
	err = l.Modify(modReq)
	if err != nil {
		log.Errorf("error moding group: %v, %v", modReq, err)
	}
	return group_, err
}

//add user
func addUser(l *ldap.Conn, user ADUSER) (err error) {
	if val, ok := user["username"]; !ok || val == "" {
		return errors.New("no user name")
	}

	user["userPrincipalName"] = fmt.Sprintf("%v@%v", user["username"], ADconfig.ADdomain) //로그온 이름(accountname@ADdomain 형식)

	user["sAMAccountName"] = user["username"] //windows 2000 이전 사용자 로그온 이름(ADdomain\sAMAccountName 형식)

	user["countryCode"] = 410 //주소->국가(한국?)
	user["c"] = "KR"          //국가명
	user["co"] = user["c"]

	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%v,cn=%v,%v", user["username"], "Users", ADconfig.ADbasedn), []ldap.Control{})

	for key, val := range user {
		log.Infof("%v:%v", key, val)
		switch val.(type) {
		case string:
			if key != "username" && key != "distinguishedName" {
				if val != "" {
					addReq.Attribute(key, []string{val.(string)})
				}
			}
		case int:
			i := val.(int)
			addReq.Attribute(key, []string{strconv.Itoa(i)})
		default:
			log.Errorf("%v:%T, %v", key, val, val)
		}
	}

	log.Infof("Add user Sucess")
	addReq.Attribute("objectClass", []string{"top", "organizationalPerson", "person", "user"})
	if err := l.Add(addReq); err != nil {
		log.Errorf("[error adding user: %v] %v", err,addReq)
		return err
	}
	log.Infof("Add user Sucess")

	log.Infof("[adding user to group:%v, %v]", fmt.Sprintf("cn=%v,cn=Users,%v", user["username"], ADconfig.ADbasedn), "Domain Admins")
	modReq := ldap.NewModifyRequest(fmt.Sprintf("CN=Domain Admins,CN=Users,%v",  ADconfig.ADbasedn), []ldap.Control{})
	modReq.Add("member", []string{fmt.Sprintf("cn=%v,cn=Users,%v", user["username"], ADconfig.ADbasedn)})
	err = l.Modify(modReq)
	if err != nil {
		log.Errorf("error moding group: %v, %v", modReq, err)
	}
	//ldap.NewDelRequest(fmt.Sprintf("cn=%v,cn=%v,%v", user["username"], "Users", ADconfig.ADbasedn), []ldap.Control{})
	return err
}

//get user info
func getUser(l *ldap.Conn, user *USER) (retuser ADUSER, err error) {
	setLog()
	// https://cybernetist.com/2020/05/18/getting-started-with-go-ldap/

	l, _ = setupLdap()

	searchRequest := ldap.NewSearchRequest(
		ADconfig.ADbasedn,
		ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false,
		fmt.Sprintf("(&(cn=%v)(objectclass=user))", user.Username),
		UserAttributes,
		nil)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return nil, err
	}
	if len(sr.Entries) >= 1 {
		userentry := sr.Entries[0]
		log.Infof("TestSearch: %s -> num of entries = %d", searchRequest.Filter, len(sr.Entries))

		aduser := ADUSER{} //NewADUser()
		log.Infoln(aduser)
		//sr.PrettyPrint(4)
		for i := range userentry.Attributes {
			//log.Infoln(i)
			//log.Infoln(userentry.Attributes[i].Name)
			//log.Infoln(userentry.Attributes[i].Values)
			//val := ""
			if len(userentry.Attributes[i].Values) >= 2 {
				aduser[userentry.Attributes[i].Name] = userentry.Attributes[i].Values
			} else {
				aduser[userentry.Attributes[i].Name] = userentry.Attributes[i].Values[0]
			}
		}
		ret, err := json.Marshal(aduser)
		log.Infoln(string(ret))
		return aduser, err
	} else {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}
}

//list user
func listUser(l *ldap.Conn) (retusers []ADUSER) {
	setLog()
	// https://cybernetist.com/2020/05/18/getting-started-with-go-ldap/

	l, _ = setupLdap()

	searchRequest := ldap.NewSearchRequest(
		ADconfig.ADbasedn,
		ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false,
		"(&(objectCategory=person)(objectClass=user))",
		UserAttributes,
		nil)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("TestSearch: %s -> num of entries = %d", searchRequest.Filter, len(sr.Entries))
	var adusers = []ADUSER{}
	for _, userentry := range sr.Entries {
		aduser := ADUSER{} //NewADUser()
		//log.Infoln(aduser)
		//sr.PrettyPrint(4)
		for i := range userentry.Attributes {
			//log.Infoln(i)
			//log.Infoln(userentry.Attributes[i].Name)
			//log.Infoln(userentry.Attributes[i].Values)
			//val := ""
			if len(userentry.Attributes[i].Values) >= 2 {
				aduser[userentry.Attributes[i].Name] = userentry.Attributes[i].Values
			} else {
				aduser[userentry.Attributes[i].Name] = userentry.Attributes[i].Values[0]
			}
			//for j, item := range userentry.Attributes[i].Values {
			//	if j == 0 {
			//		val = item
			//	} else {
			//		val = fmt.Sprintf("%v, %v", val, item)
			//	}
			//}
			//aduser[userentry.Attributes[i].Name] = val
		}
		adusers = append(adusers, aduser)
	}
	ret, err := json.Marshal(adusers)
	log.Infoln(string(ret))
	return adusers
}

//user mod
func modUser(l *ldap.Conn, user ADUSER) (user_ ADUSER, err error) {
	if val, ok := user["username"]; !ok || val == "" {
		return user_, errors.New("no user name")
	}
	modReq := ldap.NewModifyRequest(fmt.Sprintf("cn=%v,cn=%v,%v", user["username"], "Users", ADconfig.ADbasedn), []ldap.Control{})

	user_ = NewADUser(user).ToMap()

	for key, val := range user_ {
		switch val.(type) {
		case []string:
			if key != "username" && key != "memberOf" && key != "distinguishedName" {
				log.Infof("%v:%v", key, len(val.([]string)))
				modReq.Replace(key, val.([]string))
			}
		case int:
			i := val.(int)
			modReq.Replace(key, []string{strconv.Itoa(i)})
		case string:
			if key != "username" && val.(string) != "" && key != "distinguishedName" {
				modReq.Replace(key, []string{val.(string)})
			}
		default:
			log.Errorf("%v:%T, %v", key, val, val)
		}
	}
	if err := l.Modify(modReq); err != nil {
		log.Error("error moding user:", modReq, err)
	}
	return user_, err
}

//user password change
func setPassword(l *ldap.Conn, user ADUSER, password string) error {
	setLog()
	client := &http.Client{}
	l = l
	data := url.Values{}
	data.Set("username", user["username"].(string))
	data.Set("password", password)

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("http://%v:9017/api/v1/user", ADconfig.ADserver), strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, err := client.Do(req)
	if err != nil {
		log.Errorln(err)
		return err
	}
	defer func() {
		if err = resp.Body.Close(); err != nil {
			log.Errorf("response cannot close, %v", err)
		}
	}()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(err)
		log.Errorln(respBody)
		return err
	}
	var responseData = make(map[string]interface{})
	err = json.Unmarshal(respBody, &responseData)
	log.Println("response Status : ", resp.Status)
	//log.Println("response Headers : ", resp.Header)
	//log.Println("response Body : ", string(respBody))
	log.Infof("response Body_parsed : %v", responseData)
	log.Infof("response stderr %v: %v",  responseData["stderr"] != "", responseData["stderr"])
	if responseData["stderr"] != "" {

		errstr:=fmt.Sprint(responseData["stderr"])
		err = errors.New(errstr)
		log.Infof("stderr %v",  err)
		return err
	}
	//cmd=string(respBody)
	return nil
}

//delete user
func delUser(l *ldap.Conn, user ADUSER) (err error) {
	log.Debugf("user : %v", user)
	if user != nil && user["username"] == "" {
		return errors.New("no user name")
	}

	log.Debugf("username : %v", user["username"])
	delreq := ldap.NewDelRequest(fmt.Sprintf("cn=%v,cn=%v,%v", user["username"], "Users", ADconfig.ADbasedn), []ldap.Control{})

	if err := l.Del(delreq); err != nil {
		log.Error("error deleting user:", delreq, err)
		return err
	}
	return err
}

func addComputer(l *ldap.Conn, comname string) (err error) {
	//computer add
	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%v,cn=%v,%v", comname, "Computers", ADconfig.ADbasedn), []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "computer", "organizationalPerson", "person", "user"})
	addReq.Attribute("name", []string{comname})
	addReq.Attribute("sAMAccountName", []string{fmt.Sprintf("%v$", comname)})
	addReq.Attribute("dNSHostName", []string{fmt.Sprintf("%v.%v", comname, ADconfig.ADdomain)})
	addReq.Attribute("instanceType", []string{fmt.Sprintf("%d", 0x00000004)})
	if err := l.Add(addReq); err != nil {
		log.Error("error adding computer:", addReq, err)
	}
	return err

}


//add computer to group

func addComputerToGroup(l *ldap.Conn, comname string, groupname string) (err error) {

	modReq := ldap.NewModifyRequest(fmt.Sprintf("cn=%v,cn=Computers,%v", comname , ADconfig.ADbasedn), []ldap.Control{})
	modReq.Replace("distinguishedName", []string{fmt.Sprintf("cn=%v,ou=%v,%v", comname ,groupname, ADconfig.ADbasedn)})

	err = l.Modify(modReq)
	if err != nil {
		log.Errorf("error moding group: %v, %v", modReq, err)
	}
	return err
}

// del computer from group
func delComputerFromGroup(l *ldap.Conn, comname string, groupname string) (err error) {

	modReq := ldap.NewModifyRequest(fmt.Sprintf("cn=%v,ou=%v,%v", comname ,groupname, ADconfig.ADbasedn), []ldap.Control{})
	modReq.Replace("distinguishedName", []string{fmt.Sprintf("cn=%v,cn=Computers,%v", comname , ADconfig.ADbasedn)})

	err = l.Modify(modReq)
	if err != nil {
		log.Errorf("error moding group: %v, %v", modReq, err)
	}
	return err
}