package main

import "fmt"

func main() {
	// create the new
	map1 := make(map[string]int)
	fmt.Println(map1)
	//init
	map2 := map[string]int{"one": 1, "two": 2, "three": 3}

	//traverse
	for key, value := range map2 {
		fmt.Println("Key: ", key, "Value: ", value)
	}

	// detect key1 existance
	//val1, isPresent = map2["one"]
	delete(map2, "one") //delete
}
