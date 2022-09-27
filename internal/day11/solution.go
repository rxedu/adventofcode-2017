package day11

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

type Step struct {
	q int
	r int
	s int
}

type Loc struct {
	q int
	r int
	s int
}

func (x Loc) add(v Step) Loc {
	return Loc{x.q + v.q, x.r + v.r, x.s + v.s}
}

func (x Loc) dist(y Loc) int {
	return math.Max(
		math.Max(math.AbsInt(y.q-x.q), math.AbsInt(y.r-x.r)),
		math.AbsInt(y.s-x.s))
}

func solvePartOne(input []Step) int {
	origin := Loc{0, 0, 0}
	x := origin
	for _, v := range input {
		x = x.add(v)
	}
	return x.dist(origin)
}

func solvePartTwo(input []Step) int {
	origin := Loc{0, 0, 0}
	x := origin
	var max int
	for _, v := range input {
		max = math.Max(max, x.dist(origin))
		x = x.add(v)
	}
	return max
}

func parse(input string) []Step {
	steps := strings.Split(input, ",")
	arr := make([]Step, len(steps))
	for i, step := range steps {
		arr[i] = parseStep(step)
	}
	return arr
}

func parseStep(step string) Step {
	vectors := map[string]Step{
		"nw": {-1, 0, 1},
		"n":  {0, -1, 1},
		"ne": {1, -1, 0},
		"sw": {-1, 1, 0},
		"s":  {0, 1, -1},
		"se": {1, 0, -1},
	}
	return vectors[step]
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
