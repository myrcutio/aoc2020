package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func findPairsForSum(values []int, targetSum int) (int, int, error) {
	matchValues := map[int]bool{}

	for _, value := range values {
		v := value
		match := targetSum - v
		if matchValues[v] {
			return v, match, nil
		} else {
			matchValues[match] = true
		}
	}
	return 0,0, errors.New("could not find a match")
}

func findTripletsForSum(values []int, targetSum int) (int, int, int, error) {
	for i, value := range values {
		v := value
		diff := targetSum - v
		p1, p2, err := findPairsForSum(values[i+1:], diff)
		if err == nil {
			return v, p1, p2, nil
		}
	}
	return 0, 0, 0, errors.New("could not find a match")
}

func main() {
	csvFile, _ := filepath.Abs("./pkg/day1/input.csv")
	in, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(in)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var values []int
	for _, record := range records {
		r, _ := strconv.Atoi(record[0])
		values = append(values, r)
	}

	p1, p2, err := findPairsForSum(values, 2020)
	if err != nil {
		fmt.Print("couldn't find a pair match")
	} else {
		fmt.Printf("pair results of %d and %d is %+v", p1, p2, p1 * p2)
	}

	t1, t2, t3, err := findTripletsForSum(values, 2020)
	if err != nil {
		fmt.Print("\ncouldn't find a pair match")
	} else {
		fmt.Printf("\npair results of %d, %d, and %d is %+v", t1, t2, t3, t1 * t2 * t3)
	}
}