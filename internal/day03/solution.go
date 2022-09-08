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
	return spiralSum(input).z
}

type Address struct {
	n int
	z int
	x int
	y int
}

func spiral(to int) Address {
	adr := Address{x: 0, y: 0, z: 1, n: 0}

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

func spiralSum(to int) Address {
	adr := Address{x: 0, y: 0, z: 1, n: 0}
	currentAdrs := []Address{adr}
	var prevAdrs []Address

	if adr.z > to {
		return adr
	}

	for adr.z <= to {
		adr.x++
		adr.n++
		adr.z = nextZ(adr, currentAdrs, prevAdrs)
		prevAdrs = currentAdrs
		currentAdrs = []Address{adr}
		if adr.z > to {
			return adr
		}

		for adr.y < adr.n {
			adr.y++
			adr.z = nextZ(adr, currentAdrs, prevAdrs)
			currentAdrs = append(currentAdrs, adr)
			if adr.z > to {
				return adr
			}
		}

		for adr.x > -adr.n {
			adr.x--
			adr.z = nextZ(adr, currentAdrs, prevAdrs)
			currentAdrs = append(currentAdrs, adr)
			if adr.z > to {
				return adr
			}
		}

		for adr.y > -adr.n {
			adr.y--
			adr.z = nextZ(adr, currentAdrs, prevAdrs)
			currentAdrs = append(currentAdrs, adr)
			if adr.z > to {
				return adr
			}
		}

		for adr.x < adr.n {
			adr.x++
			adr.z = nextZ(adr, currentAdrs, prevAdrs)
			currentAdrs = append(currentAdrs, adr)
			if adr.z > to {
				return adr
			}
		}
	}

	adr.x++
	adr.n++
	adr.z = nextZ(adr, currentAdrs, prevAdrs)
	return adr
}

func nextZ(cur Address, a []Address, b []Address) int {
	sum := 0

	for _, adr := range a {
		diffX := adr.x - cur.x
		diffY := adr.y - cur.y
		if abs(diffX)+abs(diffY) == 1 || diffX*diffX+diffY*diffY == 2 {
			sum += adr.z
		}
	}

	for _, adr := range b {
		diffX := adr.x - cur.x
		diffY := adr.y - cur.y
		if abs(diffX)+abs(diffY) == 1 || diffX*diffX+diffY*diffY == 2 {
			sum += adr.z
		}
	}

	return sum
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
