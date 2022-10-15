package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url you want to build out a sitemap for")
	flag.Parse()

	fmt.Println(*urlFlag)
}
