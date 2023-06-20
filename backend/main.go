package main

import (
	"api/initializers"
	"api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	// Create a new Gin router
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Authorization"}
	// enable credentials
	config.AllowCredentials = true

	r.Use(cors.New(config))

	todoGroup := r.Group("/todo")
	{
		// Register the todo routes
		routes.SetupTodoRoutes(todoGroup)
	}

	doneGroup := r.Group("/done")
	{
		// Register the done routes
		routes.SetupDoneRoutes(doneGroup)
	}

	// Start the server
	r.Run(":8080")
}
