package main

import "fmt"

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float64
}

func main() {
	b := B{A{6, 9}, 6.2, 9.1}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}