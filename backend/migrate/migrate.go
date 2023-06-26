package main

import (
	"api/initializers"
	"api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Done{})
	initializers.DB.AutoMigrate(&models.Todo{})
}
