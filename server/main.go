package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	// Here we would typically read from the connection and handle messages.
	for {
		buf := make([]byte, 1024*4)
		fmt.Println("Reading the data sent by the client...")
		n, err := conn.Read(buf[:4])
		fmt.Println(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn.Read err =", err)
			return
		}
	}
}

func main() {
	fmt.Println("Server is listening on port 8889")
	listen, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Listen err =", err)
	}
	// Once the server successfully listens, it should wait for client connections.
	for {
		fmt.Println("Waiting for the client to connect to the server...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err =", err)
		}
		// Once a client connects, start a goroutine to handle the connection.
		go process(conn)
	}
}
