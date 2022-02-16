package main

import (
	"database/sql"
	"os"
)

type User struct {
	Id                string   `json:"id"`
	Uuid              string   `json:"uuid"`
	UserName          string   `json:"user_name"`
	CreateDate        string   `json:"create_date"`
	Password          string   `json:"password"`
	Removed           string   `json:"removed"`
	C                 string   `json:"c"`                 // 국가코드 영문
	Cn                string   `json:"cn"`                // 유저아이디
	Co                string   `json:"co"`                // 국가코드 영문
	CountryCode       string   `json:"countryCode"`       // 국가코드 숫자
	DistinguishedName string   `json:"distinguishedName"` // 고유이름
	GivenName         string   `json:"givenName"`         // 사용자 이름
	Mail              string   `json:"mail"`              // e-mail
	MemberOf          string   `json:"memberOf"`          // 소속 그룹
	Name              string   `json:"name"`              // 사용자 계정
	SAMAccountName    string   `json:"sAMAccountName"`    // 사용자 계정
	Sn                string   `json:"sn"`                // 사용자 성
	TelephoneNumber   string   `json:"telephoneNumber"`   // 전화번호
	Title             string   `json:"title"`             // 직책
	UserPrincipalName string   `json:"userPrincipalName"` // 사용자 도메인 계정정보
	Department        string   `json:"department"`        // 사용자 부서
	ClusterName       string   `json:"clusterName"`
	DomainName        string   `json:"domainName"`
	Login             bool     `json:"login"`
	UserNameDc        string   `json:"username"`
	Groups            []string `json:"groups"`
	IsAdmin           bool     `json:"isAdmin"`
	Token             string   `json:"token"`
}

func selectUserDBDetail(userName string) User {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
	}
	defer db.Close()
	log.Info("DB connect success")
	user := User{}
	err = db.QueryRow("SELECT uuid, user_name, create_date, password FROM users WHERE removed IS NULL AND user_name = ? ORDER BY id DESC ", userName).Scan(
		&user.Uuid, &user.UserName, &user.CreateDate, &user.Password)
	if err != nil {
		log.Error("user select query FAILED")
		log.Error(err.Error())
	}

	return user
}
