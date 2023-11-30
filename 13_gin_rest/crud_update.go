package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func putHandler(c *gin.Context, db *sql.DB) {
	studentId := c.Param("id")
	newName := c.Query("name")

	_, err := db.Exec(
		"UPDATE students SET student_name=$1 WHERE student_id=$2",
		newName,
		studentId,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Name update success",
	})
}
