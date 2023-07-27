package router

import (
	"github.com/gin-gonic/gin"
	"valtec/log"
)

func Router() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(log.Logger())
	r.Use(corsMiddleware())

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

	return r
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, FETCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
