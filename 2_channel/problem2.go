package main

import (
	"fmt"
)

func main() {
	Problem2()
}

// Problem2 Deadlock
func Problem2() {
	n := make(chan int)
	n <- 1
	fmt.Println("Total n", <-n)
}
