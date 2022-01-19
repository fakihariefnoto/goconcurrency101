package main

import (
	"fmt"
)

func notwork() {
	fmt.Println("Percobaan 1")

	go func() {
		fmt.Println("Output dari goroutine baru")
	}()

	fmt.Println("Output dari main goroutine")
}
