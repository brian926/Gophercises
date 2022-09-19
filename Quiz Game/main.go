package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Quiz struct {
	Question string
	Answer   string
}

func main() {
	csvFile, err := os.Open("quiz.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(csvFile)
	var quizList []Quiz

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		var rec Quiz
		rec.Question = record[0]
		rec.Answer = record[1]

		quizList = append(quizList, rec)
	}

	var result = 0

	for i, s := range quizList {
		fmt.Printf("Question %v: %v\n", i+1, s.Question)

		var ans string
		fmt.Scanln(&ans)
		if ans == s.Answer {
			result += 1
		}
	}

	fmt.Printf("%v Corret out of %v", result, len(quizList))
}
