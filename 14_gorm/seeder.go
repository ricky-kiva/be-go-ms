package main

import (
	"14_gorm/models"

	"github.com/jinzhu/gorm"
)

func seederMethod(DB *gorm.DB) {
	data := models.Character{
		Id:     1,
		Name:   "Rickyslash",
		Level:  8,
		Gender: "Male",
	}

	DB.Create(&data)
}
