package day18

import (
	"fmt"
	"strconv"
	"strings"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input)))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input)))
}

type Instruction interface {
	exec(i *int, registers map[string]int, recovered *[]int, played *[]int)
}

type PlayInstruction struct {
	x Register
}

func (instruction PlayInstruction) exec(i *int, registers map[string]int, recovered *[]int, played *[]int) {
	x := instruction.x.Get(registers)
	j := *i + 1
	*i = j
	*played = append(*played, x)
}

type SetInstruction struct {
	x Register
	y Register
}

func (instruction SetInstruction) exec(i *int, registers map[string]int, recovered *[]int, played *[]int) {
	x := instruction.x.Label
	y := instruction.y.Get(registers)
	j := *i + 1
	*i = j
	registers[x] = y
}

type AddInstruction struct {
	x Register
	y Register
}

func (instruction AddInstruction) exec(i *int, registers map[string]int, recovered *[]int, played *[]int) {
	x := instruction.x.Label
	y := instruction.y.Get(registers)
	j := *i + 1
	*i = j
	registers[x] += y
}

type MultiplyInstruction struct {
	x Register
	y Register
}

func (instruction MultiplyInstruction) exec(i *int, registers map[string]int, recovered *[]int, played *[]int) {
	x := instruction.x.Label
	y := instruction.y.Get(registers)
	j := *i + 1
	*i = j
	registers[x] *= y
}

type ModInstruction struct {
	x Register
	y Register
}

func (instruction ModInstruction) exec(i *int, registers map[string]int, recovered *[]int, played *[]int) {
	x := instruction.x.Label
	y := instruction.y.Get(registers)
	j := *i + 1
	*i = j
	registers[x] %= y
}

type RecoverInstruction struct {
	x Register
}

func (instruction RecoverInstruction) exec(i *int, registers map[string]int, recovered *[]int, played *[]int) {
	x := instruction.x.Get(registers)
	j := *i + 1
	*i = j
	if x == 0 {
		return
	}
	*recovered = append(*recovered, x)
}

type JumpInstruction struct {
	x Register
	y Register
}

func (instruction JumpInstruction) exec(i *int, registers map[string]int, recovered *[]int, played *[]int) {
	x := instruction.x.Get(registers)
	y := instruction.y.Get(registers)
	offset := 1
	if x > 0 {
		offset = y
	}
	j := *i + offset
	*i = j
}

type Register struct {
	Label string
}

func (r Register) Get(registers map[string]int) int {
	x, err := strconv.Atoi(r.Label)
	if err != nil {
		return registers[r.Label]
	}
	return x
}

func solvePartOne(input []Instruction) int {
	var recovered []int
	var played []int
	registers := make(map[string]int, 2*len(input))
	for i := 0; 0 <= i && i < len(input); {
		input[i].exec(&i, registers, &recovered, &played)
		if len(recovered) > 0 && len(played) > 0 {
			return played[len(played)-1]
		}
	}
	return 0
}

func solvePartTwo(input []Instruction) int {
	return 0
}

func parse(input string) []Instruction {
	rows := strings.Split(input, "\n")

	instructions := make([]Instruction, len(rows))
	for i, instruction := range rows {
		instructions[i] = parseInstruction(instruction)
	}
	return instructions
}

func parseInstruction(instruction string) Instruction {
	parts := strings.Split(instruction, " ")
	switch parts[0] {
	case "snd":
		return PlayInstruction{Register{parts[1]}}
	case "set":
		return SetInstruction{Register{parts[1]}, Register{parts[2]}}
	case "add":
		return AddInstruction{Register{parts[1]}, Register{parts[2]}}
	case "mul":
		return MultiplyInstruction{Register{parts[1]}, Register{parts[2]}}
	case "mod":
		return ModInstruction{Register{parts[1]}, Register{parts[2]}}
	case "rcv":
		return RecoverInstruction{Register{parts[1]}}
	case "jgz":
		return JumpInstruction{Register{parts[1]}, Register{parts[2]}}
	}
	panic(fmt.Sprintf("Cannot parse instruction %v", instruction))
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
