package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

const charset = "abcdefghijklmnopqrstuvwxyz"

func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(
			len(charset))]
	}
	return string(b)
}

func producer(c chan string, n int) {
	fmt.Println("producer running")
	for i := 0; i < n; i++ {
		c <- randomString(rand.Intn(10) + 1)
	}
	fmt.Println("producer finished")
}

func consumer(c chan string, o chan string, name string) {
	fmt.Printf("consumer %s running\n", name)
	for s := range c {
		o <- fmt.Sprintf("%s: %s", name, strings.ToUpper(s))
	}
	fmt.Printf("consumer %s finished\n", name)
}

func countChars(c chan string, name string) {
	fmt.Printf("countChars %s running\n", name)
	for s := range c {
		fmt.Printf("%s: %d %s\n", name, len(s), s)
	}
	fmt.Printf("countChars %s finished\n", name)
}

func main() {
	c := make(chan string)
	o := make(chan string)

	wgD := sync.WaitGroup{}
	wgD.Add(1)
	go func(n string) {
		defer wgD.Done()
		countChars(o, n)
	}("D1")

	wgCons := sync.WaitGroup{}
	consumers := []string{"C1", "C2"}
	wgCons.Add(len(consumers))
	for _, con := range consumers {
		go func(n string) {
			defer wgCons.Done()
			consumer(c, o, n)
		}(con)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		producer(c, 1000)
	}()

	wg.Wait()
	close(c)

	wgCons.Wait()
	close(o)

	wgD.Wait()
}
