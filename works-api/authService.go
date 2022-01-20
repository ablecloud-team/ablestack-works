package main

import (
	"database/sql"
	"os"
)

type User struct {
	Id         string `json:"id"`
	Uuid       string `json:"uuid"`
	UserName   string `json:"user_name"`
	CreateDate string `json:"create_date"`
	Password   string `json:"password"`
	Removed    string `json:"removed"`
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
