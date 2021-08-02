package main

import (
	auth "github.com/korylprince/go-ad-auth/v3"
	//log "github.com/sirupsen/logrus"
)

//var log = logrus.New().WithField("who", "LOGIN")

func login(conn *auth.Conn, id string, pw string) (logged bool, groups []string, isAdmin bool, err error){
	setLog()
	_, status, err := ConnectAD()
	if err != nil{
		log.Error(err)
		return false, nil, false, err
	}
	status, err = Auth(conn, id, pw)
	if err != nil || status == false{
		log.Error(err)
		return false, nil, false, err
	}
	_, _, groups, _ = listGroups(conn, id)

	isAdmin, err = inGroup(conn, id, "Administrators")

	return status, groups, isAdmin, err

}
