package main

import (
	"log"
	"net/http"

	"rantr/config"
	routes "rantr/router"

	"github.com/gin-gonic/gin"
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
		routes.RegisterUserRoutes(v1)
		routes.RegisterPostRoutes(v1)
	}


	return r
}

func main() {
	if err := config.DbConfig(); err != nil {
        log.Fatalf("Failed to initialize MongoDB: %v", err)
    }

	r := setRouter()
	r.Run(":8765")
}
