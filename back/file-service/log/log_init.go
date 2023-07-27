package log

import (
	"fmt"
	"github.com/Liyihwa/logger"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
	"time"
	"valtec-fileservice/pkg/config"
)

var logName string

func info(datetime string, level string, message string) string {
	return fmt.Sprintf("%s,%s,%s\n", datetime, level, message)
}

func init() {
	logName = config.GetConfigSevere("logger", "filename")
	f, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic("open file " + logName + " error: " + err.Error())
	}
	level := logger.LevelFromString(config.GetConfigSevere("project", "level"))

	logger.AddConfig(&logger.Config{
		Level:          level,
		UseColor:       false,
		DataTimeFormat: "2006-01-02 15:04:05.000",
		LogMethods:     [4]logger.LogMethod{info, info, info, info},
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

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("msg", &strings.Builder{})
		c.Next()
		duration := time.Since(start)
		val, _ := c.Get("msg")
		sb, _ := val.(*strings.Builder)
		msg := sb.String()

		logger.Info("[GIN]%6.3fs|%s|%s|%-7s|\"%s\" : %s", duration.Seconds(), code(c.Writer.Status()), c.ClientIP(), c.Request.Method, c.Request.URL.Path, msg)
	}
}
