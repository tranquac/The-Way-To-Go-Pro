package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello, how is it, going, tranvanquac?"
	var manyQ = "qqqqqqqqqqqqqqqqq"

	fmt.Printf("Number of H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyQ)
	fmt.Printf("%d\n", strings.Count(manyQ, "qq"))
}
