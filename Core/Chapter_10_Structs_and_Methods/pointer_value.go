package main

import "fmt"

type B struct {
	thing int
}

func (b *B) Change() { b.thing = 1 }

func (b B) write() string { return fmt.Sprint(b) }

func main() {
	var b1 B // b1 is value
	b1.Change()
	fmt.Println(b1.write())

	b2 := new(B) // b2 is pointer
	b2.Change()
	fmt.Println(b2.write())
}
