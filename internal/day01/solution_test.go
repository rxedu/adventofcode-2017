package day01

import (
	"sync"
	"testing"
)

type example struct {
	i int
	o int
}

func TestExamples(t *testing.T) {
	examples := []example{
		{i: 1122, o: 3},
		{i: 1111, o: 4},
		{i: 1234, o: 0},
		{i: 91212129, o: 9},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e example) {

			defer wg.Done()
			got := solve(e.i)
			if got != e.o {
				t.Errorf("\n example (%d => %d)\nsolution (%d => %d)", e.i, got, e.i, e.o)
			}
		}(e)
	}

	wg.Wait()
}
