package day17

import (
	"strconv"

	"github.com/rxedu/adventofcode-2017-go/internal/math"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input)))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input)))
}

func solvePartOne(input int) int {
	steps := 2017
	buffer, cur := spin(steps, input, steps)
	return buffer[cur+1]
}

func solvePartTwo(input int) int {
	steps := 50000000
	buffer, _ := spin(steps, input, 3)
	return buffer[1]
}

func spin(steps int, jumps int, maxSize int) ([]int, int) {
	buffer := make([]int, maxSize)

	cur := 0
	for i := 1; i <= steps; i++ {
		size := i
		idx := ((cur + jumps) % size) + 1
		n := math.Min(idx, maxSize-1)
		copy(buffer[n:], buffer[n-1:])
		buffer[n] = i
		cur = idx
	}
	return buffer, cur
}

func parse(input string) int {
	v, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return v
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
