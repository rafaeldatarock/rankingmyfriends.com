package main

import (
	"api/database"
	"api/server"

	"github.com/gin-gonic/gin"
)

func main() {
	//Initialize db
	database.OpenConnection()
	database.MigrateModels()

	//Start webserver
	router := gin.Default()
	server.InitializeRoutes(router)
	router.Run("localhost:8080")
}
