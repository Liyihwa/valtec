package router

import (
	"github.com/gin-gonic/gin"
	"valtec/business/model"
	"valtec/pkg/database"
	"valtec/pkg/json"
	"valtec/pkg/redis"
	"valtec/pkg/res"
)

func selectHeros(router *gin.Engine) {
	router.GET("/heros", func(c *gin.Context) {
		if redis.HandleJsonCache(c, "heros") {
			return
		}

		var heros []model.Hero
		database.Select(nil, &heros, "id", "avatar", "name", "c", "x", "q", "e")
		herosJson := json.ToJson(heros)
		redis.AddJsonCache("heros", herosJson)
		c.JSON(200, res.Ok().Data(herosJson))
	})
}
