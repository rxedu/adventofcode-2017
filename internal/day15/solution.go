package day15

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

type Generators struct {
	a int
	b int
}

func solvePartOne(input Generators) int {
	return 0
}

func solvePartTwo(input Generators) int {
	return 0
}

func parse(input string) Generators {
	gens := strings.Split(input, "\n")
	genA := strings.Split(gens[0], "Generator A starts with ")
	genB := strings.Split(gens[1], "Generator B starts with ")

	a, err := strconv.Atoi(genA[1])
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(genB[1])
	if err != nil {
		panic(err)
	}

	return Generators{a, b}
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
