package main

import "fmt"

func main() {
	slFrom := []int{1, 2, 3, 4}
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo) // output: [1 2 3 4 0 0 0 0 0 0]
	fmt.Printf("Copied %d element\n", n) //n = 3

	sl3 := []int{1, 2, 3, 4}
	sl3 = append(sl3, 5,6,7)
	fmt.Println(sl3)
}
