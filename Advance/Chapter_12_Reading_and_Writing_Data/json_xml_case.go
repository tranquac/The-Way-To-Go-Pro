package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type thing struct {
	Field1 int
	Field2 string
}

func main() {
	//x := `<x><field1>423</field1><field2>hello from xml</field2></x>`
	j := `{"field1": 423, "field2": "hello from json"}`

	tj := thing{}
	fmt.Printf("From JSON: %#v\n", tj)
	if err := json.Unmarshal([]byte(j), &tj); err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	fmt.Printf("From JSON: %#v\n", tj)
	//fmt.Printf("From XML: %#v\n", tx)
}
