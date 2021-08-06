package main

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

func setLog(){
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
				frame, more  := frames.Next()
				names = append(names, frame.Function)
				if !more {
					break
				}
				return fmt.Sprintf("start [%s:%d][%s()]", path.Base(frame.File), frame.Line, frame.Function)
			}
			//startlogger.Infof("- more:%v | %s", more, frame.Function)
			return fmt.Sprintf("start [%s]",names)
		},
	})
	startlogger.SetReportCaller(true)
	startlogger.Infof("Starting ")
}