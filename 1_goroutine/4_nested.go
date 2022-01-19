package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Percobaan 3")

	go loop1(5)

	fmt.Println("Output dari main goroutine")
	time.Sleep(2000 * time.Millisecond)

}

func loop1(max int) {
	for i := 1; i < max; i++ {
		fmt.Println("Dari loop 1 -> ", i)
		go loop2(i)
	}
}

func loop2(min int) {
	for i := min; i < 5; i++ {
		fmt.Printf("Dari loop 2-%v -> %v\n", x, i)
	}
}
