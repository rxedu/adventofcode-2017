package day14

import (
	"strconv"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(input))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(input))
}

func solvePartOne(input string) int {
	return 0
}

func solvePartTwo(input string) int {
	return 0
}

func stringToBytes(input string) []int {
	arr := make([]int, len(input))
	for i, str := range input {
		arr[i] = int(str)
	}
	return arr
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
