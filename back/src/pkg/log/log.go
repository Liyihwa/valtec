package log

import (
	"github.com/Liyihwa/logwa"
	"os"
	"valtec/pkg/config"
)

var Log *logwa.MultipleLogger

func init() {
	Log = logwa.NewMultipleLogger()

	if fileOn := config.GetBoolSever("logwa", "fileOn"); fileOn {
		level := logwa.LevelFromString(config.GetStringSevere("logwa", "fileLevel"))
		logMethods := logwa.DefaultMethods()
		dataTimeFmt := logwa.DefaultDateTimeFormat
		filePath := config.GetStringSevere("logwa", "filePath")
		f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
		if err != nil {
			panic("open file " + filePath + " error: " + err.Error())
		}
		Log.AddLogger(logwa.NewBaseLogger(
			logwa.Config{
				Level:          level,
				LogMethods:     logMethods,
				DataTimeFormat: dataTimeFmt,
				Target:         f,
				UseColor:       false,
			}))
	}

	if stdOn := config.GetBoolSever("logwa", "stdOn"); stdOn {
		level := logwa.LevelFromString(config.GetStringSevere("logwa", "stdLevel"))
		logMethods := logwa.DefaultMethods()
		dataTimeFmt := logwa.DefaultDateTimeFormat

		Log.AddLogger(logwa.NewBaseLogger(logwa.Config{
			Level:          level,
			UseColor:       true,
			DataTimeFormat: dataTimeFmt,
			LogMethods:     logMethods,
			Target:         os.Stdout,
		}))
	}
}

func Debug(fmtString string, args ...any) {
	Log.Debug(fmtString, args...)
}

func Info(fmtString string, args ...any) {
	Log.Info(fmtString, args...)
}

func Warn(fmtString string, args ...any) {
	Log.Warn(fmtString, args...)
}

func Erro(fmtString string, args ...any) {
	Log.Erro(fmtString, args...)
}
