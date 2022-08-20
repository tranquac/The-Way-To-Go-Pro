package main

import "fmt"

func trace(s string) {
	fmt.Println("Entering:", s)
}

func untrace(s string) {
	fmt.Println("Leaving:", s)
}

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}

func main() {
	b()
}
