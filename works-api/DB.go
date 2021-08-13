package main

import (
	"database/sql"
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func login1() {
	//var name, email, phone string
	q := getWorksInfo()
	fmt.Println(q.Database.User.ID)
	db, err := sql.Open(os.Getenv("MsqlType"), os.Getenv("DbInfo"))
	checkError(err)
	defer db.Close()

	fmt.Println("connect success")

	count := 0
	//rows, err := db.Query("SELECT user FROM works.users WHERE user='abc' AND password='abc'")
	err = db.QueryRow("SELECT count(user) as count FROM works.users WHERE user='abc' AND password='abc'").Scan(&count)
	checkError(err)
	//defer rows.Close()
	//user := rows.Scan(&user)


	fmt.Println("jkoweojpwerpoeqwropkqweopkqwe")
	fmt.Println(count)

	//for rows.Next() {
	//	checkError(err)
	//	fmt.Println("rows", name, email, phone)
	//}
	return
}
