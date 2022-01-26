package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "geeksforgeeks"
	ch <- "hello"
	ch <- "geeks"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
