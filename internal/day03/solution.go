package day03

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
	adr := spiral(input)
	return abs(adr.x) + abs(adr.y)
}

func solvePartTwo(input int) int {
	return 0
}

type address struct {
	n int
	z int
	x int
	y int
}

func spiral(to int) address {
	adr := address{x: 0, y: 0, z: 1, n: 0}

	if adr.z == to {
		return adr
	}

	for adr.z < to {
		adr.x++
		adr.z++
		adr.n++
		if adr.z == to {
			return adr
		}

		for adr.y < adr.n {
			adr.y++
			adr.z++
			if adr.z == to {
				return adr
			}
		}

		for adr.x > -adr.n {
			adr.x--
			adr.z++
			if adr.z == to {
				return adr
			}
		}

		for adr.y > -adr.n {
			adr.y--
			adr.z++
			if adr.z == to {
				return adr
			}
		}

		for adr.x < adr.n {
			adr.x++
			adr.z++
			if adr.z == to {
				return adr
			}
		}
	}

	return adr
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
