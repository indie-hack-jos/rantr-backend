package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rantr/router"
)

func setRouter() *gin.Engine {
	// client := dbConfig() // Assuming the function is exported (starts with a capital letter)

	r := gin.Default()

	// ping test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PONG",
			"status":  http.StatusOK,
			"state":   "success",
		})
	})

	v1 := r.Group("/api/v1")
	{
		routes.RegisterRoutes(v1)
	}

	return r
}

func main() {
	r := setRouter()
	r.Run(":8765")
}
