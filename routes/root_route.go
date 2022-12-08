package routes

import (
	"github.com/gin-gonic/gin"
)

func RootRoutes(route *gin.RouterGroup) {
	root := route.Group("/")
	{
		root.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Welcome to Go REST API example",
			})
		})

		root.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
