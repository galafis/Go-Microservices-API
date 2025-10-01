package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/galafis/Go-Microservices-API/internal/models"
)

func GetUsers(c *gin.Context) {
	users := []models.User{
		{ID: 1, Name: "Gabriel Lafis", Email: "gabriel@example.com"},
		{ID: 2, Name: "John Doe", Email: "john@example.com"},
	}

	c.JSON(http.StatusOK, models.APIResponse{
		Status:  "success",
		Message: "Users retrieved successfully",
		Data:    users,
		Author:  "Gabriel Demetrios Lafis",
	})
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.APIResponse{
			Status:  "error",
			Message: "Invalid request data",
			Author:  "Gabriel Demetrios Lafis",
		})
		return
	}

	// In a real application, you would save the user to a database
	// For this example, we simulate an auto-increment ID
	user.ID = 3 

	c.JSON(http.StatusCreated, models.APIResponse{
		Status:  "success",
		Message: "User created successfully",
		Data:    user,
		Author:  "Gabriel Demetrios Lafis",
	})
}

