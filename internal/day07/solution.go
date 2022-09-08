package day07

import (
	"strconv"
	"strings"
)

func SolvePartOne(input string) string {
	return solvePartOne(parse(input))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input)))
}

type Node struct {
	name   string
	weight int
	links  []string
}

type Tree struct {
	root Node
	subs []Tree
}

func (t Tree) subweight() int {
	w := 0
	for _, n := range t.subs {
		w += n.root.weight
		w += n.subweight()
	}
	return w
}

func solvePartOne(input []Node) string {
	root := findRoot(input)
	return root.name
}

func solvePartTwo(input []Node) int {
	// TODO
	// make the tree
	// for each subtree, compute which one has weight not equal to the rest
	// if all are equal (or no more subtrees), then the current subtree is the problem
	// the diff between this node's weight and the other weights on it's level is the solution
	return 0
}

func findRoot(input []Node) Node {
	if len(input) < 1 {
		return Node{}
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
			return node
		}
	}
	return Node{}
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

func serialize(output int) string {
	return strconv.Itoa(output)
}
