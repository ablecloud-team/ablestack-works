package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	auth "github.com/korylprince/go-ad-auth/v3"
	"reflect"

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
		"description"}

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

type ADUser struct {
	sn                         string //:="새"
	givenName                  string //:="사용자"
	initials                   string //:="이니셜"
	accountname                string //= "newuser"
	userPrincipalName          string //:=fmt.Sprintf("%v@%v",accountname, domain) //로그온 이름(accountname@domain 형식)
	username                   string //:=fmt.Sprintf("%v %v %v.", sn, givenName, initials)
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

	countryCode int    //:= 410//주소->국가(한국?)
	c           string //:="KR"//국가명
	co          string //:=c
}

func NewADUser() (user *ADUser) {
	user = &ADUser{}
	user.sn = ""
	user.givenName = ""                  //"사용자"
	user.initials = ""                   //"이니셜"
	user.accountname = ""                //= "newuser"
	user.userPrincipalName = ""          //fmt.Sprintf("%v@%v",accountname, domain) //로그온 이름(accountname@domain 형식)
	user.username = ""                   //fmt.Sprintf("%v %v %v.", sn, givenName, initials)
	user.sAMAccountName = ""             //accountname //windows 2000 이전 사용자 로그온 이름(domain\sAMAccountName 형식)
	user.description = ""                //"설명"//설명
	user.info = ""                       //"참고"//참고내용
	user.title = ""                      //"책임"//직함
	user.o = ""                          //company // ldap 회사명
	user.company = ""                    //"에이블"//회사명
	user.postOfficeBox = ""              //"사서함"//사서함 주소
	user.physicalDeliveryOfficeName = "" //"사무실"//사무실
	user.streetAddress = ""              //"만년동"//주소
	user.I = ""                          //"서구"//구/군/시
	user.st = ""                         //"대전"//시/도
	user.department = ""                 //"개발팀"
	user.mail = ""                       //"ycyun@ablecloud.io"//메일주소
	user.telephoneNumber = ""            //"010"//일반->전화
	user.pager = ""                      //"011"//전화 -> 호출기
	user.mobile = ""                     //"016"//전화->휴대폰
	user.facsimileTelephoneNumber = ""   //"팩스"//전화->팩스
	user.homePhone = ""                  //"042"//전화->집
	user.ipPhone = ""                    //"070"//전화->ip전화
	user.postalCode = ""                 //"35200"//주소->우편번호
	user.manager = ""                    //"CN=User3,CN=Users,DC=dc1,DC=local"//AD 상사이름
	user.wWWHomePage = ""                //"https://www."//홈페이지 주소
	user.countryCode = 410               //:= 410//주소->국가(한국?)
	user.c = ""                          //"KR"//국가명
	user.co = user.c
	return user
}

//user add
func addUser(l *ldap.Conn, user *ADUser) (err error) {
	if user.accountname == "" {
		return errors.New("No user name")
	}

	user.userPrincipalName = fmt.Sprintf("%v@%v", user.accountname, domain) //로그온 이름(accountname@domain 형식)
	user.username = fmt.Sprintf("%v %v %v.", user.sn, user.givenName, user.initials)
	user.sAMAccountName = user.accountname //windows 2000 이전 사용자 로그온 이름(domain\sAMAccountName 형식)


	//description:="설명"//설명
	//info:="참고"//참고내용
	//title:="책임"//직함
	//o:=company // ldap 회사명
	//company:="에이블"//회사명
	//postOfficeBox:="사서함"//사서함 주소
	//physicalDeliveryOfficeName:="사무실"//사무실
	//streetAddress:="만년동"//주소
	//I:="서구"//구/군/시
	//st:="대전"//시/도
	//department:="개발팀"
	//mail:="ycyun@ablecloud.io"//메일주소
	//telephoneNumber:="010"//일반->전화
	//pager:="011"//전화 -> 호출기
	//mobile:="016"//전화->휴대폰
	//facsimileTelephoneNumber:="팩스"//전화->팩스
	//homePhone:="042"//전화->집
	//ipPhone:="070"//전화->ip전화
	//postalCode:="35200"//주소->우편번호
	//manager:="CN=User3,CN=Users,DC=dc1,DC=local"//AD 상사이름
	//wWWHomePage:="https://www."//홈페이지 주소

	user.countryCode = 410 //주소->국가(한국?)
	user.c = "KR"          //국가명
	user.co = user.c

	addReq := ldap.NewAddRequest(fmt.Sprintf("cn=%v,cn=%v,%v", user.username, "Users", ADbasedn), []ldap.Control{})

	target := reflect.ValueOf(user)
	elements := target.Elem()
	numFields := target.Elem().NumField()

	//fmt.Printf("Type: %s\n", target.Type()) // 구조체 타입명
	for i := 0; i < numFields; i++ {
		mValue := target.Elem().Field(i)
		mType := elements.Type().Field(i)
		//fmt.Printf("%10s %10s ==> %10v\n",
		//	mType.Name, // 이름
		//	mType.Type, // 타입
		//	mValue,     // 값
		//)
		if mType.Type.String()=="string" && (mValue.String()!="" && mType.Name!="accountname" && mType.Name!="username")  {
			addReq.Attribute(mType.Name, []string{mValue.String()})
		}
	}
	//
	addReq.Attribute("objectClass", []string{"top", "organizationalPerson", "person", "user"})
	if err := l.Add(addReq); err != nil {
		log.Error("error adding user:", addReq, err)
	}
	return err
}

