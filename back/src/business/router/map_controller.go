package router

import (
	"github.com/gin-gonic/gin"
	"valtec/business/model"
	"valtec/pkg/database"
	"valtec/pkg/json"
	"valtec/pkg/redis"
	"valtec/pkg/res"
)

func selectMaps(router *gin.Engine) {
	router.GET("/maps", func(c *gin.Context) {
		if redis.HandleJsonCache(c, "maps") {
			return
		}

		maps := make([]model.Map, 0)
		database.Select(&model.Map{}, &maps, "name", "url", "avatar")

		mapsJson := json.ToJson(maps)
		redis.AddJsonCache("maps", mapsJson)
		c.JSON(200, res.Ok().Data(mapsJson))
	})
}
