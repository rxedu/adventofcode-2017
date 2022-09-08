package day03

import (
	"sync"
	"testing"
)

type example struct {
	i int
	o int
}

func TestPartOneExamples(t *testing.T) {
	examples := []example{
		{i: 1, o: 0},
		{i: 12, o: 3},
		{i: 23, o: 2},
		{i: 1024, o: 31},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e example) {

			defer wg.Done()
			got := solvePartOne(e.i)
			if got != e.o {
				t.Errorf("\n example (%d => %d)\nsolution (%d => %d)", e.i, e.o, e.i, got)
			}
		}(e)
	}

	wg.Wait()
}

func TestPartTwoExamples(t *testing.T) {
	examples := []example{
		{i: 0, o: 1},
		{i: 1, o: 2},
		{i: 10, o: 11},
		{i: 351, o: 362},
		{i: 747, o: 806},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e example) {

			defer wg.Done()
			got := solvePartTwo(e.i)
			if got != e.o {
				t.Errorf("\n example (%d => %d)\nsolution (%d => %d)", e.i, e.o, e.i, got)
			}
		}(e)
	}

	wg.Wait()
}
