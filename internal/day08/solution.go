package day08

import (
	"strconv"
	"strings"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input)))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input)))
}

type Instruction struct {
	reg       string
	change    int
	condition Condition
}

type Condition struct {
	reg string
	op  string
	val int
}

func solvePartOne(input []Instruction) int {
	return 0
}

func solvePartTwo(input []Instruction) int {
	return 0
}

func parse(input string) []Instruction {
	strRows := strings.Split(input, "\n")

	arr := make([]Instruction, len(strRows))
	for i, strRow := range strRows {
		parts := strings.Split(strRow, " ")

		change, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}

		val, err := strconv.Atoi(parts[6])
		if err != nil {
			panic(err)
		}

		arr[i] = Instruction{
			reg:    parts[0],
			change: change,
			condition: Condition{
				reg: parts[4],
				op:  parts[5],
				val: val,
			},
		}
	}

	return arr
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
