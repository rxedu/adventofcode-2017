package day17

import (
	"strconv"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input)))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input)))
}

func solvePartOne(input int) int {
	steps := 2017
	buffer, cur := spin(steps, input)
	return buffer[cur+1]
}

func solvePartTwo(input int) int {
	steps := 50000000
	buffer, _ := spin(steps, input)
	return buffer[1]
}

func spin(steps int, jumps int) ([]int, int) {
	buffer := []int{0}

	cur := 0
	for i := 1; i <= steps; i++ {
		size := len(buffer)
		idx := ((cur + jumps) % size) + 1
		buffer = append(buffer, 0)
		copy(buffer[idx:], buffer[idx-1:])
		buffer[idx] = i
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
