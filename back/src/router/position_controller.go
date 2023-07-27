package router

import (
	"github.com/gin-gonic/gin"
	"valtec/database"
	"valtec/model"
	"valtec/pkg/res"
)

func selectPositions(router *gin.Engine) {
	router.GET("/positions", func(c *gin.Context) {
		where := model.Position{}
		err := c.ShouldBindQuery(&where)
		if err != nil {
			panic("selectPositions err: " + err.Error())
		}
		var positions []model.Position
		database.Select(where, &positions, "stand_x", "stand_y", "put_x", "put_y")
		c.JSON(200, res.Ok().Data(positions))
	})
}

func selectPosition(router *gin.Engine) {
	router.GET("/position", func(c *gin.Context) {
		where := model.Position{}
		err := c.ShouldBindQuery(&where)
		if err != nil {
			panic("selectPosition err: " + err.Error())
		}
		position := model.Position{}
		dbRes := database.SelectOne(where, &position, "description", "like", "dislike")
		if dbRes.RowsAffected == 0 {
			c.JSON(200, res.Fail().Msg("No such id"))
		} else {
			c.JSON(200, res.Ok().Data(position))
		}
	})
}
