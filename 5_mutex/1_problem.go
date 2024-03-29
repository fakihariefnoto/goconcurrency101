package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	x = x + 1
	return
}
func main() {
	var w sync.WaitGroup

	fmt.Println("Pengulangan sebanyak 10k untuk menambahkan variable x")
	for i := 0; i < 10000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()

	fmt.Println("Nilai dari x", x)
}
