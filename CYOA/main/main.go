package main

import (
	"flag"
	"fmt"
	"os"

	cyoa ".."
)

func main() {
	jsonFileName := flag.String("json", "../gopher.json", "JSON of the actual story")
	flag.Parse()

	jsonFile, err := os.Open(*jsonFileName)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(jsonFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)
}
