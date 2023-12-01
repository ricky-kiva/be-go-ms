package main

import (
	"14_gorm/models"

	"github.com/jinzhu/gorm"
)

func migrateMethod(DB *gorm.DB, callback func()) {
	DB.AutoMigrate(models.Character{})
	callback()
}
