package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func operation(records [][]string) (string, error) {
	// TODO: implement this function
	return "...", nil
}

func main() {
	csvFile, _ := filepath.Abs("./pkg/day{{dayNumber}}/input.csv")
	in, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(in)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	answer, _ := operation(records)

	fmt.Printf("current data transform: %s", answer)
}