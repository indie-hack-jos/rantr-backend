package routes

import (
	controllers "rantr/controller"
	"rantr/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.RouterGroup) {
	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.LoginUser)
	router.GET("/auth/user", middleware.AuthMiddleware(), controllers.AuthUser)

	// Register other routes similarly
	// store user routes
}
