package main

func main() {
	// intro gin by making "/ping" endPoint
	r := setRouterForPing()
	r.Run(":8080") // run the endPoint on port `:8080`
}
