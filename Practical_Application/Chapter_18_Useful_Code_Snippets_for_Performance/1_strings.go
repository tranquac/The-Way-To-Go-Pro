package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// (1) How to modify a character in a string:
	str := "hello"
	c := []byte(str)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(s2)

	// (2) How  to get a substring in a string
	substr := str[2:5]
	fmt.Println(substr)

	// (3) How to using for and for-range traverse a string
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i]) // gives only the bytes:
	}
	
	for _, ch := range str {
		fmt.Println(ch)
	}

	// (4) How to get the number of bytes in a string: (len(str))
	// fastest: utf8.RuneCountInString(str)
	fmt.Println("The length of string:", utf8.RuneCountInString(str))

	// (5) How to concat
	str1 := "van "
	str2 := "quac"
	str1 += str2
	fmt.Println(str1)
}
