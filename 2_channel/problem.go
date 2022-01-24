package main

import (
	"fmt"
)

func main() {
	Problem()
}

func Problem() {
	var n int

	for i := 0; i < 100; i++ {
		go sum(&n, i)
	}

	fmt.Println("total n", n)
}

func sum(n *int, i int) {
	*n += i
}
