package main

import (
	"os"

	"github.com/SanjeevChoubey/go_Jwt_Redis_Mongo/middleware"
	"github.com/SanjeevChoubey/go_Jwt_Redis_Mongo/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	router.GET("/api-1", func(c *gin.Context) {

		c.JSON(200, gin.H{"success": "Access granted for api-1"})

	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}
