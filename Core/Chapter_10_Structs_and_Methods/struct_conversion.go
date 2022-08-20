package main

import "fmt"

type number struct {
	f float64
}

func main() {
	a := number{5.0}
	b := number{5.0}
	var c = number(b)
	fmt.Println(a, b, c)
}
