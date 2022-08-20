package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hi, I'm tranvanquac, Hi"

	fmt.Printf("The position of \"tranvanquac\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "tranvanquac"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	// The position of "tranvanquac" is: 8
	// The position of the first instance of "Hi" is: 0
	// The position of the last instance of "Hi" is: 21
}
