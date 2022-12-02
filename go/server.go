package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.New()

	server.GET("/", func(c *gin.Context) {
		c.JSON(200,
			"Ok")
	})

	server.Run("localhost:8080")
}
