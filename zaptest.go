package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
func main() {

	log, _ := zap.NewProduction()
	log.Warn("Warning Test")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
