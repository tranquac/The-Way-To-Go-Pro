package main

import "fmt"

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 234, "delta": 12, "echo": 623, "juliet": 65, "kilo": 43, "lima": 98}
)

//string -> key; int -> value

func main() {
	invMap := make(map[int]string, len(barVal))
	fmt.Println("Before invert:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
		invMap[v] = k
	}
	println()
	fmt.Println("After invert:")
	for k, v := range invMap {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
}

// Before invert:
// Key: juliet, Value: 65 / Key: kilo, Value: 43 / Key: lima, Value: 98 / Key: alpha, Value: 34 / Key: bravo, Value: 56 / Key: charlie, Value: 234 / Key: delta, Value: 12 / Key: echo, Value: 623 / 
// After invert:
// Key: 65, Value: juliet / Key: 43, Value: kilo / Key: 98, Value: lima / Key: 34, Value: alpha / Key: 56, Value: bravo / Key: 234, Value: charlie / Key: 12, Value: delta / Key: 623, Value: echo /