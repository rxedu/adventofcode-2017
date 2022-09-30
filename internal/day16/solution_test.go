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
		{i: []Step{
			SpinStep{1},
			ExchangeStep{3, 4},
			PartnerStep{"e", "b"},
		}, o: []string{"b", "a", "e", "d", "c"}},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e example) {

			defer wg.Done()
			got := solvePartOne(e.i, []string{"a", "b", "c", "d", "e"})
			ok := true
			for i, v := range got {
				if v != e.o[i] {
					ok = false
				}
			}
			if !ok {
				t.Errorf("\n example (%v => %v)\nsolution (%v => %v)", e.i, e.o, e.i, got)
			}
		}(e)
	}

	wg.Wait()
}

func TestPartTwoExamples(t *testing.T) {
	examples := []example{
		{i: []Step{
			SpinStep{1},
			ExchangeStep{3, 4},
			PartnerStep{"e", "b"},
		}, o: []string{"c", "e", "a", "d", "b"}},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e example) {

			defer wg.Done()
			got := solvePartTwo(e.i, []string{"a", "b", "c", "d", "e"}, 2)
			ok := true
			for i, v := range got {
				if v != e.o[i] {
					ok = false
				}
			}
			if !ok {
				t.Errorf("\n example (%v => %v)\nsolution (%v => %v)", e.i, e.o, e.i, got)
			}
		}(e)
	}

	wg.Wait()
}
