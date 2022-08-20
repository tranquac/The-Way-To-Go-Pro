package main

import "fmt"

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i := 0; i < 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {

}

func (b *Bird) Quack() {
	fmt.Println("I am quacking!")
}

func (b *Bird) Walk() {
	fmt.Println("I am walking!")
}

func main() {
	b := new(Bird)
	DuckDance(b)
}

// interface la giao dien khai bao cac function de struct cai dat cac fuction duoc khai bao do