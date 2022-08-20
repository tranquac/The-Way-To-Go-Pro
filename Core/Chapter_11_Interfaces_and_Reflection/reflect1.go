package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.5
	fmt.Println("Type: ", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("Value: ",v)
	fmt.Println("Type: ", v.Type())
	fmt.Println("Kind: ", v.Kind())
	fmt.Println("value: ", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
}
