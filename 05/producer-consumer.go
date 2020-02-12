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
	for i := 0; i < n; i++ {
		c <- randomString(8)
	}
	close(c)
}

func consumer(c chan string, name string) {
	for s := range c {
		fmt.Println(name, strings.ToUpper(s))
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	c := make(chan string)

	wg := sync.WaitGroup{}
	consumers := []string{"C1", "C2"}
	wg.Add(len(consumers) + 1)
	for _, con := range consumers {
		go func(n string) {
			defer wg.Done()
			consumer(c, n)
		}(con)
	}

	go func() {
		defer wg.Done()
		producer(c, 1000000)
	}()

	wg.Wait()
}
