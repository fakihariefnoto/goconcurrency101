package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Percobaan 2")

	go loop1()

	fmt.Println("Output dari main goroutine")
	time.Sleep(2000 * time.Millisecond)

}

func loop1() {
	for i := 1; i < 5; i++ {
		fmt.Println("Dari loop 1 -> ", i)
		go loop2(i)
	}
}

func loop2(x int) {
	for i := x; i < 5; i++ {
		fmt.Printf("Dari loop 2-%v -> %v\n", x, i)
	}
}
