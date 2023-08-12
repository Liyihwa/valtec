package router

import (
	"github.com/gin-gonic/gin"
	"valtec/file-service/middleware"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.LogMiddleware())
	router.Use(middleware.CorsMiddleware())

	upload(router)
	deleteFile(router)
	return router
}
