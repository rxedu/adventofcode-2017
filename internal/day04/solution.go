package day04

import (
	"sort"
	"strconv"
	"strings"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input)))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input)))
}

func solvePartOne(input []string) int {
	var count int
	for _, v := range input {
		if isValidPartOnePassphrase(v) {
			count++
		}
	}
	return count
}

func solvePartTwo(input []string) int {
	var count int
	for _, v := range input {
		if isValidPartTwoPassphrase(v) {
			count++
		}
	}
	return count
}

func isValidPartOnePassphrase(pass string) bool {
	words := strings.Split(pass, " ")
	m := make(map[string]int, len(words))
	for _, v := range words {
		m[v]++
		if m[v] > 1 {
			return false
		}
	}
	return true
}

func isValidPartTwoPassphrase(pass string) bool {
	words := strings.Split(pass, " ")
	m := make(map[string]int, len(words))
	for _, v := range words {
		k := strings.Split(v, "")
		sort.Strings(k)
		v = strings.Join(k, "")
		m[v]++
		if m[v] > 1 {
			return false
		}
	}
	return true
}

func parse(input string) []string {
	return strings.Split(input, "\n")
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
