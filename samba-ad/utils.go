package main

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)



func setLog(msg ...string){
	if len(msg)==0{
		msg = append(msg,"Starting")
	}
	startlogger := logrus.New()
	startlogger.SetFormatter(&nested.Formatter{
		HideKeys: false,
		//FieldsOrder: []string{"component", "category"},
		CallerFirst: false,
		CustomCallerFormatter: func(f *runtime.Frame) string {

			File := make([]string, 0)
			Line := make([]int, 0)
			Func := make([]string, 0)
			pc := make([]uintptr, 10)
			n := runtime.Callers(11, pc)
			if n == 0 {
				return "[]"
			}
			pc = pc[:10]
			frames := runtime.CallersFrames(pc)
			names := make([]string, 2)
			for {
				frame, more  := frames.Next()
				names = append(names, frame.Function)
				if !more {
					break
				}
				File=append(File, frame.File)
				Line=append(Line, frame.Line)
				Func=append(Func, frame.Function)
				//fmt.Println(File, Line, Func)
			}
			return fmt.Sprintf(" [%s:%d][%s()] from [%s:%d][%s()] ", path.Base(File[0]), Line[0], Func[0], path.Base(File[1]), Line[1], Func[1])
			//startlogger.Infof("- more:%v | %s", more, frame.Function)
			//return fmt.Sprintf("start [%s]",names)
		},
	})
	startlogger.SetReportCaller(true)
	startlogger.Infof("%v", msg)
}
