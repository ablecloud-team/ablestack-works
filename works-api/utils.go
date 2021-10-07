package main

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"net/url"
	"path"
	"regexp"
	"runtime"
)
import b64 "encoding/base64"

const (
	WorkspaceString = "workspace"
	InstanceString  = "instance"
	ServiceDaaS     = "ServiceDaaS"
	ServiceWorks    = "ServiceWorks"
	WorkspaceName   = "WorkspaceName"
	AblecloudWorks  = "Ablecloud.Works"
	AgentOK         = "AgentOK"
	AgentCheck      = "AgentCheck"
	Enable          = "Enable"
	Disable         = "Disable"
	UserVm          = "UserVm"
	ALL             = "all"

	VMStart   = "VMStart"
	VMStop    = "VMStop"
	VMsDeploy = "VMsDeploy"
	VMDestroy = "VMDestroy"

	MsgDBConnectError      = "DB connect error"
	MsgDBConnectOK         = "DB connect success"
	MsgDBNoRows            = "sql: no rows in result set"
	BaseErrorCode          = 9000
	SignatureErrorCode     = 9001
	SQLConnectError        = 9100
	SQLQueryError          = 9101
	NotFound404            = "404 Not Found"
	Conflict409            = "409 Conflict"
	Gone410                = "410 Gone"
	OK200                  = "200 OK"
	Created201             = "201 Created"
	Accepted202            = "202 Accepted"
	NotFound               = "Not Found"
	MessageAccountNotFound = "message.account.not.found"
	MessageSignatureError  = "message.Signature.error"
	MessageAgentUpdateOK   = "message.Agent.update.ok"
)

func setLog() {
	startlogger := logrus.New()
	startlogger.SetFormatter(&nested.Formatter{
		HideKeys: false,
		//FieldsOrder: []string{"component", "category"},
		CallerFirst: false,
		CustomCallerFormatter: func(f *runtime.Frame) string {

			pc := make([]uintptr, 10)
			n := runtime.Callers(11, pc)
			if n == 0 {
				return "[]"
			}
			pc = pc[:10]
			frames := runtime.CallersFrames(pc)
			names := make([]string, 10)
			for {
				frame, more := frames.Next()
				names = append(names, frame.Function)
				if !more {
					break
				}
				return fmt.Sprintf("start [%s:%d][%s()]", path.Base(frame.File), frame.Line, frame.Function)
			}
			//startlogger.Infof("- more:%v | %s", more, frame.Function)
			return fmt.Sprintf("start [%s]", names)
		},
	})
	startlogger.SetReportCaller(true)
	startlogger.Infof("Starting ")
}

func RegexpPassword(payload string) bool {
	//passwordConvention := "^(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])(?=.*[*.!@$%^&(){}[]:;<>,.?/~_+-=|\\]).{8,32}$"
	sc := 0
	fmt.Println(payload)
	re1, _ := regexp.Compile("[0-9a-zA-Z!@#$%^&+=]{8,15}")
	if !re1.MatchString(payload) {
		fmt.Println("re1")
		return false
	}
	re2, err := regexp.Compile("[0-9]")
	if err != nil {
		fmt.Println(err)
	}
	if re2.MatchString(payload) {
		fmt.Println("re2")
		sc = sc + 1
	}
	re3, err := regexp.Compile("[a-zA-Z]")
	if err != nil {
		fmt.Println(err)
	}
	if re3.MatchString(payload) {
		fmt.Println("re3")
		sc = sc + 1
	}
	re5, err := regexp.Compile("[!@#$%^&+=]")
	if err != nil {
		fmt.Println(err)
	}
	if re5.MatchString(payload) {
		fmt.Println("re5")
		sc = sc + 1
	}
	return sc == 3
}

func postfixFill(value int) string {
	return fmt.Sprintf("%03d", value)
}

func baseEncoding(payload string) string {

	//return b64.URLEncoding.EncodeToString([]byte(payload))
	return url.QueryEscape(b64.StdEncoding.EncodeToString([]byte(payload)))
}
