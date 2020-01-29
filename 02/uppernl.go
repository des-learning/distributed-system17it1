package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//jalan upper.go, simpan stdout
	//jalan nl.go dengan stdin-nya di set dengan stdout dari upper.go, simpan stdout
	//tampilkan stdout dar nl.go
	fmt.Println("spawn process")

	upper := exec.Command("./upper")
	// pasang stdin upper dengan stdin program main
	upper.Stdin = os.Stdin
	upperStdout, _ := upper.StdoutPipe()
	defer upperStdout.Close() // close stdout when finished

	// pasang stdin nl dengan stdout upper, pasang stdout upper ke stdout main
	nl := exec.Command("./nl")
	nl.Stdin = upperStdout
	nl.Stdout = os.Stdout

	// jalankan command
	fmt.Println("run upper")
	upper.Start()
	fmt.Println("run nl")
	nl.Start()

	// tunggu sampai selesai
	upper.Wait()
	nl.Wait()
}
