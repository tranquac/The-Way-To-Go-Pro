package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "this is an example of a string"

	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "th"))
}
