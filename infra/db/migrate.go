package db

import (
	"fmt"
	"note_pad/models"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	fmt.Println("✅ Migrations applied")
	return nil
}
