package day16

import (
	"strings"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input), createDancers()))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input), createDancers()))
}

type Step interface{}

func solvePartOne(input []Step, dancers []string) []string {
	return dancers
}

func solvePartTwo(input []Step, dancers []string) []string {
	return dancers
}

func parse(input string) []Step {
	return []Step{}
}

func serialize(output []string) string {
	return strings.Join(output, "")
}

func createDancers() []string {
	return []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"o",
		"p",
	}
}
