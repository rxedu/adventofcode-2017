package adventofcode

import (
	"github.com/rxedu/adventofcode-2017-go/internal/day01"
	"github.com/rxedu/adventofcode-2017-go/internal/day02"
	"github.com/rxedu/adventofcode-2017-go/internal/day03"
	"github.com/rxedu/adventofcode-2017-go/internal/day04"
	"github.com/rxedu/adventofcode-2017-go/internal/day05"
)

var solversPartOne []Solver
var solversPartTwo []Solver

func initSolvers() {
	solversPartOne = []Solver{
		day01.SolvePartOne,
		day02.SolvePartOne,
		day03.SolvePartOne,
		day04.SolvePartOne,
		day05.SolvePartOne,
	}

	solversPartTwo = []Solver{
		day01.SolvePartTwo,
		day02.SolvePartTwo,
		day03.SolvePartTwo,
		day04.SolvePartTwo,
		day05.SolvePartTwo,
	}
}
