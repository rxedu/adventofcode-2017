package day13

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

func solvePartOne(input map[int]int) int {
	maxDepth := getMaxDepth(input)

	severity := 0
	for i := 0; i <= maxDepth; i++ {
		r := input[i]
		if isScannerAtTop(r, i) {
			severity += i * r
		}
	}

	return severity
}

func solvePartTwo(input map[int]int) int {
	maxDelay := 100000000
	for delay := 0; delay < maxDelay; delay++ {
		if willGetThrough(input, delay) {
			return delay
		}
	}
	return 0
}

func willGetThrough(input map[int]int, delay int) bool {
	maxDepth := getMaxDepth(input)
	for i := 0; i <= maxDepth; i++ {
		r := input[i]
		t := i + delay
		if isScannerAtTop(r, t) {
			return false
		}
	}
	return true
}

func isScannerAtTop(r int, t int) bool {
	if r == 0 {
		return false
	}
	period := 2 * (r - 1)
	return t%period == 0
}

func getMaxDepth(input map[int]int) int {
	var maxDepth int
	for idx := range input {
		maxDepth = math.Max(idx, maxDepth)
	}
	return maxDepth
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
