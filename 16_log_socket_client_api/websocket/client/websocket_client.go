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

	// `Dial` connect to the address on the named network
	connection, err := net.Dial(SERVER_TYPE, fmt.Sprintf("%s:%s", SERVER_HOST, SERVER_PORT))
	if err != nil {
		panic("Error in dialing!")
	}

	defer connection.Close()

	// writes data to the connection
	_, err = connection.Write([]byte("Hey you, don't tell me there's no hope at all"))
	if err != nil {
		fmt.Println("Error writing data!")
	}

	// --- SECOND STAGE

	// read message from server
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading data!")
	}
	fmt.Printf("Received: %s\n", string(buffer[:mLen]))
}
