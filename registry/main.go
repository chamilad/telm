package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/.well-known/terraform.json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"modules.v1": "http://localhost:8080/modules/v1/",
		})
	})

	r.Run()
}
