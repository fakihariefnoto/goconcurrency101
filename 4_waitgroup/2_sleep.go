package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Masalah 1")
	a := 100
	b := 150
	c := 10
	d := 1000
	fmt.Println("Hasil dari add(a)+sub(b)+mul(c)+div(d) adalah = ")
	go add(&a)
	go sub(&b)
	go mul(&c)
	go div(&d)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Hasilnya ", a+b+c+d)

	a = 100
	b = 150
	c = 10
	d = 1000
	fmt.Println("Hasil yang diharapkan : 200 + 100 + 100 + 100 = ", add(&a)+sub(&b)+mul(&c)+div(&d))
}

// add function to add x with 100
func add(x *int) int {
	*x += 100
	return *x
}

// sub function to substract x with 50
func sub(x *int) int {
	*x -= 50
	return *x
}

// mul function to multiply x with 10
func mul(x *int) int {
	*x *= 10
	return *x
}

// mul function to div x with 10
func div(x *int) int {
	*x /= 10
	return *x
}
