package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers returns all users
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of users"})
}

// GetUserByID returns a specific user
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "User details", "id": id})
}

// CreateUser creates a new user
func CreateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// UpdateUser updates a user
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "User updated", "id": id})
}

// DeleteUser deletes a user
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "User deleted", "id": id})
}
