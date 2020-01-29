package main

import (
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	fmt.Println("client connected")
	for {
		var text string
		_, err := fmt.Fscanf(conn, "%s\n", &text)
		if err != nil {
			break
		}
		fmt.Fprintf(conn, "%s\n", strings.ToUpper(text))
	}
	fmt.Println("client disconnected")
}

func main() {
	sock, err := net.Listen("tcp", ":9090")
	if err != nil {
		return
	}
	fmt.Println("server ready")
	defer sock.Close()

	for {
		conn, err := sock.Accept()
		if err != nil {
			break
		}
		defer conn.Close()

		go handleConnection(conn)
	}
}
