package main

import (
	"fmt"
	"time"
)

type myTime struct {
	t time.Time //anonymous field
}

func (t myTime) first3Chars() string {
	return t.t.String()[0:3]
}

func main() {
	m := myTime{time.Now()}
	fmt.Println("Full time now", m.t.String())
	fmt.Println("First 3 chars", m.first3Chars())

}
