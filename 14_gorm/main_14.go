package main

import (
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

	migrateMethod(DB)

	defer DB.Close()
}
