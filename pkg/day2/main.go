package main

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func isValidPassword(policy, password string) bool {
	policyVals := strings.Split(policy, " ")
	valueRange := strings.Split(policyVals[0], "-")
	minRange, _ := strconv.Atoi(valueRange[0])
	maxRange, _ := strconv.Atoi(valueRange[1])


	if meetsCharacterRangeReq(minRange, maxRange, policyVals[1], password) {
		return true
	}

	return false
}

func meetsCharacterRangeReq(min, max int, character, password string) bool {
	i := strings.Count(password, character)

	return i >= min && i <= max
}

func iteratePasswords(records [][]string) int {
	validPasswordCount := 0
	for _, record := range records {
		r := record[0]
		passRecord := strings.Split(r, ": ")
		if isValidPassword(passRecord[0], passRecord[1]) {
			validPasswordCount ++
		}
	}
	return validPasswordCount
}

func isNewValidPassword(policy, password string) bool {
	policyVals := strings.Split(policy, " ")
	valueRange := strings.Split(policyVals[0], "-")
	minRange, _ := strconv.Atoi(valueRange[0])
	maxRange, _ := strconv.Atoi(valueRange[1])

	if meetsCharacterPositionReq(minRange, maxRange, policyVals[1], password) {
		return true
	}

	return false
}

func meetsCharacterPositionReq(first, second int, character, password string) bool {
	char1 := string(password[first-1])
	char2 := string(password[second-1])
	return char1 != char2 && (char1 == character || char2 == character)
}

func newIteratePasswords(records [][]string) int {
	validPasswordCount := 0
	for _, record := range records {
		r := record[0]
		passRecord := strings.Split(r, ": ")
		if isNewValidPassword(passRecord[0], passRecord[1]) {
			validPasswordCount ++
		}
	}
	return validPasswordCount
}

func main() {
	csvFile, _ := filepath.Abs("./pkg/day2/input.csv")
	in, err := os.Open(csvFile)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(in)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("valid password count: %d", iteratePasswords(records))
	log.Printf("new valid password count: %d", newIteratePasswords(records))
}