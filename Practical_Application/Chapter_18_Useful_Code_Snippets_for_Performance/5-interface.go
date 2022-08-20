package main

import "fmt"

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("param #%d is a bool\n", i)
		case float64:
			fmt.Printf("param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("param #%d is an int\n", i)
		case nil:
			fmt.Printf("param #%d is nil\n", i)
		case string:
			fmt.Printf("param #%d is a string\n", i)
		default:
			fmt.Printf("param #%d’s type is unknown\n", i)
		}
	}
}

func main() {
	// (1) How to detect vwhether implements the interface Stringer
	// if v, ok := v.(Stringer); ok {
	// 	fmt.Printf("implements String(): %s\n", v.String())
	// }
		classifier("tran van quac", true)
	// (2) How to use an interface to implement a type classification function:
}
