package main

import "fmt"

func main() {
	var arrAge = [5]int{18, 43, 33, 64, 12}
	var arrLazy = [...]int{1, 2, 3, 4, 5}
	//var arrLazy2 = []int{5, 6, 7, 8, 22}
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}

	for i := 0; i < len(arrAge); i++ {
		fmt.Printf("Age at %d is %d\n", i, arrAge[i])
	}

	for i := 0; i < len(arrLazy); i++ {
		fmt.Printf("Number at %d is %d\n", i, arrLazy[i])
	}
	fmt.Printf("\n")

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}
}
