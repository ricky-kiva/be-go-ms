package main

import (
	"database/sql"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func rowsToStruct(rows *sql.Rows, dest interface{}) error {
	// get reflection of `dest`
	// `Elem()` used when `ValueOf()` is a pointer/interface
	// - if after `ValueOf()`, it returns the "value" of the pointer/interface (dereferencing)
	destValue := reflect.ValueOf(dest).Elem()

	// `Elem()` after `Type()` return the element's type of a `reference type`
	// create `slice` type `interface{}` with number of fields of `destValue` (which is 5)
	args := make([]interface{}, destValue.Type().Elem().NumField())

	for rows.Next() {
		// make new pointer of empty struct
		rowPointer := reflect.New(destValue.Type().Elem())

		// access `rowPointer`'s value (dereference)
		rowValue := rowPointer.Elem()

		// iterately assign `pointer` of each `field` of `rowValue`
		for i := 0; i < rowValue.NumField(); i++ {
			args[i] = rowValue.Field(i).Addr().Interface()
		}

		// `rows.Scan()` reads `row` that is pointed by `rows.Next()`
		// - then pass the value to `args`
		if err := rows.Scan(args...); err != nil {
			return err
		}

		// `reflect.Append()` appends `rowValue` to `destValue` reflection
		// `.Set()` sets value of a reflection
		destValue.Set(reflect.Append(destValue, rowValue))
	}

	return nil
}

func getAllHander(c *gin.Context, db *sql.DB) {
	var students []Student

	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	rowsToStruct(rows, &students)

	if students == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "success",
			"message": "Data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   students,
	})
}

func getHandler(c *gin.Context, db *sql.DB) {
	var student []Student

	studentId := c.Param("id")

	rows, err := db.Query(
		"SELECT * from students WHERE student_id = $1",
		studentId,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
	}

	rowsToStruct(rows, &student)

	if student == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "success",
			"message": "Data not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   student,
	})
}
