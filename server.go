// Golang HTTP server for testing


package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer ln.Close()
	fmt.Println("HTTP Server started on 127.0.0.1:8080")

	for{
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		

		fmt.Println("[Connected] from", conn.RemoteAddr())
		response := "HTTP/1.1 200 OK\r\nContent-Length: 13\r\n\r\nTesting Server"
		conn.Write([]byte(response))
		conn.Close()
	}
}
