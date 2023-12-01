package main

import (
	"14_gorm/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	DB, err := gorm.Open(
		"postgres",
		"postgresql://postgres:postgres_pw@127.0.0.1/postgres?sslmode=disable",
	)

	if err != nil {
		panic("DB connection failed")
	}

	fmt.Println("DB connection success")

	var seederCallback = func() {
		data := models.Character{}

		// fetches all records from "relevant table" inside the `DB`
		// return `true` if there is no record inside the "relevant table"
		if DB.Find(&data).RecordNotFound() {
			fmt.Println("Executing Seeder method")
			seederMethod(DB)
		}
	}

	migrateMethod(DB, seederCallback)

	defer DB.Close()
}
