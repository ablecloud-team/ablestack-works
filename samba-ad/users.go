package main

import (
	"errors"
	"fmt"
	auth "github.com/korylprince/go-ad-auth/v3"
	"io/ioutil"
	"os/exec"
)

type USER struct {
	id string `uri:"id" binding:"required"`
	name string `uri:"username" binding:"optional"`
}

func login(conn *auth.Conn, id string, pw string) (logged bool, groups []string, isAdmin bool, err error){
	setLog()

	status, err := Auth(conn, id, pw)
	if err != nil || status == false{
		log.Error(err)
		return false, nil, false, err
	}
	_, _, groups, _ = listGroups(conn, id)

	isAdmin, err = inGroup(conn, id, "Administrators")

	return status, groups, isAdmin, err

}


func setPassword(username string, password string) (stdout string, stderr string, err error) {
	setLog()
	smbPassword:=exec.Command("/usr/local/samba/bin/smbpasswd", "-s", username)
	pwIn, _:=smbPassword.StdinPipe()
	pwOut, _:=smbPassword.StdoutPipe()
	pwErr, _:=smbPassword.StderrPipe()

	smbPassword.Start()

	pwIn.Write([]byte(fmt.Sprintf("%v\n", password)))
	pwIn.Write([]byte(fmt.Sprintf("%v\n", password)))
	pwIn.Close()

	bout, _ :=ioutil.ReadAll(pwOut)
	berr, err :=ioutil.ReadAll(pwErr)
	stdout = string(bout)
	stderr = string(berr)
	log.Infoln(stdout)
	log.Errorln(stderr)
	if stderr != ""{
		err = errors.New(stderr)
	}
	return stdout, stderr, err
}