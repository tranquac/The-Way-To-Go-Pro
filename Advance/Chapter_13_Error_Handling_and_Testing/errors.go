package main

import (
	"errors"
	"fmt"
)

var errNotFound = errors.New("Not found error")

func main() {
	fmt.Printf("error: %v", errNotFound)
}
