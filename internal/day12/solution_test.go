package day12

import (
	"sync"
	"testing"
)

type example struct {
	i map[int][]int
	o int
}

func TestPartOneExamples(t *testing.T) {
	examples := []example{
		{i: map[int][]int{0: {}}, o: 1},
		{i: map[int][]int{0: {1}, 1: {0}}, o: 2},
		{i: map[int][]int{
			0: {2},
			1: {1},
			2: {0, 3, 4},
			3: {2, 4},
			4: {2, 3, 6},
			5: {6},
			6: {4, 5},
		}, o: 6},
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
		{i: map[int][]int{}, o: 0},
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
