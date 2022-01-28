package main

import (
	"fmt"
	"time"
)

var (
	printNow chan bool
)

func main() {
	printNow = make(chan bool)

	go printer()
	go sender()

	for {
	}
}

func printer() {
	for {
		if <-printNow {
			fmt.Println("Hey i got command from other go routinr to print this!")
		}
	}
}

func sender() {
	for {
		for i := 0; i < 100; i++ {
			printNow <- true
			time.Sleep(500 * time.Millisecond)
		}
	}
}
