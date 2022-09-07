package day02

import (
	"sync"
	"testing"
)

type example struct {
	i [][]int
	o int
}

func TestPartOneExamples(t *testing.T) {
	examples := []example{
		{i: [][]int{}, o: 0},
		{i: [][]int{{1}}, o: 0},
		{i: [][]int{{1}, {2}}, o: 0},
		{i: [][]int{{1, 2}}, o: 1},
		{i: [][]int{{1, 2}, {2, 5}}, o: 4},
		{i: [][]int{
			{5, 1, 9, 5},
			{7, 5, 3},
			{2, 4, 6, 8},
		}, o: 18},
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
		{i: [][]int{{}}, o: 0},
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
