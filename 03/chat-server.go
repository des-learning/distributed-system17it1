package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"regexp"
)

var scanUsername = regexp.MustCompile(`My name is: (\w+)`)

func hello(conn net.Conn) error {
	scanner := bufio.NewScanner(conn)
	writer := bufio.NewWriter(conn)
	for scanner.Scan() {
		received := scanner.Text()
		if received != "Hello, server" {
			return errors.New("invalid hello")
		}
		writer.WriteString("Hello, please say your name")
		writer.Flush()
		break
	}
	return nil
}

func getUsername(conn net.Conn) (string, error) {
	var received string
	fmt.Fscanf(conn, "%s", &received)
	username := scanUsername.FindStringSubmatch(received)
	if len(username) != 2 {
		return "", errors.New("expected username")
	}
	return username[1], nil
}

type userConnection struct {
	username string
	conn     net.Conn
}

var userList = map[string]userConnection{}

func register(username string, conn net.Conn) userConnection {
	uc := userConnection{username, conn}
	// add uc to user list
	userList[username] = uc
	return uc
}

func chatLoop(user userConnection) {
	for {
		var received string
		_, err := fmt.Fscanf(user.conn, "%s", &received)
		if err != nil {
			return
		}
		// write to all other user
		for _, u := range userList {
			if u != user {
				fmt.Fprintf(u.conn, "%s: %s\n", user.username, received)
			}
		}
	}
}

func handleChat(conn net.Conn) {
	// handshake phase
	err := hello(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	username, err := getUsername(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	chatUser := register(username, conn)
	chatLoop(chatUser)
}

func main() {
	sock, err := net.Listen("tcp", ":9090")
	if err != nil {
		return
	}
	defer sock.Close()
	fmt.Println("server is ready")

	for {
		conn, err := sock.Accept()
		if err != nil {
			break
		}
		fmt.Println("client connected")
		defer conn.Close()

		handleChat(conn)
		fmt.Println("client disconnected")
	}
}
