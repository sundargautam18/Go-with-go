package routes

import (
	"github.com/gin-gonic/gin"
)

// InitRoutes initializes all routes
func InitRoutes() *gin.Engine {
	r := gin.Default()

	// Define API groups
	api := r.Group("/api/v1")
	{
		UserRoutes(api)
		// ProductRoutes(api)
	}

	return r
}
