package main

import "fmt"

type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func main() {
	fmt.Println(IntVector{1, 2, 3, 4}.Sum())
	v2  := IntVector{1, 2, 3, 10}
	fmt.Println(v2.Sum())
}
