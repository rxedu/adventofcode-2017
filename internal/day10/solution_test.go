package day10

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
		{i: []int{3, 4, 1, 5}, o: 12},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e example) {

			defer wg.Done()
			got := solvePartOne(e.i, 5)
			if got != e.o {
				t.Errorf("\n example (%v => %v)\nsolution (%v => %v)", e.i, e.o, e.i, got)
			}
		}(e)
	}

	wg.Wait()
}

func TestPartTwoExamples(t *testing.T) {
	examples := []example{
		{i: []int{1, 1}, o: 1},
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
