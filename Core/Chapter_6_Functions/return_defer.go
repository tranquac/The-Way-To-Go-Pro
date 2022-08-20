package main

import "fmt"

func main() {
	fmt.Println(f())
}

func f() (ret string) {
	defer func() {
		ret += ", tranvanquac"
	}()
	return "hello"
}

//2 -> defer run after return
