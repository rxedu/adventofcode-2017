package day09

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

func solvePartOne(input []string) int {
	return 0
}

func solvePartTwo(input []string) int {
	return 0
}

func parse(input string) []string {
	return strings.Split(input, "\n")
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
