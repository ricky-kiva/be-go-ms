package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRouterForPing(r *gin.Engine) {
	// make endpoint `/ping` with method `GET`
	r.GET("/ping", func(ctx *gin.Context) {

		// will respond status OK & string "Pong"
		// ctx.String(http.StatusOK, "Pong")

		// will respond status OK & a JSON
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
			"value":  "Pong",
		})
	})
}
