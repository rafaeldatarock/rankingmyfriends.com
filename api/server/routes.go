package server

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/player", controllers.GetPlayer)
	router.GET("/newplayer", controllers.NewPlayer)
}
