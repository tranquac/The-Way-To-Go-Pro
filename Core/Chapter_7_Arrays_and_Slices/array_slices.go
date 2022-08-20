package main

import "fmt"

func main() {
	var arr1 [10]int
	var slice1 []int = arr1[2:5]

	//load the array with integer: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	//print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at index %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1)) //len(arr1) - (len(arr1) - len(slice1))

	//grow the slice
	slice1 = slice1[5:7]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}
