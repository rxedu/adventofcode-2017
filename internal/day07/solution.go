package day07

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

type Node struct {
	name   string
	weight int
	links  []string
}

func solvePartOne(input []Node) string {
	if len(input) < 1 {
		return ""
	}

	var links []string
	for _, v := range input {
		links = append(links, v.links...)
	}

	for _, node := range input {
		uniq := true
		for _, link := range links {
			if link == node.name {
				uniq = false
				break
			}
		}
		if uniq {
			return node.name
		}
	}
	return Node{}.name
}

func solvePartTwo(input []Node) string {
	return ""
}

func parse(input string) []Node {
	strRows := strings.Split(input, "\n")

	matrix := make([]Node, len(strRows))
	for i, strRow := range strRows {
		parts := strings.Split(strRow, " ")
		weight, err := strconv.Atoi(
			strings.TrimSuffix(strings.TrimPrefix(parts[1], "("), ")"))
		if err != nil {
			panic(err)
		}
		linkParts := strings.Split(strRow, " -> ")

		var links []string
		if len(linkParts) > 1 {
			links = strings.Split(linkParts[1], ", ")
		}

		matrix[i] = Node{
			name:   parts[0],
			weight: weight,
			links:  links,
		}
	}

	return matrix
}

func serialize(output string) string {
	return output
}
