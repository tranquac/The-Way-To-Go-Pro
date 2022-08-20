package main

import "fmt"

func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	n := 0
	reply := &n
	Multiply(5, 7, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 35
	fmt.Println("Multiply:", n)      // Multiply: 35
}
