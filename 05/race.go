package main

import (
	"fmt"
	"time"
)

var counter int = 0

func inc(worker string) {
	prev := counter
	counter = counter + 1
	fmt.Printf("%s: %d %d\n", worker, prev, counter)
	time.Sleep(100 * time.Millisecond)
}

func dec(worker string) {
	prev := counter
	counter = counter - 1
	fmt.Printf("%s: %d %d\n", worker, prev, counter)
	time.Sleep(100 * time.Millisecond)
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
	time.Sleep(5 * time.Second)
	fmt.Printf("counter: %d\n", counter)
}
