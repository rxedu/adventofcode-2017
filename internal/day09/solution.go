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
	groups, _ := removeGarbage(input)
	for _, v := range groups {
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
	_, count := removeGarbage(input)
	return count
}

func parse(input string) string {
	return input
}

func serialize(output int) string {
	return strconv.Itoa(output)
}

func removeGarbage(input string) (string, int) {
	result := ""
	inGarbage := false
	isNegated := false
	count := 0
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
			if inGarbage {
				count++
			}
			inGarbage = true
			continue
		}

		if !inGarbage {
			result += string(v)
		} else {
			count++
		}
	}

	multicomma := regexp.MustCompile("(,)+")
	result = multicomma.ReplaceAllString(result, ",")
	result = strings.ReplaceAll(result, ",}", "}")
	result = strings.ReplaceAll(result, "{,", "{")
	return result, count
}
