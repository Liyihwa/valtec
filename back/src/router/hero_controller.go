package router

import (
	"github.com/gin-gonic/gin"
	"valtec/database"
	"valtec/model"
	"valtec/pkg/res"
)

func selectHeros(router *gin.Engine) {
	router.GET("/heros", func(c *gin.Context) {
		var heros []model.Hero
		database.Select(nil, &heros, "avatar", "name", "c", "x", "q", "e")
		c.JSON(200, res.Ok().Data(heros))
	})
}
