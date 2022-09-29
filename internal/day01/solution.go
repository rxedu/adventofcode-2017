package day01

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
	if len(input) == 0 {
		return 0
	}

	sum := 0

	for i, cur := range input {
		nextIdx := (i + 1) % len(input)
		next := input[nextIdx]

		if cur == next {
			sum += cur
		}
	}

	return sum
}

func solvePartTwo(input []int) int {
	if len(input) == 0 {
		return 0
	}

	if len(input)%2 != 0 {
		return 0
	}

	sum := 0

	for i, cur := range input {
		nextIdx := (i + len(input)/2) % len(input)
		next := input[nextIdx]

		if cur == next {
			sum += cur
		}
	}

	return sum
}

func parse(input string) []int {
	strSlice := strings.Split(input, "")

	arr := make([]int, len(strSlice))
	for i, v := range strSlice {
		w, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		arr[i] = w
	}
	return arr
}

func serialize(output int) string {
	return strconv.Itoa(output)

}
