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

type Countdown struct {
	t int
	d int
	h int
	m int
	s int
}

func main() {
	csvFileName := flag.String("csv", "quiz.csv", "csv file in format of 'question,answer'")
	flag.Parse()

	deadline := flag.String("deadline", "", "Deadline for the countdown time")
	flag.Parse()

	v, err := time.Parse(time.RFC3339, *deadline)
	if err != nil {
		exit("Failed to parse Timer")
	}

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
	result := runQuiz(quizList)

	for range time.Tick(1 * time.Second) {
		timeRemaining := getTimeRemaining(v)

		if timeRemaining.t <= 0 {
			exit("End of timer")
		}
	}

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

func getTimeRemaining(t time.Time) Countdown {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return Countdown{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}
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
