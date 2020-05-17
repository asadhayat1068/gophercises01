package main

import "flag"

func main() {
	csvFileName := flag.String(
		"csv", "problems.csv",
		"A csv file in the format of question, answer")
	flag.Parse()
	_ = csvFileName
}