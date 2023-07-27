package router

import (
	"github.com/gin-gonic/gin"
	"valtec-fileservice/log"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(log.Logger())
	router.Use(corsMiddleware())

	upload(router)
	deleteFile(router)
	return router
}

// 跨域中间件
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
