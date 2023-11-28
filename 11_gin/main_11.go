package main

import "github.com/gin-gonic/gin"

func main() {
	// initiate Gin engine
	r := gin.Default()

	// make "/ping" endPoint
	setRouterForPing(r)

	// make "/v1" group
	setRouterGroupV1(r)

	r.Run(":8080") // run router on port `:8080`
}
