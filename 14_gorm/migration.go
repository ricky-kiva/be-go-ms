package main

import (
	"14_gorm/models"

	"github.com/jinzhu/gorm"
)

func migrateMethod(DB *gorm.DB) {
	DB.AutoMigrate(models.Character{})
}
