package main

import (
	"duckAPI/controller/middleware"
	"duckAPI/controller/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	validApiKey := os.Getenv("VALID_API_KEY")

	r := gin.Default()
	r.Use(middleware.ApiKeyAuthMiddleware(validApiKey))

	routes.InitRoutes(&r.RouterGroup)

	apiListenPort := "0.0.0.0:" + port

	r.Run(apiListenPort)
}
