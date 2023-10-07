package middleware

import (
	"fmt"
	"github.com/Liyihwa/logwa"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"time"
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

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("msg", &strings.Builder{})
		c.Next()
		duration := time.Since(start)
		val, _ := c.Get("msg")
		sb, _ := val.(*strings.Builder)
		msg := sb.String()
		logwa.Info("[GIN]%6.3fs|%s|%s|%-7s|\"%s\" : %s", duration.Seconds(), code(c.Writer.Status()), c.ClientIP(), c.Request.Method, c.Request.URL.Path, msg)
		Log.Info("[GIN]%6.3fs|%d|%s|%-7s|\"%s\" : %s", duration.Seconds(), c.Writer.Status(), c.ClientIP(), c.Request.Method, c.Request.URL.Path, msg)
	}
}

//10.16.19.138
