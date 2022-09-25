package day09

import (
	"strconv"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input)))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input)))
}

func solvePartOne(input string) int {
	return 0
}

func solvePartTwo(input string) int {
	return 0
}

func parse(input string) string {
	return input
}

func serialize(output int) string {
	return strconv.Itoa(output)
}

func removeGarbage(input string) string {
	result := ""
	inGarbage := false
	isNegated := false
	for _, v := range input {
		if isNegated {
			isNegated = false
			continue
		}

		if string(v) == "!" {
			isNegated = true
			continue
		}

		if !inGarbage && string(v) == "<" {
			inGarbage = true
			continue
		}

		if inGarbage && string(v) == ">" {
			inGarbage = false
			continue
		}

		if !inGarbage && string(v) != "," {
			result += string(v)
		}
	}
	return result
}
