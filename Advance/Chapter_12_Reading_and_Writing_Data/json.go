package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type Vcard struct {
	FirstName string
	LastName  string
	Addresses []Address
	Remark    string
}

func main() {
	pa := Address{"private", "Bejing", "HaNoi"}
	wa := Address{"work", "Boom", "Belgium"}
	vc := Vcard{"Jan", "Kersschot", []Address{pa, wa}, "none"}
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
}
