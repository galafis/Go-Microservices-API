package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/galafis/Go-Microservices-API/internal/models"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.APIResponse{
		Status:  "healthy",
		Message: "Service is running properly",
		Author:  "Gabriel Demetrios Lafis",
	})
}

