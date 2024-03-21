package routes

import (
	controllers "rantr/controller"
	"rantr/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(router *gin.RouterGroup) {
	router.POST("/post/create", middleware.AuthMiddleware(), controllers.CreatePost)
	router.GET("/post/get", controllers.GetPosts)

	// Register other routes similarly
}
