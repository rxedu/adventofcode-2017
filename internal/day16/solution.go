package day16

import (
	"fmt"
	"strconv"
	"strings"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input), createDancers()))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input), createDancers(), 1000000000))
}

type Step interface {
	dance([]string)
}

type SpinStep struct {
	steps int
}

func (s SpinStep) dance(dancers []string) {
	size := len(dancers)
	tmp := make([]string, size)
	for i := range dancers {
		idx := (i + s.steps) % size
		tmp[idx] = dancers[i]
	}
	copy(dancers, tmp)
}

type ExchangeStep struct {
	a int
	b int
}

func (s ExchangeStep) dance(dancers []string) {
	dancers[s.a], dancers[s.b] = dancers[s.b], dancers[s.a]
}

type PartnerStep struct {
	a string
	b string
}

func (s PartnerStep) dance(dancers []string) {
	a, ok := findDancerIndex(dancers, s.a)
	if !ok {
		return
	}
	b, ok := findDancerIndex(dancers, s.b)
	if !ok {
		return
	}
	dancers[a], dancers[b] = dancers[b], dancers[a]
}

func solvePartOne(input []Step, dancers []string) []string {
	for _, step := range input {
		step.dance(dancers)
	}
	return dancers
}

func solvePartTwo(input []Step, dancers []string, dances int) []string {
	loop, ok := findLoopPoint(input, dancers)

	count := dances
	if ok {
		count = dances % loop
	}

	for i := 0; i < count; i++ {
		for _, step := range input {
			step.dance(dancers)
		}
	}
	return dancers
}

func findLoopPoint(input []Step, dancers []string) (int, bool) {
	max := 1000000
	tmp := make([]string, len(dancers))
	copy(tmp, dancers)
	for i := 1; i < max; i++ {
		for _, step := range input {
			step.dance(tmp)

			eq := true
			for j, v := range dancers {
				eq = eq && tmp[j] == v
			}
			if eq {
				return i, true
			}
		}
	}
	return 0, false
}

func findDancerIndex(dancers []string, name string) (int, bool) {
	for i, v := range dancers {
		if v == name {
			return i, true
		}
	}
	return 0, false
}

func parse(input string) []Step {
	rows := strings.Split(input, ",")

	steps := make([]Step, len(rows))
	for i, step := range rows {
		steps[i] = parseStep(step)
	}
	return steps
}

func parseStep(step string) Step {
	switch string(step[0]) {
	case "s":
		x, err := strconv.Atoi(step[1:])
		if err != nil {
			panic(err)
		}
		return SpinStep{x}
	case "x":
		parts := strings.Split(step[1:], "/")
		a, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		return ExchangeStep{a, b}
	case "p":
		parts := strings.Split(step[1:], "/")
		a := parts[0]
		b := parts[1]
		return PartnerStep{a, b}
	}
	panic(fmt.Sprintf("Cannot parse step %v", step))
}

func serialize(output []string) string {
	return strings.Join(output, "")
}

func createDancers() []string {
	return []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"o",
		"p",
	}
}
