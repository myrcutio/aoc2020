package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func operation(records [][]string, defaultSlope []int) (int, error) {
	grid := transformCSVRecords(records)

	stepsToComplete := len(grid) / defaultSlope[1]
	treesHit := 0

	for i := 0; i < stepsToComplete; i++ {
		if isATree(grid, i*defaultSlope[0], i + defaultSlope[1]-1) {
			treesHit ++
		}
	}

	return treesHit, nil
}

func transformCSVRecords(records [][]string) (grid []string) {
	for _, row := range records {
		r := row
		grid = append(grid, r[0])
	}
	return grid
}

func isATree(grid []string, x, y int) bool {
	if x >= len(grid[0]) {
		x = x % len(grid[0])
	}
	return string(grid[y][x]) == "#"
}

func main() {
	csvFile, _ := filepath.Abs("./pkg/day3/input.csv")
	in, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(in)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// x-traverse, y-shift
	slopes := [][]int{
		{1,1},
		{3,1},
		{5,1},
		{7,1},
		{1,2},
	}
	answer := 0
	for _, v := range slopes {
		slopeAnswer, _ := operation(records, v)
		fmt.Printf("slope %+v: %d", v, slopeAnswer)
		if answer == 0 {
			answer = slopeAnswer
		} else {
			answer *= slopeAnswer
		}
	}

	fmt.Printf("current data transform: %d\n", answer)
}