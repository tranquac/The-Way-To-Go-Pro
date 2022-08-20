package main

import "fmt"

func main() {
	Funtion1()
}

func Funtion1() {
	fmt.Printf("In Functional at the top\n")
	defer Funtion2()
	fmt.Printf("In Functional at the end\n")
}

func Funtion2() {
	fmt.Printf("Function2: Defered util the end of the calling function!")
}
