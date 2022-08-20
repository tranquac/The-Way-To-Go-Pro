package main

import "fmt"

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{}

func main() {
	var val Any
	val = 5
	fmt.Printf("val has the value: %v\n", val)
	val = str
	fmt.Printf("val has the value: %v\n", val)
	per1 := new(Person)
	per1.name = "tran van quac"
	per1.age = 22
	val = per1
	fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", *t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}

// interface trong co the duoc gan voi moi kieu du lieu vi no khong khai bao ham nao can cai dat