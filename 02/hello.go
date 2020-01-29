package main

import "fmt"

func hello(name string) {
	fmt.Println("Hello " + name)
}

func add(a int, b int) int {
	return a + b
}

func main() {
	//var world string
	//world = "World"
	//var world string = "World"
	world := "World" // type inference
	hello(world)
	world = "budi"
	hello(world)
	fmt.Println(add(10, 13))
}
