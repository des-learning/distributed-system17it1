package main

import (
	"fmt"
	"time"
)

func cetak(s string) {
	for {
		fmt.Println(s)
	}
}

func main() {
	var a func() = func() { cetak("A") } // anonymous function/lambda
	b := func() { cetak("B") }
	/*go func() {
		for {
			fmt.Println("A")
		}
	}()

	go func() {
		for {
			fmt.Println("B")
		}
	}()*/
	go a()
	go b()
	time.Sleep(5 * time.Second)
}
