// Go Microservices API
// Author: Gabriel Demetrios Lafis

package main

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
)

type User struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type APIResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Author  string      `json:"author"`
}

func main() {
    r := gin.Default()
    
    // Middleware
    r.Use(func(c *gin.Context) {
        c.Header("X-Author", "Gabriel Demetrios Lafis")
        c.Next()
    })
    
    // Routes
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, APIResponse{
            Status:  "success",
            Message: "Go Microservices API is running",
            Author:  "Gabriel Demetrios Lafis",
        })
    })
    
    r.GET("/api/users", getUsers)
    r.POST("/api/users", createUser)
    r.GET("/api/health", healthCheck)
    
    log.Println("🚀 Go Microservices API starting on :8080")
    log.Println("👨‍💻 Created by Gabriel Demetrios Lafis")
    r.Run(":8080")
}

func getUsers(c *gin.Context) {
    users := []User{
        {ID: 1, Name: "Gabriel Lafis", Email: "gabriel@example.com"},
        {ID: 2, Name: "John Doe", Email: "john@example.com"},
    }
    
    c.JSON(http.StatusOK, APIResponse{
        Status:  "success",
        Message: "Users retrieved successfully",
        Data:    users,
        Author:  "Gabriel Demetrios Lafis",
    })
}

func createUser(c *gin.Context) {
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, APIResponse{
            Status:  "error",
            Message: "Invalid request data",
            Author:  "Gabriel Demetrios Lafis",
        })
        return
    }
    
    user.ID = 3 // Simulate auto-increment
    
    c.JSON(http.StatusCreated, APIResponse{
        Status:  "success",
        Message: "User created successfully",
        Data:    user,
        Author:  "Gabriel Demetrios Lafis",
    })
}

func healthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, APIResponse{
        Status:  "healthy",
        Message: "Service is running properly",
        Author:  "Gabriel Demetrios Lafis",
    })
}
