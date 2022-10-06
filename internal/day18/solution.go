package day18

import (
	"strconv"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input)))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input)))
}

type Instruction interface {
}

func solvePartOne(input []Instruction) int {
	return 0
}

func solvePartTwo(input []Instruction) int {
	return 0
}

func parse(input string) []Instruction {
	return []Instruction{}
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
