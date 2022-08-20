package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "The quick brown fox jumped over the lazy dog"
	s1 := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", s1)
	for _, val := range s1 {
		fmt.Printf("%s - ", val)
	}

	fmt.Println()
	str2 := "GO1|The ABC of Go|30"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n",sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)

	// Splitted in slice: [The quick brown fox jumped over the lazy dog]
	// The - quick - brown - fox - jumped - over - the - lazy - dog - 
	// Splitted in slice: [GO1 The ABC of Go 30]
	// GO1 - The ABC of Go - 30 -
	// sl2 joined by ;: GO1;The ABC of Go;30
}
