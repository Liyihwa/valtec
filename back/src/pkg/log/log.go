package log

import (
	"fmt"
	"github.com/Liyihwa/logwa"
	"os"
	"strings"
	"valtec/pkg/config"
)

var logName string
var Log *logwa.Logger

func log(datetime string, level string, message string) string {
	return fmt.Sprintf("%s|%s|%s\n", datetime, level, message)
}

func init() {
	logName = config.GetStringSevere("logwa", "filename")
	f, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic("open file " + logName + " error: " + err.Error())
	}
	levelString := config.GetStringSevere("logwa", "level")

	var level logwa.LogLevel
	switch strings.ToLower(levelString) {
	case "debug":
		level = logwa.DEBUG
	case "info":
		level = logwa.INFO
	case "warn":
	case "warning":
		level = logwa.WARNING
	case "err":
	case "erro":
	case "error":
		level = logwa.ERROR
	default:
		panic("No such level: " + strings.ToLower(levelString))
	}

	Log = logwa.NewLogger(logwa.Config{
		Level:          level,
		UseColor:       false,
		DataTimeFormat: "2006-01-02 15:04:05.000",
		LogMethods:     [4]logwa.LogMethod{log, log, log, log},
		Target:         f,
	})
}
func code(c int) string {
	if c >= 200 && c < 300 {
		return fmt.Sprintf("{_g}%d{;}", c)
	}
	if c >= 300 && c < 400 {
		return fmt.Sprintf("{_u}%d{;}", c)
	}
	if c >= 400 && c < 500 {
		return fmt.Sprintf("{_y}%d{;}", c)
	}
	if c >= 500 {
		return fmt.Sprintf("{_r}%d{;}", c)
	}
	return fmt.Sprintf("%d", c)
}

func Debug(logFmt string, args ...any) {
	logwa.Debug(logFmt, args)
	Log.Debug(logFmt, args)
}

func Info(logFmt string, args ...any) {
	logwa.Debug(logFmt, args)
	Log.Info(logFmt, args)
}

func Warn(logFmt string, args ...any) {
	logwa.Warn(logFmt, args)
	Log.Warn(logFmt, args)
}

func Erro(logFmt string, args ...any) {
	logwa.Erro(logFmt, args)
	Log.Erro(logFmt, args)
}
