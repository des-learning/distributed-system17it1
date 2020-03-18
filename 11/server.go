package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Message struct {
	Message string `json:"message"`
}

type CircuitBreaker struct {
	state int // 0 close, 1 open, 2 half open
}

func (c *CircuitBreaker) sendRequest(ctx context.Context, fn func() error) error {
	res := make(chan error, 1)
	go func() {
		res <- fn()
	}()
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout")
		case err := <-res:
			return err
		default:
		}
	}
}

func (c *CircuitBreaker) Perform(ctx context.Context, fn func() error) {
	if c.state == 1 && rand.Int31n(10)%3 == 0 {
		c.state = 2
	}
	switch c.state {
	case 0:
		log.Println("cb close, try to send request to upstream service")
		if c.sendRequest(ctx, fn) != nil {
			c.state = 1
			log.Println("service error, change cb to open state")
		}
	case 1:
		log.Println("circuit breaker open, not sending request to upstream service")
	case 2:
		log.Println("circuit breaker half open, allowing some request to check upstream service")
		if c.sendRequest(ctx, fn) != nil {
			c.state = 1
			log.Println("service still error, change cb to open state")
		} else {
			c.state = 0
			log.Println("service recovered, change cb to close state")
		}
	}
}

var cb = &CircuitBreaker{}

func handler(w http.ResponseWriter, r *http.Request) {
	message, _ := greeter("world")
	data, _ := json.Marshal(Message{Message: message})
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(data))
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		var m Message
		json.Unmarshal(body, &m)
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		var message string
		var err error
		cb.Perform(ctx, func() error {
			message, err = greeter(m.Message)
			return err
		})
		res, _ := json.Marshal(Message{Message: message})
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(res))
	}
}

func greeter(name string) (string, error) {
	time.Sleep(time.Duration(rand.Int31n(110)) * time.Millisecond)
	return "Hello " + name, nil
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
