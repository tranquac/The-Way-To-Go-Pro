package main

import (
	"fmt"
	"strconv"
)

type struct1 struct {
	name, job string
	age       int
}

func (s struct1) String() string {
	return "Name:" + s.name + ", age:" + strconv.Itoa(s.age) + ", job " + s.job
}

func main() {
	per1 := new(struct1)
	per2 := struct1{"tran van quac", "IT", 22}
	fmt.Println(per1)
	fmt.Println(per2)
}
