package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func delHandler(c *gin.Context, db *sql.DB) {
	studentId := c.Param("id")

	_, err := db.Exec(
		"DELETE FROM students WHERE student_id=$1",
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
		"message": "Delete success",
	})
}
