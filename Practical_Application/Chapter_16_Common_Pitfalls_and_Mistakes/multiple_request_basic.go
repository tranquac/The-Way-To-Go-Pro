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

func MakeRequest(url string) {
	start := time.Now()
	resp, _ := http.Get(url)
	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%.2f elapsed with response length: %d %s\n", secs, len(body), url)
}
func main() {
	start := time.Now()
	fileName := os.Args[1]
	urls := readUrlsFromFile(fileName)
	for _, url := range urls {
		MakeRequest(url)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
