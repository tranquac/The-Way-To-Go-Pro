package main

import (
	"fmt"
	"strconv"
)

type TowInts struct {
	a, b int
}

func main() {
	two := new(TowInts)
	two.a = 12
	two.b = 100
	fmt.Printf("two is: %v\n", two)
	fmt.Println("two is:", two)
	fmt.Printf("two is: %T\n", two)
	fmt.Printf("two is: %#v\n", two)
}

func (tn *TowInts) String() string {
	return "(" + strconv.Itoa(tn.a) + " / " + strconv.Itoa(tn.b) + ")"
}
