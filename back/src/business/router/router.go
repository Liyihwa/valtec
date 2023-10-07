package router

import (
	"github.com/gin-gonic/gin"
	"valtec/business/middleware"
)

func Router() *gin.Engine {

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LogMiddleware())
	r.Use(middleware.CorsMiddleware())

	//user
	register(r)
	login(r)
	emailVerif(r)

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

	return r
}
