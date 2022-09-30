package day16

import (
	"sync"
	"testing"
)

type example struct {
	i []Step
	o []string
}

func TestPartOneExamples(t *testing.T) {
	examples := []example{
		{i: []Step{}, o: []string{"a", "b", "c", "d", "e"}},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e example) {

			defer wg.Done()
			got := solvePartOne(e.i, []string{"a", "b", "c", "d", "e"})
			for i, v := range got {
				if v != e.o[i] {
					t.Errorf("\n example (%v => %v)\nsolution (%v => %v)", e.i, e.o, e.i, v)
				}
			}
		}(e)
	}

	wg.Wait()
}

func TestPartTwoExamples(t *testing.T) {
	examples := []example{
		{i: []Step{}, o: []string{"a", "b", "c", "d", "e"}},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e example) {

			defer wg.Done()
			got := solvePartTwo(e.i, []string{"a", "b", "c", "d", "e"})
			for i, v := range got {
				if v != e.o[i] {
					t.Errorf("\n example (%v => %v)\nsolution (%v => %v)", e.i, e.o, e.i, v)
				}
			}
		}(e)
	}

	wg.Wait()
}
