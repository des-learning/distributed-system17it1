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
		writer.WriteString("Hello, please say your name\n")
		writer.Flush()
		break
	}
	return nil
}

func getUsername(conn net.Conn) (string, error) {
	scanner := bufio.NewScanner(conn)
	writer := bufio.NewWriter(conn)
	var username []string
	for scanner.Scan() {
		received := scanner.Text()
		username = scanUsername.FindStringSubmatch(received)
		if len(username) != 2 {
			return "", errors.New("expected username")
		}
		writer.WriteString(fmt.Sprintf("Hello, %s\n", username[1]))
		writer.Flush()
		break
	}
	return username[1], nil
}

type userConnection struct {
	username string
	conn     net.Conn
}

func (u userConnection) send(username string, text string) {
	writer := bufio.NewWriter(u.conn)
	writer.WriteString(fmt.Sprintf("%s: %s\n", username, text))
	writer.Flush()
}

var userList = []userConnection{}

func register(username string, conn net.Conn) userConnection {
	uc := userConnection{username, conn}
	userList = append(userList, uc)
	return uc
}

func chatLoop(user userConnection) {
	scanner := bufio.NewScanner(user.conn)
	writer := bufio.NewWriter(user.conn)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "bye" {
			writer.WriteString(fmt.Sprintf("bye, %s\n", user.username))
			writer.Flush()
			fmt.Println("client disconnected")
			break
		}
		for _, otherUser := range userList {
			if otherUser != user {
				otherUser.send(user.username, text)
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
	chatUser.conn.Close()
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

		go handleChat(conn)
	}
}
