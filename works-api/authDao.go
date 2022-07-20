package main

import (
	"database/sql"
	"net/http"
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

func insertUserDB(user User) map[string]interface{} {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	resultReturn := map[string]interface{}{}
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
		resultReturn["message"] = "DB connect error"
	}
	defer db.Close()
	result, err := db.Exec("INSERT INTO users(uuid, user_name, password, create_date) VALUES (?, ?, ?, now())", getUuid(), user.Cn, user.Password)
	if err != nil {
		log.Errorf("유저를 DB 등록중에러가 발생했습니다. [%v]", err)
		resultReturn["message"] = "An error occurred while registering Async Job."
		resultReturn["status"] = SQLQueryError
	}
	n, _ := result.RowsAffected()
	if n == 1 {
		log.Info("유저가 정상적으로 등록되었습니다.")
		resultReturn["status"] = http.StatusOK
		resultReturn["message"] = "The async job has been successfully registered."
	}
	return resultReturn
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

func deleteUserDB(userName string) map[string]interface{} {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
	}
	defer db.Close()
	log.Info("DB connect success")
	returnValue := map[string]interface{}{}
	result, err := db.Exec("UPDATE users SET removed=NOW() WHERE user_name=? ", userName)
	if err != nil {
		log.Errorf("유저 삭제중 에러가 발생했습니다.\n")
		log.Errorf("%v\n", err)
		returnValue["status"] = BaseErrorCode
		returnValue["message"] = "An error occurred while registering Async Job."
	}
	n, _ := result.RowsAffected()
	if n == 1 {
		log.Infof("%v\n", result)
		log.Infof("유저가 정상적으로 삭제 되였습니다.\n")
		returnValue["status"] = http.StatusOK
		returnValue["message"] = "async job has been updated normally."
	}

	return returnValue
}

func updateUserPassword(userName string, newPassword string) DatabaseReturnValue {
	db, err := sql.Open(os.Getenv("MysqlType"), os.Getenv("DbInfo"))
	if err != nil {
		log.Error("DB connect error")
		log.Error(err)
	}
	defer db.Close()
	log.Info("DB connect success")
	databaseReturnValue := DatabaseReturnValue{}
	result, err := db.Exec("UPDATE users SET password=? WHERE user_name=? ", newPassword, userName)
	if err != nil {
		log.Errorf("유저 비밀번호 업데이트 중 에러가 발생했습니다.\n")
		log.Errorf("%v\n", err)
		databaseReturnValue.status = BaseErrorCode
		databaseReturnValue.message = "An error occurred while registering Async Job."
	}
	n, _ := result.RowsAffected()
	if n == 1 {
		log.Infof("%v\n", result)
		log.Infof("유저 비밀번호 업데이트가 정상적으로 되였습니다.\n")
		databaseReturnValue.status = http.StatusOK
		databaseReturnValue.message = "async job has been updated normally."
	}

	return databaseReturnValue
}
