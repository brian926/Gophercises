package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	cyoa ".."
)

func main() {
	port := flag.Int("port", 3000, "the port to start the web app on")
	jsonFileName := flag.String("json", "gopher.json", "JSON of the actual story")
	flag.Parse()

	jsonFile, err := os.Open(*jsonFileName)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(jsonFile)
	if err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
