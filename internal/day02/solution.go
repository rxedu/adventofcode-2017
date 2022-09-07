package day02

import (
	"strconv"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input)))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input)))
}

func solvePartOne(input [][]int) int {
	return 0
}

func solvePartTwo(input [][]int) int {
	return 0
}

func parse(input string) [][]int {
	return [][]int{}
}

func serialize(output int) string {
	return strconv.Itoa(output)

}
