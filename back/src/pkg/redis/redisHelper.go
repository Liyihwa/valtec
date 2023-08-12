package redis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"valtec/pkg/res"
)

func HandleJsonCache(c *gin.Context, key string) bool {
	json := Get(key)
	fmt.Println(json)
	if json != "" {
		c.JSON(200, res.Ok().Data(json))
		return true
	}
	return false
}

func AddJsonCache(key string, value string) {
	Set(key, value, cacheExpir)
}
