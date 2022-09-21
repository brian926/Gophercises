package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type Quiz struct {
	Question string
	Answer   string
}

func main() {
	csvFileName := flag.String("csv", "quiz.csv", "csv file in format of 'question,answer'")
	timeLimit := flag.Int("limit", 90, "The time limit for the quiz in seconds")
	flag.Parse()

	csvFile, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	r := csv.NewReader(csvFile)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file!")
	}

	quizList := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	<-timer.C

	result := runQuiz(quizList)

	fmt.Printf("%v Corret out of %v", result, len(quizList))
}

func runQuiz(quizList []Quiz) int {
	var results = 0

	for i, s := range quizList {
		fmt.Printf("Question %v: %v\n", i+1, s.Question)

		var ans string
		fmt.Scanln(&ans)
		if ans == s.Answer {
			results++
		}
	}

	return results
}

func parseLines(lines [][]string) []Quiz {
	rec := make([]Quiz, len(lines))

	for i, line := range lines {
		rec[i] = Quiz{

			Question: line[0],
			Answer:   line[1],
		}
	}

	return rec
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
