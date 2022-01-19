package main

import (
	"fmt"
	"time"
)

func loop() {
	fmt.Println("Percobaan 2")

	go loop1()
	go loop2()

	fmt.Println("Output dari main goroutine")
	time.Sleep(1000 * time.Millisecond)

}

func loop1() {
	for i := 0; i < 10; i++ {
		fmt.Println("Dari loop 1 -> ", i)
	}
}

func loop2() {
	for i := 10; i > 0; i-- {
		fmt.Println("Dari loop 2 -> ", i)
	}
}
