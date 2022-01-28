package main

import (
	"fmt"
	"sync"
)

var x = 0

func incrementSolve(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	x = x + 1
	mu.Unlock()
	return
}
func main() {
	var w sync.WaitGroup
	var mu sync.Mutex

	fmt.Println("Pengulangan sebanyak 10k untuk menambahkan variable x")
	for i := 0; i < 10000; i++ {
		w.Add(1)
		go incrementSolve(&w, &mu)
	}
	w.Wait()

	fmt.Println("Nilai dari x", x)
}
