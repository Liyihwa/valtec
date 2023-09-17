package redis

import (
	"github.com/gin-gonic/gin"
	"valtec/pkg/json"
	"valtec/pkg/res"
)

func HandleJsonCache(c *gin.Context, key string) bool {
	cache := Get(key)
	if cache != "" {
		var data interface{}
		json.ToJsonObj(cache, &data)
		c.JSON(200, res.Ok().Data(data))
		return true
	}
	return false
}

func AddJsonCache(key string, value any) {
	jsons := json.ToJson(value)
	Set(key, jsons, cacheExpir)
}
