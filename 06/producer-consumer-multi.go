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

func producer(c chan string, n int, name string) {
	fmt.Printf("producer %s running\n", name)
	for i := 0; i < n; i++ {
		c <- randomString(rand.Intn(10) + 1)
	}
	fmt.Printf("producer %s finished\n", name)
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

	d := []string{"D1", "D2"}
	wgD := sync.WaitGroup{}
	wgD.Add(len(d))
	for _, i := range d {
		go func(n string) {
			defer wgD.Done()
			countChars(o, n)
		}(i)
	}

	wgCons := sync.WaitGroup{}
	consumers := []string{"C1", "C2"}
	wgCons.Add(len(consumers))
	for _, con := range consumers {
		go func(n string) {
			defer wgCons.Done()
			consumer(c, o, n)
		}(con)
	}

	p := []string{"P1", "P2", "P3", "P4"}
	wg := sync.WaitGroup{}
	wg.Add(len(p))
	for _, i := range p {
		go func(n string) {
			defer wg.Done()
			producer(c, 1000, n)
		}(i)
	}

	wg.Wait()
	close(c)

	wgCons.Wait()
	close(o)

	wgD.Wait()
}
