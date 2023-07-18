package controllers

import (
	"api/initializers"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create a done item
func PostDone(c *gin.Context) {
	// Get data off req body
	var body struct {
		Subject string
	}

	c.Bind(&body)

	// Create a done item
	done := models.Done{Subject: body.Subject}
	results := initializers.DB.Create(&done)

	if results.Error != nil {
		// Handle the error
		c.AbortWithStatus(http.StatusBadRequest)
	}

	// Return it
	c.JSON(http.StatusCreated, gin.H{
		"done": done.Subject,
	})
}

// Get a list of dones
func GetAllDones(c *gin.Context) {

	// Get the todos
	var dones []models.Done
	results := initializers.DB.Order("id asc").Find(&dones)

	if results.Error != nil {
		// Handle the error
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	// Return it
	c.JSON(http.StatusOK, gin.H{
		"dones": dones,
	})
}

// Delete a done item
func DeleteDone(c *gin.Context) {
	// Get data off req body
	id := c.Param("id")

	// Delete from todos
	initializers.DB.Delete(&models.Done{}, id)

	// Return it
	c.Status(http.StatusOK)
}
