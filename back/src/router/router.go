package router

import (
	"github.com/gin-gonic/gin"
	"valtec/log"
)

func Router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(log.Logger())

	//user
	register(r)
	login(r)

	//map
	selectMaps(r)

	//position
	selectPosition(r)
	selectPositions(r)

	//hero
	selectHeros(r)

	//admin
	addPosition(r)
	addMap(r)
	addHero(r)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "valtec")
	})

	return r
}
