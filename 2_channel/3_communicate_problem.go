package main

import (
	"fmt"
	"time"
)

var (
	printNow bool
)

func main() {
	printNow = false

	go printer()
	go sender()

	for {
	}
}

func printer() {
	for {
		if printNow {
			fmt.Println("Hey variable changed, i need to print this!")
			printNow = false
		}
	}
}

func sender() {
	for {
		printNow = true
		time.Sleep(1000 * time.Millisecond)
	}
}
