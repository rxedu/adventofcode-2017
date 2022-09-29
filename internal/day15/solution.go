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

const aFactor = 16807
const bFactor = 48271
const divisor = 2147483647

type Generators struct {
	a Generator
	b Generator
}

type Generator struct {
	prev   uint64
	factor uint64
}

func (g *Generator) next() {
	next := (g.prev * g.factor) % divisor
	g.prev = next
}

func solvePartOne(input Generators) int {
	cycles := 40000000
	count := 0
	for i := 0; i < cycles; i++ {
		if judge(input.a.prev, input.b.prev) {
			count++
		}
		input.a.next()
		input.b.next()
	}
	return count
}

func solvePartTwo(input Generators) int {
	return 0
}

func judge(a uint64, b uint64) bool {
	return uint16(a) == uint16(b)
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

	return Generators{
		Generator{uint64(a), aFactor},
		Generator{uint64(b), bFactor},
	}
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
