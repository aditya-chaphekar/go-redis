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
	connection, err := listner.Accept()
	if err != nil {
		fmt.Println("Error Accepting Connection")
		os.Exit(1)
	}
	buff := make([]byte, 1024)
	length, err := connection.Read(buff)
	returnBuff := []byte("+PONG\r\n")
	connection.Write(returnBuff)
	for i := 0; i <= length; i++ {
		if string(buff[i]) == "\n" {
			connection.Write(returnBuff)
		}
	}

}
