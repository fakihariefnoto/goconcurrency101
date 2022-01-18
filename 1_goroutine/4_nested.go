package goroutine

import (
	"fmt"
	"time"
)

func nested() {
	fmt.Println("Percobaan 2")

	go loop1()

	fmt.Println("Output dari main goroutine")
	time.Sleep(2000 * time.Millisecond)

}

func nestedloop1() {
	for i := 1; i < 5; i++ {
		fmt.Println("Dari loop 1 -> ", i)
		go nestedloop2(i)
	}
}

func nestedloop2(x int) {
	for i := x; i < 5; i++ {
		fmt.Printf("Dari loop 2-%v -> %v\n", x, i)
	}
}
