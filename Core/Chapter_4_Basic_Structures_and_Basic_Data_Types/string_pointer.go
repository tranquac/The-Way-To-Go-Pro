package main

import "fmt"

func main() {
	s := "say hi from VietNam"
	var p *string = &s
	*p = "lalala, love Ha Noi"

	fmt.Printf("Here is the pointer p: %p\n", p)
	fmt.Printf("Here is the string *p: %s\n", *p)
	fmt.Printf("Here is the string s: %s",s)

	// Here is the pointer p: 0xc000088220
	// Here is the string *p: lalala, love Ha Noi
	// Here is the string s: lalala, love Ha Noi
}
