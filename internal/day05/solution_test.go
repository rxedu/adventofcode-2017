package day05

import (
	"sync"
	"testing"
)

type example struct {
	i []int
	o int
}

func TestPartOneExamples(t *testing.T) {
	examples := []example{
		{i: []int{}, o: 0},
		{i: []int{0}, o: 2},
		{i: []int{0, 1}, o: 3},
		{i: []int{0, 3, 0, 1, -3}, o: 5},
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
		{i: []int{}, o: 0},
		{i: []int{0}, o: 1},
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
