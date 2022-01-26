package main

import (
	"fmt"
)

func main() {

	// create a buffered channel
	// with a capacity of 5.
	ch := make(chan string, 5)
	ch <- "First Message !"
	ch <- "Second Message !"
	ch <- "Third Message !"
	ch <- "Fourth Message !"
	ch <- "Fifth Message !"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
