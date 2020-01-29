package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func greetServer(username string, conn net.Conn) error {
	scanner := bufio.NewScanner(conn)
	writer := bufio.NewWriter(conn)

	writer.WriteString("Hello, server\n")
	writer.Flush()
	var err error
loop:
	for scanner.Scan() {
		text := scanner.Text()
		switch text {
		case "Hello, please say your name":
			writer.WriteString(fmt.Sprintf("My name is: %s\n", username))
			writer.Flush()
		case fmt.Sprintf("Hello, %s", username):
			err = nil
			break loop
		default:
			err = fmt.Errorf("error from server: %s", text)
			break loop
		}
	}
	return err
}

func handleServerResponse(username string, conn net.Conn, wg sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()
		if text == fmt.Sprintf("bye %s", username) {
			break
		}
		fmt.Println(text)
	}
}

func handleWriteToServer(conn net.Conn, wg sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(conn)

	for scanner.Scan() {
		text := scanner.Text()
		writer.WriteString(fmt.Sprintf("%s\n", text))
		writer.Flush()
	}
}

func chatLoop(username string, conn net.Conn) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go handleServerResponse(username, conn, wg)
	go handleWriteToServer(conn, wg)
	wg.Wait()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter your username: ")
	var username string
	for scanner.Scan() {
		username = scanner.Text()
		break
	}

	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		return
	}
	fmt.Println("connected to server")
	defer conn.Close()

	err = greetServer(username, conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	chatLoop(username, conn)
}
