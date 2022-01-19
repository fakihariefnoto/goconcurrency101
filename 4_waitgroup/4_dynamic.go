package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Masalah 2")
	fmt.Println("Print 1 - 1000 dengan goroutine tanpa perlu berurutan dan print finish jika sudah selesai")

	var wg sync.WaitGroup

	wg.Add(1000)
	for urutan := 1; urutan <= 1000; urutan++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(urutan)
	}
	wg.Wait()
	fmt.Println("Finish")

}
