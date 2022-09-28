package day13

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

func solvePartOne(input map[int]int) int {
	return 0
}

func solvePartTwo(input map[int]int) int {
	return 0
}

func parse(input string) map[int]int {
	rows := strings.Split(input, "\n")
	layers := make(map[int]int, len(rows))
	for _, link := range rows {
		parts := strings.Split(link, ": ")

		idx, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		val, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		layers[idx] = val
	}

	return layers
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
