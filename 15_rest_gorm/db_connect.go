package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func connectDB() (*gorm.DB, error) {
	conn := "postgresql://postgres:postgres_pw@127.0.0.1/postgres?sslmode=disable"
	db, err := gorm.Open("postgres", conn)
	return db, err
}
