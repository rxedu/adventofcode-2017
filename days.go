package adventofcode

import (
	"github.com/rxedu/adventofcode-2017-go/internal/day01"
	"github.com/rxedu/adventofcode-2017-go/internal/day02"
	"github.com/rxedu/adventofcode-2017-go/internal/day03"
)

func createSolvers(part int) []Solver {
	if part == 1 {
		return []Solver{
			day01.SolvePartOne,
			day02.SolvePartOne,
			day03.SolvePartOne,
		}
	}

	if part == 2 {
		return []Solver{
			day01.SolvePartTwo,
			day02.SolvePartTwo,
			day03.SolvePartTwo,
		}
	}

	return []Solver{}
}
