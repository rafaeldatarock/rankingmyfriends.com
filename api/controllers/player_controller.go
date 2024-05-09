package controllers

import (
	"net/http"

	"api/repositories"

	"github.com/gin-gonic/gin"
)

func GetPlayer(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, repositories.GetPlayer())
}

func NewPlayer(c *gin.Context) {
	repositories.NewPlayer()
	c.IndentedJSON(http.StatusOK, "OK")
}
