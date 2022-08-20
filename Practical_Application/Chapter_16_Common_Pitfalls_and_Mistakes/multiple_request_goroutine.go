package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func readUrlsFromFile(filename string) []string {
	url_file, err := os.Open(filename)
	if err != nil {
		log.Fatal("File could not be read\n")
	}
	defer url_file.Close()
	uScanner := bufio.NewScanner(url_file)
	var urls []string
	for uScanner.Scan() {
		urls = append(urls, uScanner.Text())
	}
	if err := uScanner.Err(); err != nil {
		log.Fatal(err)
	}
	return urls
}

func MakeRequest(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err == nil {
		secs := time.Since(start).Seconds()
		body, _ := ioutil.ReadAll(resp.Body)
		ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
	} else {
		ch <- fmt.Sprintf("%v", err)
	}
}
func main() {
	start := time.Now()
	ch := make(chan string)
	fileName := os.Args[1]
	urls := readUrlsFromFile(fileName)
	for _, url := range urls {
		go MakeRequest(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
