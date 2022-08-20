package main

import "fmt"

type TowInts struct {
	a int
	b int
}

func main() {
	a := TowInts{4, 89}
	// b := new(TowInts)
	fmt.Printf("The sum is: %d\n", a.AddThem())
	fmt.Printf("The sum and param is: %d\n", a.AddToParam(6))
}

func (tn *TowInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TowInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
