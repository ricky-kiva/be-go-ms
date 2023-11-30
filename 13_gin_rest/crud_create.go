package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func postHandler(c *gin.Context, db *sql.DB) {
	var newStudent Student

	// `Bind` will parse request body to Struct
	if c.Bind(&newStudent) == nil { //
		_, err := db.Exec(
			"insert into students values ($1, $2, $3, $4, $5)",
			newStudent.Id,
			newStudent.Name,
			newStudent.Age,
			newStudent.Address,
			newStudent.Phone,
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Exec error: " + err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Success create"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Binding Error"})
	}
}
