package main

import "strconv"

type ADUSER map[string]interface{}

type ADUser struct {
	sn                         string   `uri:"sn" form:"sn"`                                                 //성
	givenName                  string   `uri:"givenName" form:"givenName"`                                   //이름
	initials                   string   `uri:"initials" form:"initials"`                                     //이니셜
	accountname                string   `uri:"accountname" form:"accountname"`                               //어카운트명
	userPrincipalName          string   `uri:"userPrincipalName" form:"userPrincipalName"`                   //로그온 이름(accountname@domain 형식)
	username                   string   `uri:"username" form:"username"`                  //works 로그인명
	sAMAccountName             string   `uri:"sAMAccountName" form:"sAMAccountName"`                         //windows 2000 이전 사용자 로그온 이름(domain\sAMAccountName 형식)
	description                string   `uri:"description" form:"description"`                               //설명
	info                       string   `uri:"info" form:"info"`                                             //참고내용
	title                      string   `uri:"title" form:"title"`                                           //직함
	o                          string   `uri:"o" form:"o"`                                                    //ldap 회사명
	company                    string   `uri:"company" form:"company"`                                       //AD회사명
	postOfficeBox              string   `uri:"postOfficeBox" form:"postOfficeBox"`                           //사서함 주소
	physicalDeliveryOfficeName string   `uri:"physicalDeliveryOfficeName" form:"physicalDeliveryOfficeName"` //사무실주소
	streetAddress              string   `uri:"streetAddress" form:"streetAddress"`                           //주소
	I                          string   `uri:"I" form:"I"`                                                   //구/군/시
	st                         string   `uri:"st" form:"st"`                                                 //시/도
	department                 string   `uri:"department" form:"department"`                                 //부서
	mail                       string   `uri:"mail" form:"mail"`                                             //메일주소
	telephoneNumber            string   `uri:"telephoneNumber" form:"telephoneNumber"`                       //일반->전화
	pager                      string   `uri:"pager" form:"pager"`                                           //전화 -> 호출기
	mobile                     string   `uri:"mobile" form:"mobile"`                                         //전화->휴대폰
	facsimileTelephoneNumber   string   `uri:"facsimileTelephoneNumber" form:"facsimileTelephoneNumber"`     //전화->팩스
	homePhone                  string   `uri:"homePhone" form:"homePhone"`                                   //전화->집
	ipPhone                    string   `uri:"ipPhone" form:"ipPhone"`                                       //전화->ip전화
	postalCode                 string   `uri:"postalCode" form:"postalCode"`                                 //주소->우편번호
	manager                    string   `uri:"manager" form:"manager"`                                       //상사dn("CN=User3,CN=Users,DC=dc1,DC=local")
	wWWHomePage                string   `uri:"wWWHomePage" form:"wWWHomePage"`                               //홈페이지 주소
	memberOf                   []string `uri:"memberOf" form:"memberOf"`                                     //그룹 dn 목록
	countryCode                int      `uri:"countryCode" form:"countryCode"`                               //주소->국가숫자코드(한국:410)
	c                          string   `uri:"c" form:"c"`                                                   //주소->국가영문코드명(한국:KR)
	co                         string   `uri:"co" form:"co"`                                                 //:=c
}

func NewADUser(aduser ADUSER) (aduserstruct *ADUser) {
	aduserstruct = &ADUser{}

	if val, ok := aduser["sn"]; ok {
		aduserstruct.sn = val.([]string)[0]
	}
	if val, ok := aduser["givenName"]; ok {
		aduserstruct.givenName = val.([]string)[0]
	}
	if val, ok := aduser["initials"]; ok {
		aduserstruct.initials = val.([]string)[0]
	}
	if val, ok := aduser["accountname"]; ok {
		aduserstruct.accountname = val.([]string)[0]
	}
	if val, ok := aduser["userPrincipalName"]; ok {
		aduserstruct.userPrincipalName = val.([]string)[0]
	}
	if val, ok := aduser["sAMAccountName"]; ok {
		aduserstruct.sAMAccountName = val.([]string)[0]
	}
	if val, ok := aduser["description"]; ok {
		aduserstruct.description = val.([]string)[0]
	}
	if val, ok := aduser["info"]; ok {
		aduserstruct.info = val.([]string)[0]
	}
	if val, ok := aduser["title"]; ok {
		aduserstruct.title = val.([]string)[0]
	}
	if val, ok := aduser["o"]; ok {
		aduserstruct.o = val.([]string)[0]
	}
	if val, ok := aduser["company"]; ok {
		aduserstruct.company = val.([]string)[0]
	}
	if val, ok := aduser["postOfficeBox"]; ok {
		aduserstruct.postOfficeBox = val.([]string)[0]
	}
	if val, ok := aduser["physicalDeliveryOfficeName"]; ok {
		aduserstruct.physicalDeliveryOfficeName = val.([]string)[0]
	}
	if val, ok := aduser["streetAddress"]; ok {
		aduserstruct.streetAddress = val.([]string)[0]
	}
	if val, ok := aduser["I"]; ok {
		aduserstruct.I = val.([]string)[0]
	}
	if val, ok := aduser["st"]; ok {
		aduserstruct.st = val.([]string)[0]
	}
	if val, ok := aduser["department"]; ok {
		aduserstruct.department = val.([]string)[0]
	}
	if val, ok := aduser["mail"]; ok {
		aduserstruct.mail = val.([]string)[0]
	}
	if val, ok := aduser["telephoneNumber"]; ok {
		aduserstruct.telephoneNumber = val.([]string)[0]
	}
	if val, ok := aduser["pager"]; ok {
		aduserstruct.pager =val.([]string)[0]
	}
	if val, ok := aduser["mobile"]; ok {
		aduserstruct.mobile = val.([]string)[0]
	}
	if val, ok := aduser["facsimileTelephoneNumber"]; ok {
		aduserstruct.facsimileTelephoneNumber = val.([]string)[0]
	}
	if val, ok := aduser["homePhone"]; ok {
		aduserstruct.homePhone = val.([]string)[0]
	}
	if val, ok := aduser["ipPhone"]; ok {
		aduserstruct.ipPhone = val.([]string)[0]
	}
	if val, ok := aduser["postalCode"]; ok {
		aduserstruct.postalCode = val.([]string)[0]
	}
	if val, ok := aduser["manager"]; ok {
		aduserstruct.manager = val.([]string)[0]
	}
	if val, ok := aduser["wWWHomePage"]; ok {
		aduserstruct.wWWHomePage = val.([]string)[0]
	}
	if val, ok := aduser["username"]; ok {
		aduserstruct.username = val.(string)
	}
	if val, ok := aduser["memberOf"]; ok {
		aduserstruct.memberOf = val.([]string)
	}
	if val, ok := aduser["countryCode"]; ok {
		aduserstruct.countryCode,_ = strconv.Atoi(val.([]string)[0])

	}
	if val, ok := aduser["c"]; ok {
		aduserstruct.c = val.([]string)[0]
	}
	if val, ok := aduser["co"]; ok {
		aduserstruct.co = val.([]string)[0]
	}
	return aduserstruct
}

func (aduserstruct *ADUser) ToMap()(aduser ADUSER){
	aduser = ADUSER{
		"sn": aduserstruct.sn,
		"givenName": aduserstruct.givenName,
		"initials": aduserstruct.initials,
		"accountname": aduserstruct.accountname,
		"userPrincipalName": aduserstruct.userPrincipalName,
		"sAMAccountName": aduserstruct.sAMAccountName,
		"description": aduserstruct.description,
		"info": aduserstruct.info,
		"title": aduserstruct.title,
		"o": aduserstruct.o,
		"company": aduserstruct.company,
		"postOfficeBox": aduserstruct.postOfficeBox,
		"physicalDeliveryOfficeName": aduserstruct.physicalDeliveryOfficeName,
		"streetAddress": aduserstruct.streetAddress,
		"I": aduserstruct.I,
		"st": aduserstruct.st,
		"department": aduserstruct.department,
		"mail": aduserstruct.mail,
		"telephoneNumber": aduserstruct.telephoneNumber,
		"pager": aduserstruct.pager,
		"mobile": aduserstruct.mobile,
		"facsimileTelephoneNumber": aduserstruct.facsimileTelephoneNumber,
		"homePhone": aduserstruct.homePhone,
		"ipPhone": aduserstruct.ipPhone,
		"postalCode": aduserstruct.postalCode,
		"manager": aduserstruct.manager,
		"wWWHomePage": aduserstruct.wWWHomePage,
		"memberOf": aduserstruct.memberOf,
		"countryCode": aduserstruct.countryCode,
		"c": aduserstruct.c,
		"co": aduserstruct.co,
		"username": aduserstruct.username,
	}
	return aduser
}