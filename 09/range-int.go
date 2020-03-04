package main

import (
	"fmt"
	"log"
)

type IntRange struct {
	start int
	end   int
	delta int
}

func (r IntRange) Stream() (<-chan int, error) {
	out := make(chan int, 1)
	switch {
	case r.end >= r.start && r.delta >= 0:
		go func() {
			for i := r.start; i < r.end; i += r.delta {
				out <- i
			}
			close(out)
		}()
		return out, nil
	case r.end < r.start && r.delta < 0:
		go func() {
			for i := r.start; i > r.end; i += r.delta {
				out <- i
			}
			close(out)
		}()
		return out, nil
	}
	return nil, fmt.Errorf("invalid range")
}

func main() {
	a, err := IntRange{1, 1000000000, 1}.Stream()
	if err != nil {
		log.Fatal(err)
	}

	for i := range a {
		fmt.Println(i)
	}
}
