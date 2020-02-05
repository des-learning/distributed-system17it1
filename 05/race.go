package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var counter int = 0
var mutex = sync.Mutex{}

func inc(worker string) {
	mutex.Lock()
	prev := counter
	counter = counter + 1
	fmt.Printf("%s: %d %d\n", worker, prev, counter)
	mutex.Unlock()
	n := time.Duration(rand.Intn(100)) * time.Millisecond
	time.Sleep(n)
}

func dec(worker string) {
	mutex.Lock()
	prev := counter
	counter = counter - 1
	fmt.Printf("%s: %d %d\n", worker, prev, counter)
	mutex.Unlock()
	n := time.Duration(rand.Intn(100)) * time.Millisecond
	time.Sleep(n)
}

func work(name string, n int, add bool) {
	for i := 0; i < n; i++ {
		if add {
			inc(name)
		} else {
			dec(name)
		}
	}
}

type test struct {
	n   int
	add bool
}

func main() {
	workers := map[string]test{
		"W1": {5, false},
		"W2": {10, true},
		"W3": {7, true},
		"W4": {4, false}}
	for name, value := range workers {
		go work(name, value.n, value.add)
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("counter: %d\n", counter)
}
