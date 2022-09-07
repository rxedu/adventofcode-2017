package adventofcode

import (
	"github.com/rxedu/adventofcode-2017-go/internal/day01"
	"github.com/rxedu/adventofcode-2017-go/internal/day02"
)

func createSolvers(part int) []Solver {
	if part == 1 {
		return []Solver{
			day01.SolvePartOne,
			day02.SolvePartOne,
		}
	}

	if part == 2 {
		return []Solver{
			day01.SolvePartTwo,
			day02.SolvePartTwo,
		}
	}

	return []Solver{}
}
