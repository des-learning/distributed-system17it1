package main

import (
	"fmt"
	"os/exec"
)

func main() {
	out, err := exec.Command("./hello").Output()
	if err != nil {
		fmt.Println("Error running command: hello")
		return
	}
	fmt.Printf("Program output is:\n%s", string(out))
}
