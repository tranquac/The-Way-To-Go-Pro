// multidim_array.go
package main

import "fmt"

const (
	// WIDTH  = 1920
	// HEIGHT = 1080
	WIDTH =	5
	HEIGHT = 4
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func main() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}
	fmt.Println(screen)

	for row := range screen {
		for column := range screen[0] {
			screen[row][column] = 1
		}
	}

	fmt.Println(screen)
}