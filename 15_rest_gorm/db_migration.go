package main

import (
	"15_rest_gorm/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.TitanSoldier{})

	data := models.TitanSoldier{}
	if db.Find(&data).RecordNotFound() {
		fmt.Println("Executing seeder")
		seederSoldier(db)
	}
}

func seederSoldier(db *gorm.DB) {
	data := models.TitanSoldier{
		Id:             1,
		Name:           "Eren Yeager",
		Division:       "Survey Corps",
		Specialization: "Titan shifter",
	}

	db.Create(&data)
}
