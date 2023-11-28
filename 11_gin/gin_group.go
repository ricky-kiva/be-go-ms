package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SimpleUser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func setRouterGroupV1(r *gin.Engine) {
	// initiate group
	v1 := r.Group("v1")

	// initialize `user` URI with `name` Param
	v1.GET("/user/:name", func(ctx *gin.Context) {
		param := ctx.Param("name") // get `name` param from URL
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"value":  param,
		})
	})

	// POST value on `user` URI
	v1.POST("/user", func(ctx *gin.Context) {
		var data SimpleUser

		ctx.BindJSON(&data) // store data POSTed in the transaction to `data` variable

		// simulate process by re-passing as response (usually, it should be processed like saving on DB, etc.)
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"value":  data,
		})
	})

	// add Query Param `name` to `/user` URI
	v1.GET("/user", func(ctx *gin.Context) {
		query := ctx.Query("name")

		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"value":  query,
		})
	})
}
