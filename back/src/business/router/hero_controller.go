package router

import (
	"github.com/gin-gonic/gin"
	"valtec/business/model"
	"valtec/pkg/database"
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

		var ans []model.ResObject
		for _, h := range heros {
			ans = append(ans, h.GetInfo())
		}

		redis.AddJsonCache("heros", ans)
		c.JSON(200, res.Ok().Data(ans))
	})
}
