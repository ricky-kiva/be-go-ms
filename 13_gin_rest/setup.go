package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	// postgresql for Golang
	_ "github.com/lib/pq" // `_` means it is not directly used in the code
)

func setupRouter() *gin.Engine {
	// setup PostgreSQL connection
	// "db_type://username:password@ip_address/db_name?sslmode=`ssl_mode`"
	conn := "postgressql://postgres:postgres_pw@127.0.0.1/postgres?sslmode=disable"

	_, err := sql.Open("postgres", conn) // "postgres" is the driverName, which is PostgreSQL
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	return r
}
