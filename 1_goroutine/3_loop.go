package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Percobaan 2")

	go loop1(10) // with param
	go loop2()   // without param

	fmt.Println("Output dari main goroutine")
	time.Sleep(1000 * time.Millisecond)

}

func loop1(x int) {
	for i := 0; i < x; i++ {
		fmt.Println("Dari loop 1 -> ", i)
	}
}

func loop2() {
	for i := 10; i > 0; i-- {
		fmt.Println("Dari loop 2 -> ", i)
	}
}
