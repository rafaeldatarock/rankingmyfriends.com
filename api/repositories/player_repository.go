package repositories

import (
	"api/database"
	"api/models"
)

func NewPlayer() {
	database.GetConnection().Create(&models.Player{
		Name:     "John",
		Email:    "john@doe.nl",
		Password: "notsosecret",
	})
}

func GetPlayer() string {
	var player models.Player
	database.GetConnection().First(&player, "name = ?", "John")

	return player.Name
}
