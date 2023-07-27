package router

import (
	"github.com/gin-gonic/gin"
	"valtec/database"
	"valtec/model"
	"valtec/pkg/res"
)

func selectMaps(router *gin.Engine) {
	router.GET("/maps", func(c *gin.Context) {
		maps := make([]model.Map, 0)
		database.Select(&model.Map{}, &maps, "name", "url", "avatar")
		c.JSON(200, res.Ok().Data(maps))
	})
}
