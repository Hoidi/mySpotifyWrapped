package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitMyServer() {
	r := gin.Default()

	r.GET("/test", get())
	r.GET("/", rest())

	r.Run(":4000")
}

func get() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.JSON(http.StatusOK, gin.H{
			"message": "Test okej",
		})
	}
}

func rest() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Not test",
		})
	}
}
