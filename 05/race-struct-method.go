package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type shared struct {
	counter int
	mutex   sync.Mutex
}

func (s *shared) inc(worker string) {
	s.mutex.Lock()
	prev := a.counter
	s.counter = s.counter + 1
	fmt.Printf("%s: %d %d\n", worker, prev, s.counter)
	s.mutex.Unlock()
	n := time.Duration(rand.Intn(100)) * time.Millisecond
	time.Sleep(n)
}

func (s *shared) dec(worker string) {
	s.mutex.Lock()
	prev := s.counter
	s.counter = s.counter - 1
	fmt.Printf("%s: %d %d\n", worker, prev, s.counter)
	s.mutex.Unlock()
	n := time.Duration(rand.Intn(100)) * time.Millisecond
	time.Sleep(n)
}

var a = &shared{} // a = new shared()

func work(name string, n int, add bool) {
	for i := 0; i < n; i++ {
		if add {
			a.inc(name)
		} else {
			a.dec(name)
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
	fmt.Printf("counter: %d\n", a.counter)
}
