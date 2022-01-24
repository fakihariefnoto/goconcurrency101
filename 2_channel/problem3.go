package main

import (
	"fmt"
)

func main() {
	Problem3()
}

// Problem3 sum with channel
func Problem3() {
	sqrch := make(chan int)
	cubech := make(chan int)

	go calcSquares(sqrch)
	go calcCubes(cubech)

	squares, cubes := <-sqrch, <-cubech

	fmt.Println(squares, cubes)
	fmt.Println("Final output", squares+cubes)
}

func calcSquares(squareop chan int) {
	var sum int
	for i := 1; i < 5; i++ {
		sum += i * i
	}
	squareop <- sum
}

func calcCubes(cubeop chan int) {
	var sum int
	for i := 1; i < 5; i++ {
		sum += i * i * i
	}
	cubeop <- sum
}
