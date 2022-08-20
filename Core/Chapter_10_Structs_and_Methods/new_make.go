package main

import "fmt"

type Foo map[string]string

type Bar struct {
	thingOne string
	thingTwo int
}

func main() {
	// OK:
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1
	fmt.Println(y)
	// not OK:
	// z := make(Bar) //cannot make Bar; type must be slice, map, or channel
	// z.thingOne = "hello"
	// z.thingTwo = 1

	// OK:
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"
	// not OK:
	// u := new(Foo)
	// (*u)["x"] = "goodbye" // !! panic !!: runtime error: assignment to entry in nil map
	// (*u)["y"] = "world"
}
