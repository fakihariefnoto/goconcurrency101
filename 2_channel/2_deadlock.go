package main

import (
	"fmt"
)

func main() {
	n := make(chan int)
	n <- 1 // data stuck here, there is no cap
	fmt.Println("Total n", <-n)
}
