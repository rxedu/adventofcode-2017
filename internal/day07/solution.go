package day07

import (
	"math"
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
	links  []string
}

func solvePartOne(input []Node) string {
	if len(input) < 1 {
		return ""
	}

	root := 0
	weight := math.MaxInt
	for i, v := range input {
		if v.weight < weight {
			weight = v.weight
			root = i
		}
	}
	return input[root].name
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
			links:  strings.Split(strings.Split(strRow, " -> ")[1], ", "),
		}
	}

	return matrix
}

func serialize(output string) string {
	return output
}
