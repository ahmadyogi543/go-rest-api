package main

import (
	"go-rest-api/routes"
	"go-rest-api/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadDotenv()
}

func main() {
	app := gin.Default()
	v1 := app.Group("/api/v1")

	routes.RootRoutes(v1)
	routes.UsersRoutes(v1)

	app.Run()
}
