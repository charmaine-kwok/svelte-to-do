package controllers

import (
	"api/initializers"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create a todo item
func PostTodo(c *gin.Context) {
	// Get data off req body
	var body struct {
		Subject string
	}

	c.Bind(&body)

	// Create a todo
	todo := models.Todo{Subject: body.Subject}
	results := initializers.DB.Create(&todo)

	if results.Error != nil {
		// Handle the error
		c.AbortWithStatus(http.StatusBadRequest)
	}

	// Return it
	c.JSON(http.StatusCreated, gin.H{
		"todo": todo.Subject,
	})
}

// Get a list of todos
func GetAllTodos(c *gin.Context) {
	// Get the todos
	var todos []models.Todo
	results := initializers.DB.Order("id asc").Find(&todos)

	if results.Error != nil {
		// Handle the error
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Return the todos
	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

// Delete a todo item
func DeleteTodo(c *gin.Context) {
	// Get data off req body
	id := c.Param("id")

	// Delete from todos
	initializers.DB.Delete(&models.Todo{}, id)

	// Return it
	c.Status(http.StatusOK)
}

// Update a todo item
func UpdateTodo(c *gin.Context) {
	// Get data off req body
	id := c.Param("id")

	var body struct {
		Subject string
	}

	c.Bind(&body)

	// Update todo item
	initializers.DB.Model(&models.Todo{}).Where("id = ?", id).Update("subject", body.Subject)

	// Return it
	c.JSON(http.StatusOK, gin.H{
		"todo": body.Subject,
	})
}
