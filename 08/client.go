package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"des.com/hellogrpc/hello"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connect to server: %v", err)
	}
	client := hello.NewHelloClient(conn)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
		defer cancel()
		response, err := client.SayHello(ctx, &hello.HelloRequest{Name: text})
		if err != nil {
			fmt.Fprintf(os.Stderr, "errof request say hello to server: %v\n", err)
			continue
		}
		fmt.Printf("Response: %s\n", response.Message)
	}

}
