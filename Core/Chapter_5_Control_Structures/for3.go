package main

import "fmt"

func main() {
	var i int = 5

	for {
		i = i - 1
		fmt.Printf("The value is %d\n", i)
		if i < 0 {
			break
		}
	}
}
