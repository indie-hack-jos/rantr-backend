package users

import (
	"github.com/gin-gonic/gin"
	controllers "rantr/controller"
	"rantr/middleware"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/login", controllers.LoginUser)
	router.GET("/auth/user", middleware.AuthMiddleware(), controllers.AuthUser)

	// Register other routes similarly
}
