package main

import "fmt"

func IntRange(start int, end int, delta int) <-chan int {
	out := make(chan int, 1)
	switch {
	case start <= end && delta >= 0:
		go func() {
			for i := start; i < end; i += delta {
				out <- i
			}
			close(out)
		}()
	case start >= end && delta < 0:
		go func() {
			for i := start; i > end; i += delta {
				out <- i
			}
			close(out)
		}()
	default:
		go func() {
			close(out)
		}()
	}
	return out
}

func IntSlice(slice []int) <-chan int {
	out := make(chan int, 1)
	go func() {
		for _, i := range slice {
			out <- i
		}
		close(out)
	}()
	return out
}

func IntFilter(numbers <-chan int, fun func(int) bool) <-chan int {
	out := make(chan int, 1)
	go func() {
		for i := range numbers {
			if fun(i) {
				out <- i
			}
		}
		close(out)
	}()
	return out
}

func IntMap(numbers <-chan int, fun func(int) int) <-chan int {
	out := make(chan int, 1)
	go func() {
		for i := range numbers {
			out <- fun(i)
		}
		close(out)
	}()
	return out
}

func IntReduce(start int, numbers <-chan int, fun func(int, int) int) int {
	result := start
	for i := range numbers {
		result = fun(result, i)
	}
	return result
}

func main() {
	/*for i := range IntRange(1, 100, 1) {
		fmt.Println(i)
	}
	for i := range IntRange(100, 1, -1) {
		fmt.Println(i)
	}*/
	/*for i := range IntSlice([]int{1, 2, 3, 4, 5, 6, 7}) {
		fmt.Println(i)
	}*/
	satuSampai999 := IntRange(1, 1000, 1)
	ganjil := func(x int) bool { return x%2 != 0 }
	genap := func(x int) bool { return x%2 == 0 }
	/*for i := range IntFilter(satuSampai999, ganjil) {
		fmt.Println(i)
	}*/
	/*kali3 := func(x int) int { return x * 3 }
	for i := range IntMap(satuSampai999, kali3) {
		fmt.Println(i)
	}*/
	sum := func(x int, y int) int { return x + y }
	//fmt.Println(IntReduce(0, satuSampai999, sum))
	fmt.Println(IntReduce(0, IntFilter(satuSampai999, ganjil), sum))
	satuSampai999 = IntRange(1, 1000, 1)
	fmt.Println(IntReduce(0, IntFilter(satuSampai999, genap), sum))
}
