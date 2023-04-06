package main

import (
	"user-jwt-auth/intitializers"

	"github.com/gin-gonic/gin"
)

func init() {
	intitializers.LoadConfigVariales()
	intitializers.ConnectDb()
	intitializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
