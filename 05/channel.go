package main

import "fmt"

func kali(n chan int) {
	/* for i := range n {
		fmt.Println(i * 10)
	} */
	for {
		i, ok := <-n
		if !ok {
			break
		}
		fmt.Println(i * 10)
	}
}

func main() {
	a := make(chan int)

	go kali(a)

	for _, x := range []int{10, 3, 4, 2, 1, 5, 8} {
		a <- x
	}
}
