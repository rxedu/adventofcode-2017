package adventofcode

import (
	"github.com/rxedu/adventofcode-2017-go/internal/day01"
)

func createSolvers() []Solver {
	return []Solver{
		day01.Solve,
	}
}
