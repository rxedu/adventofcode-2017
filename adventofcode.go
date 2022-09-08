package adventofcode

type Solver = func(string) string

func SolveDay(day int, part int, input string) (string, bool) {
	solver, ok := GetSolver(day, part)
	if !ok {
		return solver(""), false
	}

	solution := solver(input)
	return solution, true
}

func GetSolver(day int, part int) (Solver, bool) {
	solvers := getSolvers(part)

	defaultSolver := func(input string) string { return "" }

	if day < 1 {
		return defaultSolver, false
	}

	if day > len(solvers) {
		return defaultSolver, false
	}

	return solvers[day-1], true
}

func NumSolvedDays() int {
	if len(solversPartOne) == 0 || len(solversPartTwo) == 0 {
		initSolvers()
	}

	return len(solversPartTwo)
}

func getSolvers(part int) []Solver {
	if len(solversPartOne) == 0 || len(solversPartTwo) == 0 {
		initSolvers()
	}

	if part == 1 {
		return solversPartOne
	}

	if part == 2 {
		return solversPartTwo
	}

	return []Solver{}
}
