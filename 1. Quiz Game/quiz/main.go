package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer string
}

func main() {
	csvFileName := flag.String(
		"csv", "problems.csv",
		"A csv file in the format of question, answer")
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
	fmt.Println(problems)
}

func parseLines(lines [][]string) []problem {
	problems :=  make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer: line[1],
		}
	}
	return problems
	
}

func exit(msg string)  {
	fmt.Println(msg)
	os.Exit(1)
}