package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rxedu/adventofcode-2017-go/internal/day10"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(input))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(input))
}

func solvePartOne(input string) int {
	_, count := makeGrid(input)
	return count
}

func solvePartTwo(input string) int {
	count := 0
	grid, _ := makeGrid(input)
	for i := range grid {
		for j, v := range grid[i] {
			if v != 0 {
				count++
			}
			clearGroupStartingAt(grid, i, j)
		}
	}
	return count
}

func clearGroupStartingAt(grid [][]int, i int, j int) {
	iInRange := 0 <= i && i < len(grid)
	jInRange := 0 <= j && j < len(grid)
	if !iInRange || !jInRange || grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	clearGroupStartingAt(grid, i+1, j)
	clearGroupStartingAt(grid, i-1, j)
	clearGroupStartingAt(grid, i, j+1)
	clearGroupStartingAt(grid, i, j-1)
}

func makeGrid(input string) ([][]int, int) {
	size := 128
	count := 0
	grid := make([][]int, size)
	for i := 0; i < size; i++ {
		h := day10.KnotHash(fmt.Sprintf("%s-%d", input, i))
		row := strings.Split(h.Bin(), "")
		grid[i] = make([]int, len(row))
		for j, v := range row {
			w, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			grid[i][j] = w
			if w == 1 {
				count++
			}
		}
	}
	return grid, count
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
