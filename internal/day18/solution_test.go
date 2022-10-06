package day18

import (
	"sync"
	"testing"
)

type example struct {
	i []Instruction
	o int
}

func TestPartOneExamples(t *testing.T) {
	examples := []example{
		{i: []Instruction{
			SetInstruction{Register{"a"}, Register{"1"}},
			AddInstruction{Register{"a"}, Register{"2"}},
			MultiplyInstruction{Register{"a"}, Register{"a"}},
			ModInstruction{Register{"a"}, Register{"5"}},
			PlayInstruction{Register{"a"}},
			SetInstruction{Register{"a"}, Register{"0"}},
			RecoverInstruction{Register{"a"}},
			JumpInstruction{Register{"a"}, Register{"-1"}},
			SetInstruction{Register{"a"}, Register{"1"}},
			JumpInstruction{Register{"a"}, Register{"-2"}},
		}, o: 4},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e example) {

			defer wg.Done()
			got := solvePartOne(e.i)
			if got != e.o {
				t.Errorf("\n example (%v => %v)\nsolution (%v => %v)", e.i, e.o, e.i, got)
			}
		}(e)
	}

	wg.Wait()
}

func TestPartTwoExamples(t *testing.T) {
	examples := []example{
		{i: []Instruction{}, o: 0},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e example) {

			defer wg.Done()
			got := solvePartTwo(e.i)
			if got != e.o {
				t.Errorf("\n example (%v => %v)\nsolution (%v => %v)", e.i, e.o, e.i, got)
			}
		}(e)
	}

	wg.Wait()
}
