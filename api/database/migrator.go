package database

import (
	"api/models"
)

func MigrateModels() {
	db.AutoMigrate(&models.Player{})
}
