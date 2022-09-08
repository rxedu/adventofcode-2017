package day05

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

func solvePartOne(input []int) int {
	count := 0

	if len(input) == 0 {
		return count
	}

	for i := 0; i < len(input); count++ {
		j := input[i]
		input[i]++
		i = i + j
	}

	return count
}

func solvePartTwo(input []int) int {
	if len(input) == 0 {
		return 0
	}
	return 0
}

func parse(input string) []int {
	rows := strings.Split(input, "\n")
	arr := make([]int, len(rows))
	for i, str := range rows {
		v, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		arr[i] = v
	}
	return arr
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
