package goroutine

import (
	"fmt"
	"time"
)

func work() {
	fmt.Println("Percobaan 2")

	go func() {
		fmt.Println("Output dari goroutine baru")
	}()

	fmt.Println("Output dari main goroutine")
	time.Sleep(1000 * time.Millisecond)
}
