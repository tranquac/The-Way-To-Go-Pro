package main

import "fmt"

func main() {
	// create
	// arr1 := new([5]int) // array
	// slice1 := make([]int, 10)

	//initlization
	arr2 := [...]int{1, 2, 3, 4, 5}
	arrKeyValue := [5]int{1: 1, 2: 2, 3: 3}
	var slice2 []int = arr2[2:5]
	fmt.Println(arrKeyValue)

	// truncate the lest element of an array or slice
	line := slice2[:len(slice2)-1]
	fmt.Println(line)
	// how to use for or for-range traverse an array (or slice)
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	for _, value := range arr2 {
		fmt.Println(value)
	}
}