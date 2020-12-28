package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func operation(records []string) (int, error) {
	var groupAnswers []string
	allCounts := map[string]int{}

	rowsRemain := true
	for rowsRemain {
		raw, remaining := identifyRows(records)
		if remaining == nil {
			rowsRemain = false
		} else {
			records = remaining
		}
		if raw != "" {
			groupAnswers = append(groupAnswers, raw)
		}
	}
	for _, v := range groupAnswers {
		allCounts = addCounts(v, allCounts)
	}

	return sumCounts(allCounts), nil
}

func addCounts(answers string, allAnswers map[string]int) map[string]int {
	groupAnswers := strings.Split(answers, " ")

	answered := map[string]int{}
	for _, v := range answers {
		a := string(v)


		if validQuestion(a) {
			answered[a]++
		}
	}

	for v := range answered {
		if answered[v] == len(groupAnswers) {
			allAnswers[v]++
		}
	}
	return allAnswers
}

func validQuestion(question string) bool {
	return question != ""
}

func sumCounts(allAnswers map[string]int) int {
	sum := 0
	for v := range allAnswers {
		sum += allAnswers[v]
	}
	return sum
}

func identifyRows(rows []string) (string, []string) {
	raw := ""
	if rows[0] == "" {
		return raw, rows[1:]
	}

	foundRows := []string{}
	for i, value := range rows {
		v := value

		if v == "" {
			for _, r := range foundRows {
				raw = raw + " " + r
			}
			remainingRows := rows[i:]
			return strings.Trim(raw, " "), remainingRows
		} else {
			foundRows = append(foundRows, v)
		}
	}
	for _, r := range foundRows {
		raw = raw + " " + r
	}

	return strings.Trim(raw, " "), nil
}

func main() {
	inputFile, _ := filepath.Abs("./pkg/day6/input.csv")
	r, _ := ioutil.ReadFile(inputFile)

	records := strings.Split(string(r), "\n")
	answer, _ := operation(records)

	fmt.Printf("current data transform: %d", answer)
}