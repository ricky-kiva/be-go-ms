package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRouterForPing() *gin.Engine {
	r := gin.Default()

	// make endpoint `/ping` with method `GET`
	r.GET("/ping", func(ctx *gin.Context) {
		// ctx.String(http.StatusOK, "Pong") // will respond status OK & string "Pong"
		ctx.JSON(http.StatusOK, gin.H{ // will respond status OK & a JSON
			"status": "success",
			"value":  "Pong",
		})
	})

	return r
}
