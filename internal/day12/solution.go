package day12

import (
	"strconv"
	"strings"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input)))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input)))
}

func solvePartOne(input map[int][]int) int {
	group := groupBy(0, input)
	return len(group)
}

func solvePartTwo(input map[int][]int) int {
	seen := make(map[int]bool, len(input))
	count := 0
	for idx := range input {
		if seen[idx] {
			continue
		}
		group := groupBy(idx, input)
		for _, v := range group {
			seen[v] = true
		}
		count++
	}
	return count
}

func groupBy(ID int, graph map[int][]int) []int {
	seen := make(map[int]bool, len(graph))
	var group []int

	var collect func(target int)
	collect = func(target int) {
		if seen[target] {
			return
		}
		group = append(group, target)
		seen[target] = true
		for _, v := range graph[target] {
			collect(v)
		}
	}
	collect(ID)

	return group
}

func parse(input string) map[int][]int {
	rows := strings.Split(input, "\n")
	links := make(map[int][]int, len(rows))
	for _, link := range rows {
		parts := strings.Split(link, " <-> ")
		connections := strings.Split(parts[1], ", ")

		idx, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		val := make([]int, len(connections))
		for i, v := range connections {
			c, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			val[i] = c
		}

		links[idx] = val
	}

	return links
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
