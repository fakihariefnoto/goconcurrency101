package main

import (
	"fmt"
	"sync"
)

var x = 0

func incrementSolve(wg *sync.WaitGroup, c chan bool) {
	defer wg.Done()
	c <- false
	x = x + 1
	<-c
	return
}
func main() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)

	fmt.Println("Pengulangan sebanyak 10k untuk menambahkan variable x")
	for i := 0; i < 10000; i++ {
		w.Add(1)
		go incrementSolve(&w, ch)
	}
	w.Wait()

	fmt.Println("Nilai dari x", x)
}
