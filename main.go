package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const defaultMessage = "Hello!"

func main() {
	// Initialize Go Gin
	engine := gin.Default()

	// Setup a simple endpoint
	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, defaultMessage)
		return

	})

	engine.Run()
}
