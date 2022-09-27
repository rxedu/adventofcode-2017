package day11

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

type Loc struct {
	q int
	r int
	s int
}

type Step struct {
	q int
	r int
	s int
}

func solvePartOne(input []Step) int {
	return 0
}

func solvePartTwo(input []Step) int {
	return 0
}

func parse(input string) []Step {
	steps := strings.Split(input, ",")
	arr := make([]Step, len(steps))
	for i, step := range steps {
		arr[i] = parseStep(step)
	}
	return arr
}

func parseStep(step string) Step {
	vectors := map[string]Step{
		"nw": Step{-1, 0, 1},
		"n":  Step{0, -1, 1},
		"ne": Step{1, -1, 0},
		"sw": Step{-1, 1, 0},
		"s":  Step{0, 1, -1},
		"se": Step{1, 0, -1},
	}
	return vectors[step]
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
