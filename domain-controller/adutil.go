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
	"strconv"

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
var UserAttributes = []string{
	"dn",
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
}
//
//var username = "user"
//var password = "Ablecloud1!"

var config = &auth.Config{
	Server:   ADserver,
	Port:     ADport,
	BaseDN:   ADbasedn,
	Security: auth.SecurityInsecureTLS,
}

func ConnectAD() (conn *auth.Conn, status bool, err error) {
	setLog()
	//connect
	conn, err = config.Connect()
	if err != nil {
		//log.Errorln(err)
		return conn, false, err
	}

	//bind
	upn, err := config.UPN(ADusername)
	if err != nil {
		//log.Errorln(err)
		upn = ADusername
	}
	status, err = conn.Bind(upn, ADpassword)
	if err != nil {
		//log.Errorln(err)
		return conn, false, err
	}
	if !status {
		return conn, false, nil
	}
	return conn, status, err
}

func Auth(conn *auth.Conn, username string, password string) (status bool, err error) {
	status, err = auth.Authenticate(config, username, password)
	return status, err
}
func listGroups(conn *auth.Conn, username string) (status bool, entry *ldap.Entry, userGroups []string, err error) {
	setLog()
	upn, err := config.UPN(username)
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
	_, _, groups, err := listGroups(conn, username)
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
	ldapsServer := fmt.Sprintf("ldaps://%v:%v", ADserver, ADport)
	//ldapServer := fmt.Sprintf("ldap://%v:%v", ADserver, ADport)
	l, err = ldap.DialURL(ldapsServer, ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	if err != nil {
		//log.Fatal(err)
	}

	upn, err := config.UPN(ADusername)
	if err != nil {
		//log.Errorln(err)
		upn = ADusername
	}
	err = l.Bind(upn, ADpassword)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	return
}

func getUser(l *ldap.Conn, user *USER) (retuser ADUSER){
	setLog()
	// https://cybernetist.com/2020/05/18/getting-started-with-go-ldap/

	l, _ = setupLdap()

	searchRequest := ldap.NewSearchRequest(
		ADbasedn,
		ldap.ScopeWholeSubtree, ldap.DerefAlways, 0, 0, false,
		fmt.Sprintf("(cn=%v)", user.Username),
		UserAttributes,
		nil)

	sr, err := l.Search(searchRequest)
	if err != nil {
		//log.Fatal(err)
	}
	userentry := sr.Entries[0]
	log.Infof("TestSearch: %s -> num of entries = %d", searchRequest.Filter, len(sr.Entries))

	aduser := ADUSER{}//NewADUser()
	log.Infoln(aduser)
	//sr.PrettyPrint(4)
	for i := range userentry.Attributes {
		log.Infoln(i)
		log.Infoln(userentry.Attributes[i].Name)
		log.Infoln(userentry.Attributes[i].Values)
		//val := ""
		if len(userentry.Attributes[i].Values) >=2{
			aduser[userentry.Attributes[i].Name] = userentry.Attributes[i].Values
		} else{
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
	ret, err :=json.Marshal(aduser)
	log.Infoln(string(ret))
	return aduser
}

func searchUser(l *ldap.Conn) (retusers []ADUSER){
	setLog()
	// https://cybernetist.com/2020/05/18/getting-started-with-go-ldap/

	l, _ = setupLdap()

	searchRequest := ldap.NewSearchRequest(
		ADbasedn,
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
	for _,userentry := range sr.Entries {
		aduser := ADUSER{} //NewADUser()
		log.Infoln(aduser)
		//sr.PrettyPrint(4)
		for i := range userentry.Attributes {
			log.Infoln(i)
			log.Infoln(userentry.Attributes[i].Name)
			log.Infoln(userentry.Attributes[i].Values)
			//val := ""
			if len(userentry.Attributes[i].Values) >=2{
				aduser[userentry.Attributes[i].Name] = userentry.Attributes[i].Values
			} else{
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
	ret, err :=json.Marshal(adusers)
	log.Infoln(string(ret))
	return adusers
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
		ADbasedn,
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
func addGroup(l *ldap.Conn, groupname string) (err error) {
	//ou add
	addReq := ldap.NewAddRequest(fmt.Sprintf("ou=%v,%v", groupname, ADbasedn), []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "organizationalUnit"})
	addReq.Attribute("name", []string{groupname})
	addReq.Attribute("ou", []string{groupname})
	addReq.Attribute("instanceType", []string{fmt.Sprintf("%d", 0x00000004)})
	if err := l.Add(addReq); err != nil {
		log.Error("error adding OU:", addReq, err)
	}

	//group add
	addReq = ldap.NewAddRequest(fmt.Sprintf("cn=%v,ou=%v,%v", groupname, groupname, ADbasedn), []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "group"})
	addReq.Attribute("name", []string{groupname})
	addReq.Attribute("sAMAccountName", []string{groupname})
	addReq.Attribute("instanceType", []string{fmt.Sprintf("%d", 0x00000004)})
	addReq.Attribute("groupType", []string{fmt.Sprintf("%d", 0x00000004|0x80000000)})
	if err := l.Add(addReq); err != nil {
		log.Error("error adding group:", addReq, err)
	}
	return err
}
func addComputer(l *ldap.Conn, comname string) (err error) {
	//computer add
	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%v,cn=%v,%v", comname, "Computers", ADbasedn), []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "computer", "organizationalPerson", "person", "user"})
	addReq.Attribute("name", []string{comname})
	addReq.Attribute("sAMAccountName", []string{fmt.Sprintf("%v$", comname)})
	addReq.Attribute("dNSHostName", []string{fmt.Sprintf("%v.%v", comname, domain)})
	addReq.Attribute("instanceType", []string{fmt.Sprintf("%d", 0x00000004)})
	if err := l.Add(addReq); err != nil {
		log.Error("error adding computer:", addReq, err)
	}
	return err

}
type ADUSER map[string] interface{}
/*
type ADUser struct {
	sn                         string //:="새"
	givenName                  string //:="사용자"
	initials                   string //:="이니셜"
	accountname                string //= "newuser"
	userPrincipalName          string //:=fmt.Sprintf("%v@%v",accountname, domain) //로그온 이름(accountname@domain 형식)
	username                   string `uri:"username" binding:"required"` //:=fmt.Sprintf("%v %v %v.", sn, givenName, initials)
	sAMAccountName             string //:=accountname //windows 2000 이전 사용자 로그온 이름(domain\sAMAccountName 형식)
	description                string //:="설명"//설명
	info                       string //:="참고"//참고내용
	title                      string //:="책임"//직함
	o                          string //:=company // ldap 회사명
	company                    string //:="에이블"//회사명
	postOfficeBox              string //:="사서함"//사서함 주소
	physicalDeliveryOfficeName string //:="사무실"//사무실
	streetAddress              string //:="만년동"//주소
	I                          string //:="서구"//구/군/시
	st                         string //:="대전"//시/도
	department                 string //:="개발팀"
	mail                       string //:="ycyun@ablecloud.io"//메일주소
	telephoneNumber            string //:="010"//일반->전화
	pager                      string //:="011"//전화 -> 호출기
	mobile                     string //:="016"//전화->휴대폰
	facsimileTelephoneNumber   string //:="팩스"//전화->팩스
	homePhone                  string //:="042"//전화->집
	ipPhone                    string //:="070"//전화->ip전화
	postalCode                 string //:="35200"//주소->우편번호
	manager                    string //:="CN=User3,CN=Users,DC=dc1,DC=local"//AD 상사이름
	wWWHomePage                string //:="https://www."//홈페이지 주소
	memberOf					[]string
	countryCode int    //:= 410//주소->국가(한국?)
	c           string //:="KR"//국가명
	co          string //:=c
}
*/

//user add
func addUser(l *ldap.Conn, user ADUSER) (err error) {
	if val, ok := user["username"]; !ok || val=="" {
		return errors.New("no user name")
	}

	user["userPrincipalName"] = fmt.Sprintf("%v@%v", user["username"], domain) //로그온 이름(accountname@domain 형식)


	user["sAMAccountName"] = user["username"] //windows 2000 이전 사용자 로그온 이름(domain\sAMAccountName 형식)

	user["countryCode"] = 410 //주소->국가(한국?)
	user["c"] = "KR"          //국가명
	user["co"] = user["c"]

	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%v,cn=%v,%v", user["username"], "Users", ADbasedn), []ldap.Control{})

	for key, val := range user{
		log.Infof("%v:%v", key, val)
		switch val.(type){
		case string:
			if key != "username" {
				addReq.Attribute(key, []string{val.(string)})
			}
		case int:
			i := val.(int)
			addReq.Attribute(key, []string{strconv.Itoa(i)})
		default:
			log.Errorf("%v:%T, %v", key, val, val)
		}
	}
	addReq.Attribute("objectClass", []string{"top", "organizationalPerson", "person", "user"})
	//addReq.Attribute("name", []string{user.username})
	//addReq.Attribute("sn", []string{user.sn})
	//addReq.Attribute("givenName", []string{user.givenName})
	//addReq.Attribute("initials", []string{user.initials})
	//addReq.Attribute("displayName", []string{user.username})
	//addReq.Attribute("cn", []string{user.username})
	//addReq.Attribute("userPrincipalName", []string{user.userPrincipalName})
	//addReq.Attribute("instanceType", []string{fmt.Sprintf("%d", 0x00000004)})
	//addReq.Attribute("countryCode", []string{fmt.Sprintf("%d", user.countryCode)})
	//addReq.Attribute("c", []string{user.c})
	//addReq.Attribute("co", []string{user.co})
	if err := l.Add(addReq); err != nil {
		log.Error("error adding user:", addReq, err)
	}
	//ldap.NewDelRequest(fmt.Sprintf("cn=%v,cn=%v,%v", user["username"], "Users", ADbasedn), []ldap.Control{})
	return err
}


//user add
func modUser(l *ldap.Conn, user ADUSER) (err error) {
	if val, ok := user["username"]; !ok || val=="" {
		return errors.New("no user name")
	}

	user["userPrincipalName"] = fmt.Sprintf("%v@%v", user["username"], domain) //로그온 이름(username@domain 형식)


	user["sAMAccountName"] = user["username"] //windows 2000 이전 사용자 로그온 이름(domain\sAMAccountName 형식)

	user["countryCode"] = 410 //주소->국가(한국?)
	user["c"] = "KR"          //국가명
	user["co"] = user["c"]

	modReq := ldap.NewModifyRequest(fmt.Sprintf("cn=%v,cn=%v,%v", user["username"], "Users", ADbasedn), []ldap.Control{})
	modReq.Add("description", []string{"An example user"})
	modReq.Replace("mail", []string{"user@example.org"})

	err = l.Modify(modReq)
	if err != nil {
		log.Fatal(err)
	}

	for key, val := range user{
		log.Infof("%v:%v", key, val)
		switch val.(type){
		case string:
			if key != "username" {
				modReq.Replace(key, []string{val.(string)})
			}
		case int:
			i := val.(int)
			modReq.Replace(key, []string{strconv.Itoa(i)})
		default:
			log.Errorf("%v:%T, %v", key, val, val)
		}
	}
	modReq.Replace("objectClass", []string{"top", "organizationalPerson", "person", "user"})
	//modReq.Attribute("name", []string{user.username})
	//modReq.Attribute("sn", []string{user.sn})
	//modReq.Attribute("givenName", []string{user.givenName})
	//modReq.Attribute("initials", []string{user.initials})
	//modReq.Attribute("displayName", []string{user.username})
	//modReq.Attribute("cn", []string{user.username})
	//modReq.Attribute("sAMAccountName", []string{user.sAMAccountName})
	//modReq.Attribute("userPrincipalName", []string{user.userPrincipalName})
	//modReq.Attribute("instanceType", []string{fmt.Sprintf("%d", 0x00000004)})
	//modReq.Attribute("countryCode", []string{fmt.Sprintf("%d", user.countryCode)})
	//modReq.Attribute("c", []string{user.c})
	//modReq.Attribute("co", []string{user.co})
	if err := l.Modify(modReq); err != nil {
		log.Error("error adding user:", modReq, err)
	}
	//ldap.NewDelRequest(fmt.Sprintf("cn=%v,cn=%v,%v", user["username"], "Users", ADbasedn), []ldap.Control{})
	return err
}
//user delete
func delUser(l *ldap.Conn, user ADUSER) (err error) {
	log.Debugf("user : %v",user)
	if user !=nil && user["username"] == "" {
		return errors.New("no user name")
	}

	log.Debugf("username : %v",user["username"])
	delreq := ldap.NewDelRequest(fmt.Sprintf("cn=%v,cn=%v,%v", user["username"], "Users", ADbasedn), []ldap.Control{})

	if err := l.Del(delreq); err != nil {
		log.Error("error deleting user:", delreq, err)
		return err
	}
	return err
}

func setPassword(l *ldap.Conn, user ADUSER, password string) (err error) {
	setLog()
	client := &http.Client{}

	data := url.Values{}
	data.Set("username", user["username"].(string))
	data.Set("password", password)

	req, err := http.NewRequest(http.MethodPatch, "http://10.1.1.8:8082/api/v1/user", strings.NewReader(data.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, err := client.Do(req)
	if err != nil {
		log.Errorln(err)
		return
	}
	defer resp.Body.Close()

	// Response 체크.
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(err)
		log.Errorln(respBody)
		return
	}
	var responseData = make(map[string]interface{})
	err = json.Unmarshal(respBody, &responseData)
	log.Println("response Status : ", resp.Status)
	log.Println("response Headers : ", resp.Header)
	log.Println("response Body : ", string(respBody))
	log.Println("response Body_parsed : ", responseData)
	if responseData["stderr"] != "" {
		return errors.New(fmt.Sprint(responseData["stderr"]))
	}
	//cmd=string(respBody)
	return nil
}
