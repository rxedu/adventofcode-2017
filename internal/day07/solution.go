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

func (t Tree) weight() int {
	return t.root.weight + t.subweight()
}

func solvePartOne(input []Node) string {
	root := findRoot(input)
	return root.name
}

func solvePartTwo(input []Node) int {
	root := findRoot(input)
	tree := createTree(root, input)
	return findCorrectWeight(tree, 0)
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

func createTree(root Node, nodes []Node) Tree {
	tree := Tree{root: root}
	for _, name := range root.links {
		subroot := findNodeByName(name, nodes)
		tree.subs = append(tree.subs, createTree(subroot, nodes))
	}
	return tree
}

func findNodeByName(name string, nodes []Node) Node {
	for _, n := range nodes {
		if n.name == name {
			return n
		}
	}
	return Node{}
}

func findCorrectWeight(tree Tree, imbalance int) int {
	weight := tree.root.weight + imbalance
	if isBalanced(tree.subs) {
		return weight
	}

	// NOTE: For the case of 2 children, we assume the input is constrained
	// such that solution will be the same for either branch.
	if len(tree.subs) < 3 {
		return findCorrectWeight(tree.subs[0], imbalance)
	}

	var w int
	if tree.subs[0].weight() == tree.subs[1].weight() || tree.subs[0].weight() == tree.subs[2].weight() {
		w = tree.subs[0].weight()
	} else {
		w = tree.subs[1].weight()
	}

	for i := 0; i < len(tree.subs); i++ {
		imbalance = w - tree.subs[i].weight()
		if imbalance != 0 {
			return findCorrectWeight(tree.subs[i], imbalance)
		}
	}

	return weight
}

func isBalanced(trees []Tree) bool {
	if len(trees) < 2 {
		return true
	}

	for i := 1; i < len(trees); i++ {
		if trees[i].subweight() != trees[i-1].subweight() {
			return false
		}
	}

	return true
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
