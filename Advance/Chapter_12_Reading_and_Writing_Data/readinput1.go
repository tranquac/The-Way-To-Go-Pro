package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.23 / 2343 / GO"
	format                 = " %f / %d / %d"
)

func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// // fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s - %s!\n", firstName, lastName)

	fmt.Sscan(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s) // From the string we read:  56.12 5212 Go
}
