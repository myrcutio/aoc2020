package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

func operation(records []string) int {
	colorTree := map[string]map[string]int{}
	for _, v := range records {
		newRule := findContainerRules(v)
		for r := range newRule {
			colorTree[r] = newRule[r]
		}
	}

	fmt.Printf("shiny gold must contain %d bags\n", calculateContainedBags("shiny gold", colorTree))

	return countValidContainers("shiny gold", colorTree)
}

func calculateContainedBags(color string, colorTree map[string]map[string]int) int {
	bagSum := 0
	for k := range colorTree[color] {
		c := colorTree[color][k]
		bagSum += c * calculateContainedBags(k, colorTree) + c
	}
	return bagSum
}

func findContainerRules(rule string) map[string]map[string]int {
	ruleSegments := strings.Split(rule, " bags contain ")
	containerBagType := ruleSegments[0]

	if strings.HasPrefix(ruleSegments[1], "no") {
		return map[string]map[string]int{
			containerBagType: {},
		}
	}
	containsRules := strings.Split(ruleSegments[1], ",")

	childRules := map[string]int{}
	for _, v := range containsRules {
		r := findBagCount(v)
		for color := range r {
			childRules[color] = r[color]
		}
	}
	return map[string]map[string]int{
		containerBagType: childRules,
	}
}

func findBagCount(rule string) map[string]int {
	segs := strings.Split(rule, " bag")
	numS := strings.Split(strings.Trim(segs[0], " "), " ")
	num, _ := strconv.Atoi(numS[0])
	bagColor := strings.Join(numS[1:], " ")

	return map[string]int{
		bagColor: num,
	}
}

func countValidContainers(color string, colorTree map[string]map[string]int) int {
	totalCount := 0

	for k := range colorTree {
		canContain := false
		if k == color {
			continue
		}
		if nodeContainsColor(color, k, colorTree) {
			canContain = true
		}
		if canContain {
			totalCount++
		}
	}

	return totalCount
}

func nodeContainsColor(matchColor, color string, colorTree map[string]map[string]int) bool {
	if colorTree[color][matchColor] > 0 {
		return true
	}
	if len(colorTree[color]) > 0 {
		childContainsColor := false
		for k := range colorTree[color] {
			if nodeContainsColor(matchColor, k, colorTree) {
				childContainsColor = true
			}
		}
		return childContainsColor
	}
	return false
}

func main() {
	inputFile, _ := filepath.Abs("./pkg/day7/input.csv")
	r, _ := ioutil.ReadFile(inputFile)

	records := strings.Split(string(r), "\n")
	answer := operation(records)

	fmt.Printf("current data transform: %d", answer)
}