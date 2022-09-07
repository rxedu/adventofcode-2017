package day01

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
		{i: []int{1, 1, 2, 2}, o: 3},
		{i: []int{1, 1, 1, 1}, o: 4},
		{i: []int{1, 2, 3, 4}, o: 0},
		{i: []int{9, 1, 2, 1, 2, 1, 2, 9}, o: 9},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e example) {

			defer wg.Done()
			got := solvePartOne(e.i)
			if got != e.o {
				t.Errorf("\n example (%d => %d)\nsolution (%d => %d)", e.i, got, e.i, e.o)
			}
		}(e)
	}

	wg.Wait()
}
