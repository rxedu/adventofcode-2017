package day08

import (
	"strconv"
	"strings"

	"github.com/rxedu/adventofcode-2017-go/internal/math"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input)))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input)))
}

type Instruction struct {
	reg       string
	change    int
	condition Condition
}

func (ins Instruction) exec(regs map[string]int) {
	if ins.condition.isTrue(regs) {
		regs[ins.reg] += ins.change
	}
}

type Condition struct {
	reg string
	op  string
	val int
}

func (con Condition) isTrue(regs map[string]int) bool {
	switch con.op {
	case "==":
		return regs[con.reg] == con.val
	case "!=":
		return regs[con.reg] != con.val
	case ">":
		return regs[con.reg] > con.val
	case ">=":
		return regs[con.reg] >= con.val
	case "<":
		return regs[con.reg] < con.val
	case "<=":
		return regs[con.reg] <= con.val
	}

	return false
}

func solvePartOne(input []Instruction) int {
	regs := make(map[string]int)

	for _, ins := range input {
		ins.exec(regs)
	}

	max := 0
	for _, v := range regs {
		max = math.Max(max, v)
	}
	return max
}

func solvePartTwo(input []Instruction) int {
	regs := make(map[string]int)

	max := 0
	for _, ins := range input {
		ins.exec(regs)
		max = math.Max(max, regs[ins.reg])
	}

	return max
}

func parse(input string) []Instruction {
	strRows := strings.Split(input, "\n")

	arr := make([]Instruction, len(strRows))
	for i, strRow := range strRows {
		parts := strings.Split(strRow, " ")

		change, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(err)
		}

		dir := 1
		if parts[1] == "dec" {
			dir = -1
		}

		val, err := strconv.Atoi(parts[6])
		if err != nil {
			panic(err)
		}

		arr[i] = Instruction{
			reg:    parts[0],
			change: dir * change,
			condition: Condition{
				reg: parts[4],
				op:  parts[5],
				val: val,
			},
		}
	}

	return arr
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
