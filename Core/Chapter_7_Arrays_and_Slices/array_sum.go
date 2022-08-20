package main

import "fmt"

func main() {
	array := [3]float64{7.3, 8.2, 1.4}
	x := Sum(&array)
	fmt.Printf("The sum of the array is: %f", x)
}

func Sum(s *[3]float64) float64 {
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	return sum
}

// func main() {
// 	array := [3]float64{7.3, 8.2, 1.4}
// 	x := Sum(array)
// 	fmt.Printf("The sum of the array is: %f", x)
// }

// func Sum(s [3]float64) float64 {
// 	sum := 0.0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	return sum
// }

// The sum of the array is: 16.900000
