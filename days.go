package adventofcode

import (
	"github.com/rxedu/adventofcode-2017/internal/day01"
)

func createSolvers() []Solver {
	return []Solver{
		day01.Solve,
	}
}
