package main

import (
	"user-jwt-auth/controllers"
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

	r.POST("/user-registration", controllers.SignUp)

	err := r.Run(":8080")
	if err != nil {
		panic("could not start the server")
	}
}
