package day02

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

func solvePartOne(input [][]int) int {
	return 0
}

func solvePartTwo(input [][]int) int {
	return 0
}

func parse(input string) [][]int {
	strRows := strings.Split(input, "\n")

	matrix := make([][]int, len(strRows))
	for i, strRow := range strRows {
		strCol := strings.Split(strRow, "\t")

		matrix[i] = make([]int, len(strRow))
		for j, v := range strCol {
			w, err := strconv.Atoi(v)
			matrix[i][j] = w
			if err != nil {
				panic(err)
			}
		}
	}

	return matrix
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
