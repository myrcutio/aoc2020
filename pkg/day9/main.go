package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func operation(records []string, preambleSize int) int {
	nums := make([]int, len(records))
	for i, v := range records {
		val, _ := strconv.Atoi(v)
		nums[i] = val
	}

	invalidNum := 0
	marker := preambleSize
	for _, v := range nums[preambleSize:] {
		if !isValidNumber(nums[marker-preambleSize: marker], v) {
			invalidNum = v
			break
		}
		marker++
	}

	cSet := findContiguousSet(nums, invalidNum)

	s := invalidNum
	l := 0
	for _, v := range cSet {
		if v < s {
			s = v
		}
		if v > l {
			l = v
		}
	}

	return s + l
}

func isValidNumber(prevNums []int, num int) bool {
	searchMap := map[int]bool{}

	for _, v := range prevNums {
		searchMap[num - v] = true
	}

	for _, v := range prevNums {
		if searchMap[v] {
			return true
		}
	}

	return false
}

func findContiguousSet(nums []int, num int) []int {
	currentSum := nums[0] + nums[1]
	s := 0
	e := 1
	for true {
		if currentSum == num {
			return nums[s:e]
		}
		if currentSum < num {
			e++
			currentSum += nums[e]
		}

		if currentSum == num {
			return nums[s:e+1]
		}
		if currentSum > num {
			currentSum -= nums[s]
			s++
		}
	}

	panic("could not find set")
}

func main() {
	inputFile, _ := filepath.Abs("./pkg/day9/input.csv")
	r, _ := ioutil.ReadFile(inputFile)

	records := strings.Split(string(r), "\n")
	answer := operation(records, 25)

	fmt.Printf("current data transform: %d", answer)
}