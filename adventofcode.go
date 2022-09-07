package adventofcode

type Solver = func(string) string

func SolveDay(day int, input string) (string, bool) {
	solver, ok := GetSolver(day)
	if !ok {
		return solver(""), false
	}

	solution := solver(input)
	return solution, true
}

func GetSolver(day int) (Solver, bool) {
	solvers := createSolvers()

	defaultSolver := func(input string) string { return "" }

	if day < 1 {
		return defaultSolver, false
	}

	if day > len(solvers) {
		return defaultSolver, false
	}

	return solvers[day-1], true
}
