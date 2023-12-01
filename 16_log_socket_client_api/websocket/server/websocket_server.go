package main

import (
	"fmt"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {

	// starts TCP server
	server, err := net.Listen(SERVER_TYPE, fmt.Sprintf("%s:%s", SERVER_HOST, SERVER_PORT))
	if err != nil {
		panic("Listening error!")
	}

	fmt.Println("Server running..")

	defer server.Close()

	fmt.Printf("Listening on %s:%s\n", SERVER_HOST, SERVER_PORT)
	fmt.Println("Waiting for client..")

	for {
		// waits `for` & accepts new connection
		connection, err := server.Accept()
		if err != nil {
			panic("Error in accepting connection!")
		}

		// executed once when a new client `connection` is established
		go processClient(connection)
	}
}

func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)

	// reads data sent from the `client` and store it to `buffer`
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error in reading!")
	}

	// converts the byte in buffer "up to" `mLen`
	fmt.Printf("Received: %s\n", string(buffer[:mLen]))

	// --- SECOND STAGE

	_, err = connection.Write([]byte("Together we stand, divided we fall"))
	if err != nil {
		fmt.Println("Error writing data!")
	}
}
