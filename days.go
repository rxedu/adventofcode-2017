package adventofcode

import (
	"github.com/rxedu/adventofcode-2017-go/internal/day01"
	"github.com/rxedu/adventofcode-2017-go/internal/day02"
	"github.com/rxedu/adventofcode-2017-go/internal/day03"
	"github.com/rxedu/adventofcode-2017-go/internal/day04"
	"github.com/rxedu/adventofcode-2017-go/internal/day05"
	"github.com/rxedu/adventofcode-2017-go/internal/day06"
	"github.com/rxedu/adventofcode-2017-go/internal/day07"
	"github.com/rxedu/adventofcode-2017-go/internal/day08"
	"github.com/rxedu/adventofcode-2017-go/internal/day09"
	"github.com/rxedu/adventofcode-2017-go/internal/day10"
	"github.com/rxedu/adventofcode-2017-go/internal/day11"
	"github.com/rxedu/adventofcode-2017-go/internal/day12"
	"github.com/rxedu/adventofcode-2017-go/internal/day13"
	"github.com/rxedu/adventofcode-2017-go/internal/day14"
	"github.com/rxedu/adventofcode-2017-go/internal/day15"
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
		day06.SolvePartOne,
		day07.SolvePartOne,
		day08.SolvePartOne,
		day09.SolvePartOne,
		day10.SolvePartOne,
		day11.SolvePartOne,
		day12.SolvePartOne,
		day13.SolvePartOne,
		day14.SolvePartOne,
		day15.SolvePartOne,
	}

	solversPartTwo = []Solver{
		day01.SolvePartTwo,
		day02.SolvePartTwo,
		day03.SolvePartTwo,
		day04.SolvePartTwo,
		day05.SolvePartTwo,
		day06.SolvePartTwo,
		day07.SolvePartTwo,
		day08.SolvePartTwo,
		day09.SolvePartTwo,
		day10.SolvePartTwo,
		day11.SolvePartTwo,
		day12.SolvePartTwo,
		day13.SolvePartTwo,
		day14.SolvePartTwo,
		day15.SolvePartTwo,
	}
}
