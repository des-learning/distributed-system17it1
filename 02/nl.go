package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	line := 0
	for scanner.Scan() {
		line++
		fmt.Printf("%4d %s\n", line, scanner.Text())
	}
}
