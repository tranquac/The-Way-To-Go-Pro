package main

import "fmt"

func main() {
	var n int16 = 18
	var m int32

	//m = n -> error
	m = int32(n)

	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)
}