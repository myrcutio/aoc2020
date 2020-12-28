package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

type instruction struct {
	op string
	value int
}

func operation(records []string) int {
	ins := make([]instruction, len(records))

	for i, r := range records {
		ins[i] = parseOperation(r)
	}

	accumulator, corruptedIndex := findCorruptedInstruction(ins)
	fmt.Printf("corrupted index: %d\n", corruptedIndex)

	return accumulator
}

// return index of corrupted instruction
func findCorruptedInstruction(ins []instruction) (int, int) {
	for i, v := range ins {
		if !(v.op == "acc") {
			flipInstruction(&ins[i])

			a, finalIndex := runInstruction(ins)
			if finalIndex == len(ins) {
				return a, finalIndex
			} else {
				flipInstruction(&ins[i])
			}
		}
	}
	panic("no corrupted instruction found")
}

func flipInstruction(in *instruction) {
	var (
		nop = "nop"
		jmp = "jmp"
	)

	if in.op == nop {
		in.op = jmp
	} else if in.op == jmp {
		in.op = nop
	}
}

func runInstruction(ins []instruction) (int, int) {
	instructionsRun := map[int]bool{}
	accumulator := 0
	i := 0
	for true {
		if i == len(ins) {
			return accumulator, i
		}
		switch ins[i].op {
		case "nop":
			i++
		case "acc":
			accumulator += ins[i].value
			i++
		case "jmp":
			i += ins[i].value
		}

		if instructionsRun[i] {
			break
		}
		instructionsRun[i] = true
	}
	return accumulator, i
}

func parseOperation(op string) instruction {
	seg := strings.Split(op, " ")
	v, _ := strconv.Atoi(seg[1][1:])
	if strings.HasPrefix(seg[1], "-") {
		v = -v
	}
	return instruction{
		op:    seg[0],
		value: v,
	}
}

func main() {
	inputFile, _ := filepath.Abs("./pkg/day8/input.csv")
	r, _ := ioutil.ReadFile(inputFile)

	records := strings.Split(string(r), "\n")
	answer := operation(records)

	fmt.Printf("current data transform: %d\n", answer)
}