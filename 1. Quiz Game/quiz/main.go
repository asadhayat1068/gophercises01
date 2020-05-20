package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer string
}

func main() {
	csvFileName := flag.String(
		"csv", "problems.csv",
		"A csv file in the format of question, answer")

	timeLimit := flag.Int("timelimit", 30, "Time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse the provided CSV file: %s\n", *csvFileName))
	}

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	fmt.Print("Waiting for Channel: ")
	fmt.Print(<-timer.C)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem %d: %s = ", i+1, p.question)
		var answer string
		fmt.Scanf("%s \n", &answer)
		if answer == p.answer {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	problems :=  make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}
	return problems
	
}

func exit(msg string)  {
	fmt.Println(msg)
	os.Exit(1)
}