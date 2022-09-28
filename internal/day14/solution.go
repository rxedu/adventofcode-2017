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
	size := 128
	count := 0
	for i := 0; i < size; i++ {
		h := day10.KnotHash(fmt.Sprintf("%s-%d", input, i))
		for _, v := range strings.Split(h.Bin(), "") {
			w, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			if w == 1 {
				count++
			}
		}
	}
	return count
}

func solvePartTwo(input string) int {
	return 0
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
