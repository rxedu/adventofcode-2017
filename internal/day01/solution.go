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

	var sum int

	for i, cur := range input {
		var next int
		if i+1 == len(input) {
			next = input[0]
		} else {
			next = input[i+1]
		}

		if cur == next {
			sum += cur
		}

	}

	return sum
}

func solvePartTwo(input []int) int {
	return 0
}

func parse(input string) []int {
	strSlice := strings.Split(input, "")

	arr := make([]int, len(strSlice))
	for i, v := range strSlice {
		w, err := strconv.Atoi(v)
		arr[i] = w
		if err != nil {
			panic(err)
		}
	}
	return arr
}

func serialize(output int) string {
	return strconv.Itoa(output)

}
