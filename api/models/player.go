package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}
