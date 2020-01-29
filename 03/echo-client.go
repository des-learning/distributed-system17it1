package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		return
	}
	fmt.Println("connected to server")
	defer conn.Close()

	for {
		var text string
		_, err := fmt.Scanf("%s\n", &text)
		if err != nil {
			break
		}
		_, err = fmt.Fprintf(conn, "%s\n", text)
		if err != nil {
			break
		}
		var input string
		_, err = fmt.Fscanf(conn, "%s\n", &input)
		if err != nil {
			break
		}
		fmt.Printf("received from server: %s\n", input)
	}
}
