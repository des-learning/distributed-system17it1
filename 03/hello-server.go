package main

import (
	"fmt"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("client connected from %s\n", conn.RemoteAddr().String())
	for count := 10; count > 0; count-- {
		fmt.Fprintf(conn, "%d\n", count)
		time.Sleep(time.Second)
	}
}

func main() {
	sock, err := net.Listen("udp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println(err)
		fmt.Println("error listening to network")
		return
	}
	defer sock.Close()

	for {
		conn, err := sock.Accept()
		if err != nil {
			fmt.Println("error accepting network connection")
			return
		}
		go handleConnection(conn)
	}
}
