package routes

import (
	"go-rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(route *gin.RouterGroup) {
	users := route.Group("/users")
	{
		users.GET("/", controllers.GetUsers)
		users.POST("/", controllers.PostUsers)
		users.PUT("/:id", controllers.PutUsers)
		users.DELETE("/:id", controllers.DeleteUsers)
	}
}
