package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
func main() {

	// creates capacity of 2
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(10 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		//time.Sleep(time.Second)
	}
}
