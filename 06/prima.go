package main

import (
	"fmt"
	"math"
	"sync"
)

type primaResult struct {
	n       int
	isPrima bool
}

func prima(in <-chan int, out chan<- primaResult) {
	for n := range in {
		p := true
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				p = false
				break
			}
		}
		out <- primaResult{n, p}
	}
}

func gatherResult(res <-chan primaResult) {

	for r := range res {
		if !r.isPrima {
			continue
		}
		fmt.Println(r.n)
	}
}

func main() {
	t := 512
	jobChan := make(chan int, t)
	resultChan := make(chan primaResult, t)

	n := 8

	primaWg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		primaWg.Add(1)
		go func() {
			defer primaWg.Done()
			prima(jobChan, resultChan)
		}()
	}

	resultWg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		resultWg.Add(1)
		go func() {
			defer resultWg.Done()
			gatherResult(resultChan)
		}()
	}

	for i := 1; i <= 10000000; i++ {
		jobChan <- i
	}

	close(jobChan)
	primaWg.Wait()

	close(resultChan)
	resultWg.Wait()
}
