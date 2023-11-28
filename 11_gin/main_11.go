package main

import "github.com/gin-gonic/gin"

func main() {
	// initiate Gin engine
	r := gin.Default()

	// make "/ping" endPoint
	setRouterForPing(r)

	r.Run(":8080") // run router on port `:8080`
}
