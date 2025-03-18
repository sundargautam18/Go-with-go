package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware protects routes
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		// Validate token logic here
		c.Next()
	}
}
