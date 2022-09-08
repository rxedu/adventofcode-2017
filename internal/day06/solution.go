package day06

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
	return 0
}

func solvePartTwo(input []int) int {
	return 0
}

func parse(input string) []int {
	cols := strings.Split(input, "\t")
	arr := make([]int, len(cols))
	for i, str := range cols {
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
