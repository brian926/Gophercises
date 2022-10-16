package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url you want to build out a sitemap for")
	flag.Parse()

	fmt.Println(*urlFlag)
	resp, err := http.Get(*urlFlag)
	if err != nil {

	}
	defer resp.Body.Close()
}
