package routes

import (
	"duckAPI/controller/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/", handler.DuckSay)
	r.GET("/env", handler.PrintServerENV)
}
