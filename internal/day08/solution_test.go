package day08

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
		{i: []Instruction{}, o: 0},
		{i: []Instruction{{reg: "a", change: 10, condition: Condition{reg: "a", op: "==", val: 0}}}, o: 10},
		{i: []Instruction{
			{reg: "b", change: 5, condition: Condition{reg: "a", op: ">", val: 1}},
			{reg: "a", change: 1, condition: Condition{reg: "b", op: "<", val: 5}},
			{reg: "c", change: 10, condition: Condition{reg: "a", op: ">=", val: 1}},
			{reg: "c", change: -20, condition: Condition{reg: "c", op: "==", val: 10}},
		}, o: 1},
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
		{i: []Instruction{{reg: "a", change: 10, condition: Condition{reg: "a", op: "==", val: 0}}}, o: 10},
		{i: []Instruction{
			{reg: "b", change: 5, condition: Condition{reg: "a", op: ">", val: 1}},
			{reg: "a", change: 1, condition: Condition{reg: "b", op: "<", val: 5}},
			{reg: "c", change: 10, condition: Condition{reg: "a", op: ">=", val: 1}},
			{reg: "c", change: -20, condition: Condition{reg: "c", op: "==", val: 10}},
		}, o: 10},
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
