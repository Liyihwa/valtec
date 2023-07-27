package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"valtec/pkg/res"
)

func ShouldEqual(c *gin.Context) {

}

func ShouldNotNull(c *gin.Context, val string, message string) {
	if val == "" {
		c.JSON(200, res.Fail().Msg(message))
	}
}

func ShouldLongerThan(c *gin.Context, val string, length int, prefix string) {
	if len(val) <= length {
		c.JSON(200, res.Fail().Msg(prefix+fmt.Sprintf("长度不得小于%d位", length)))
	}
}
