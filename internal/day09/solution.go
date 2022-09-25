package day09

import (
	"regexp"
	"strconv"
	"strings"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input)))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input)))
}

func solvePartOne(input string) int {
	score := 0
	level := 0
	for _, v := range removeGarbage(input) {
		switch string(v) {
		case "{":
			level++
		case "}":
			score += level
			level--
		}
	}
	return score
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

		if string(v) == ">" {
			inGarbage = false
			continue
		}

		if string(v) == "<" {
			inGarbage = true
			continue
		}

		if !inGarbage {
			result += string(v)
		}
	}

	multicomma := regexp.MustCompile("(,)+")
	singleCommas := multicomma.ReplaceAllString(result, ",")
	return strings.ReplaceAll(
		strings.ReplaceAll(singleCommas, ",}", "}"),
		"{,", "{")
}
