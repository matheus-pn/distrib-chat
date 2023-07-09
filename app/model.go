package app

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Username string `gorm:"not null"`
	Message  string `gorm:"not null"`
	MediaUrl string
}
