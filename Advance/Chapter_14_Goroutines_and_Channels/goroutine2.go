package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "tran van quac"
	ch <- "tran van sang"
	ch <- "tran van quyen"
	ch <- "hoang thi thu huong"
	ch <- "tran van tien"
}

func getData(ch chan string) {
	var input string

	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
