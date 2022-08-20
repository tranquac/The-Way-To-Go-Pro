package main

import "fmt"

func main() {
	s := "tran van quac"
	for i, c := range s {
		fmt.Printf("%d:%c ", i, c)
	}
}

// prints: 0:ÿ 2:界