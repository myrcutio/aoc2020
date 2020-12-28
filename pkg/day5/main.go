package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func operation(records []string) (int, error) {
	seatIds := []int{}

	for _, seat := range records {
		row := findRow(seat[:7])
		column := findColumn(seat[7:])
		seatID := convertSeatToID(row, column)
		seatIds = append(seatIds, seatID)
	}

	sort.Ints(seatIds)
	for i := seatIds[0]; i < len(seatIds); i++ {
		if seatIds[i]+1 != seatIds[i+1] {
			return seatIds[i]+1, nil
		}
	}

	return 0, errors.New("no seat ID gap found")
}

func findRow(row string) int {
	return stringToBin(row, "F", "B")
}

func findColumn(column string) int {
	return stringToBin(column, "L", "R")
}

func stringToBin(s, zero, one string) int {
	binString := ""
	for _, value := range s {
		v := string(value)
		if v == zero {
			binString += "0"
		} else if v == one {
			binString += "1"
		}
	}
	out, _ := strconv.ParseInt(binString, 2, 64)
	return int(out)
}

func convertSeatToID(row, column int) int {
	return row * 8 + column
}

func main() {
	inputFile, _ := filepath.Abs("./pkg/day5/input.csv")
	r, _ := ioutil.ReadFile(inputFile)

	records := strings.Split(string(r), "\n")
	answer, _ := operation(records)

	fmt.Printf("current data transform: %d", answer)
}