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
		fmt.Printf("Failed to open CSV file: %s\n", *csvFileName)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		fmt.Printf("Failed to parse the provided CSV file: %s\n", *csvFileName)
		os.Exit(1)
	}

	fmt.Println(lines)
}