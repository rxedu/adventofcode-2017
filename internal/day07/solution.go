package day07

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

type Node struct {
	name   string
	weight int
}

func solvePartOne(input []Node) string {
	if len(input) < 1 {
		return ""
	}
	return input[0].name
}

func solvePartTwo(input []Node) string {
	return ""
}

func parse(input string) []Node {
	strRows := strings.Split(input, "\n")

	matrix := make([]Node, len(strRows))
	for i, strRow := range strRows {
		parts := strings.Split(strRow, " ")
		weight, err := strconv.Atoi(
			strings.TrimSuffix(strings.TrimPrefix(parts[1], "("), ")"))
		if err != nil {
			panic(err)
		}
		matrix[i] = Node{
			name:   parts[0],
			weight: weight,
		}
	}

	return matrix
}

func serialize(output string) string {
	return output
}
