package day02

import (
	"strconv"
	"strings"

	"github.com/rxedu/adventofcode-2017-go/internal/math"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input)))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input)))
}

func solvePartOne(input [][]int) int {
	checksum := 0
	for _, row := range input {
		a := row[0]
		b := row[0]
		for _, v := range row {
			a = math.Max(a, v)
			b = math.Min(b, v)
		}
		checksum += a - b
	}

	return checksum
}

func solvePartTwo(input [][]int) int {
	checksum := 0
	for _, row := range input {
		checksum += findRatio(row)
	}

	return checksum
}

func findRatio(row []int) int {
	if len(row) < 2 {
		return 0
	}

	for i, n := range row {
		for j, m := range row {
			if i == j {
				continue
			}
			if n%m == 0 {
				return n / m
			}
			if m%n == 0 {
				return m / n
			}
		}
	}

	return 0
}

func parse(input string) [][]int {
	strRows := strings.Split(input, "\n")

	matrix := make([][]int, len(strRows))
	for i, strRow := range strRows {
		strCols := strings.Split(strRow, "\t")

		matrix[i] = make([]int, len(strCols))
		for j, v := range strCols {
			w, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			matrix[i][j] = w
		}
	}

	return matrix
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
