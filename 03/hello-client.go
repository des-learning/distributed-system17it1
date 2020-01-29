package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("error connection to server")
		return
	}
	defer conn.Close()
	fmt.Printf("Connected to server at %s\n",
		conn.RemoteAddr().String())

	for {
		var text string
		_, err := fmt.Fscanf(conn, "%s\n", &text)
		if err != nil {
			break
		}
		fmt.Println("Received from server: ", text)
	}
}
