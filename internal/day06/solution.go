package day06

import (
	"fmt"
	"strconv"
	"strings"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input)))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input)))
}

func solvePartOne(input []int) int {
	cycles, _ := debuggerCycles(input)
	return cycles
}

func solvePartTwo(input []int) int {
	_, loopSize := debuggerCycles(input)
	return loopSize
}

func debuggerCycles(input []int) (int, int) {
	seen := make(map[string]int)

	for count := 0; count < 10000000000; count++ {
		cfg := fmt.Sprint(input)
		if seen[cfg] > 0 {
			return count, count - seen[cfg]
		}
		seen[cfg] = count
		rebalance(input)
	}
	return 0, 0
}

func rebalance(mem []int) {
	cur := getStart(mem)
	blocks := mem[cur]
	mem[cur] = 0
	for i := 0; i < blocks; i++ {
		cur = (cur + 1) % len(mem)
		mem[cur]++
	}

}

func getStart(mem []int) int {
	largestValue := 0
	startAt := 0
	for i, v := range mem {
		if v > largestValue {
			largestValue = v
			startAt = i
		}
	}
	return startAt
}

func parse(input string) []int {
	cols := strings.Split(input, "\t")
	arr := make([]int, len(cols))
	for i, str := range cols {
		v, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		arr[i] = v
	}
	return arr
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
