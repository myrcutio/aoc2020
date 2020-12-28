package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type Path struct {
	Indices []int
	string string
}

func operation(records []string) int {
	adapters := make([]int, len(records))

	for i, r := range records {
		v, _ := strconv.ParseInt(r, 10, 16)
		adapters[i] = int(v)
	}
	sort.Ints(adapters)

	joltDiffs := countJoltDiffs(adapters)
	fmt.Printf("diff distribution: %+v\n", joltDiffs)

	permutations := countPermutations(adapters)

	return permutations
}

func countJoltDiffs(adapters []int) map[int]int {
	counted := map[int]int{}
	for i, v := range adapters {
		if i == 0 {
			counted[v]++
			continue
		}

		if v == adapters[i - 1] {
			panic(fmt.Sprintf("found duplicate: %d\n", v))
		}
		counted[v - adapters[i - 1]]++
	}
	counted[3]++
	return counted
}

func countPermutations(adapters []int) int {
	shortPath := findShortestPath(adapters)

	var f big.Int
	f.MulRange(1, 50)
	fmt.Println(&f)

	additionalPerms := 1

	for i, _ := range shortPath.Indices {
		if i == 0 {
			continue
		}
		additionalPerms *= shortPath.Indices[i] - shortPath.Indices[i - 1]
	}

	return additionalPerms
}

// returns an array of indices
func findShortestPath(adapters []int) Path {
	shortPath := Path{}
	shortPath.Indices = []int{}

	lastMarker := 0
	currentDiff := adapters[0]

	for i, v := range adapters {
		currentDiff = v - adapters[lastMarker]

		if i == 0 {
			shortPath.Indices = append(shortPath.Indices, i)
			lastMarker = i
		}
		if currentDiff > 3 {
			shortPath.Indices = append(shortPath.Indices, i-1)
			lastMarker = i-1
		}

	}
	shortPath.Indices = append(shortPath.Indices, len(adapters) - 1)

	return shortPath
}

func (p *Path) String() string {
	s := ""
	for _, v := range p.Indices {
		s += strconv.Itoa(v)
	}
	return s
}

func main() {
	inputFile, _ := filepath.Abs("./pkg/day10/input.csv")
	r, _ := ioutil.ReadFile(inputFile)

	records := strings.Split(string(r), "\n")
	answer := operation(records)

	fmt.Printf("current data transform: %d", answer)
}