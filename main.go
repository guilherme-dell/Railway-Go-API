package main

import (
	"duckAPI/controller/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	fmt.Printf("PORT=%s\n", port)

	r := gin.Default()
	routes.InitRoutes(&r.RouterGroup)

	apiListenPort := "0.0.0.0:" + port

	r.Run(apiListenPort)
}
