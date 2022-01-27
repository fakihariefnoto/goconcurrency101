package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "first messages"
	ch <- "second messages"
	ch <- "deadlock messages"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
