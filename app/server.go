package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	listner, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	for {
		connection, err := listner.Accept()
		if err != nil {
			fmt.Println("Error Accepting Connection")
			os.Exit(1)
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	buffer := make([]byte, 1024)
	for {
		_, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from connection")
			return
		}
		_, err = connection.Write([]byte("+PONG\r\n"))
		if err != nil {
			fmt.Println("Error writing to connection")
			return
		}
	}

}
