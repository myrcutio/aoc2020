package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type group struct {
	Data string
}

func operation(records []string) (int, error) {
	// TODO: implement this function
	logicalGroups := groupRows(records)

	fmt.Printf("%+v", logicalGroups)
	return 0, nil
}

func groupRows(rows []string) []group {
	var rowGroup []group
	rowsRemain := true
	for rowsRemain {
		raw, remaining := identifyRows(rows)
		if remaining == nil {
			rowsRemain = false
		} else {
			rows = remaining
		}
		if raw != "" {
			rowGroup = append(rowGroup, parseGroup(raw))
		}
	}
	return rowGroup
}

func parseGroup(s string) group {
	return group{
		Data: s,
	}
}

func identifyRows(rows []string) (string, []string) {
	raw := ""
	if rows[0] == "" {
		return raw, rows[1:]
	}

	var foundRows []string
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
	inputFile, _ := filepath.Abs("./pkg/day{{dayNumber}}/input.csv")
	r, _ := ioutil.ReadFile(inputFile)

	records := strings.Split(string(r), "\n")
	answer, _ := operation(records)

	fmt.Printf("current data transform: %d", answer)
}