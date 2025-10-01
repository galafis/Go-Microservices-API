package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/galafis/Go-Microservices-API/internal/models"
)

func TestGetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/users", GetUsers)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/users", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response models.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)
	assert.Equal(t, "Users retrieved successfully", response.Message)
	assert.Equal(t, "Gabriel Demetrios Lafis", response.Author)

	users, ok := response.Data.([]interface{})
	assert.True(t, ok)
	assert.Len(t, users, 2)
}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/users", CreateUser)

	newUser := models.User{Name: "Jane Doe", Email: "jane@example.com"}
	jsonValue, _ := json.Marshal(newUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success", response.Status)
	assert.Equal(t, "User created successfully", response.Message)
	assert.Equal(t, "Gabriel Demetrios Lafis", response.Author)

	createdUser, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, float64(3), createdUser["id"])
	assert.Equal(t, "Jane Doe", createdUser["name"])
	assert.Equal(t, "jane@example.com", createdUser["email"])
}

func TestCreateUserInvalidInput(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/users", CreateUser)

	// Invalid JSON: ID should be uint, but sending a string
	invalidJson := []byte(`{"id": "invalid", "name": "Invalid User", "email": "invalid@example.com"}`)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(invalidJson))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response models.APIResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "error", response.Status)
	assert.Equal(t, "Invalid request data", response.Message)
	assert.Equal(t, "Gabriel Demetrios Lafis", response.Author)
}

