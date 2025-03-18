package routes

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/sundargautam18/Go-with-go/controller"
	middlewares "github.com/sundargautam18/Go-with-go/middleware"
)

func UserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/users")
	{
		user.GET("/", middlewares.AuthMiddleware(), controllers.GetUsers)
		user.GET("/:id", controllers.GetUserByID)
		user.POST("/", middlewares.AuthMiddleware(), controllers.CreateUser)
		user.PUT("/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)
		user.DELETE("/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)
	}
}
