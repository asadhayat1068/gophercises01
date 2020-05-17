package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

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

	fmt.Println(lines)
}

func exit(msg string)  {
	fmt.Println(msg)
	os.Exit(1)
}