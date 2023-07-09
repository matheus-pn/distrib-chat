package app

import (
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	fmt.Println("[Database] Running AutoMigrations")
	db.AutoMigrate(&Message{})
}
