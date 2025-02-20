package router

import (
	"example/web-service-gin/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/albums", controller.GetAlbumsController)
	router.GET("/albums/:id", controller.GetAlbumByIDController)
	router.POST("/albums", controller.PostAlbumsController)

	return router
}
