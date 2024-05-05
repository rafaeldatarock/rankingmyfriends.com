package main

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users", controllers.GetUsers)
	router.Run("localhost:8080")
}
