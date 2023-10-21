package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"valtec/pkg/log"
)

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
		c.Set("msg", &strings.Builder{})
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		val, _ := c.Get("msg")
		sb, _ := val.(*strings.Builder)
		msg := sb.String()

		log.Info("[GIN]%6.3f|"+code(c.Writer.Status())+"|%s|%-5s|%s : %s", duration.Seconds(), c.ClientIP(), c.Request.Method, c.Request.URL.Path, msg)
	}
}

//10.16.19.138
