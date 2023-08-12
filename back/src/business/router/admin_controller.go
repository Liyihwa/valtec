package router

import (
	"github.com/gin-gonic/gin"
	model2 "valtec/business/model"
	"valtec/pkg/database"
	"valtec/pkg/res"
)

func addPosition(router *gin.Engine) {
	router.POST("/position", func(c *gin.Context) {
		position := model2.Position{}
		err := c.ShouldBindJSON(&position)
		if err != nil {
			panic("addPosition err: " + err.Error())
		}
		position.Like = 1
		position.Dislike = 1
		database.AddOne(&position)
		c.JSON(200, res.Ok())
	})
}

func addHero(router *gin.Engine) {
	router.POST("/hero", func(c *gin.Context) {
		hero := model2.Hero{}
		err := c.ShouldBindJSON(&hero)
		if err != nil {
			panic("addHero err: " + err.Error())
		}
		println(hero.Avatar)
		database.AddOne(&hero)
		c.JSON(200, res.Ok())
	})
}

func addMap(router *gin.Engine) {
	router.POST("/map", func(c *gin.Context) {
		_map := model2.Map{}
		err := c.ShouldBindJSON(&_map)
		if err != nil {
			panic("addMap err: " + err.Error())
		}

		database.AddOne(&_map)
		c.JSON(200, res.Ok())
	})
}
