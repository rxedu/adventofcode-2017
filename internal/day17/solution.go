package day17

import (
	"strconv"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input)))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input)))
}

func solvePartOne(input int) int {
	return 0
}

func solvePartTwo(input int) int {
	return 0
}

func parse(input string) int {
	v, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return v
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
