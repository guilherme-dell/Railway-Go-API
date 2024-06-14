package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func PrintServerENV(c *gin.Context) {
	c.String(http.StatusOK, os.Getenv("TEST_EVN"))
}
