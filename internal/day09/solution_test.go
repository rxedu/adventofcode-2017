package day09

import (
	"sync"
	"testing"
)

type example struct {
	i string
	o int
}

func TestPartOneExamples(t *testing.T) {
	examples := []example{
		{i: "{}", o: 1},
		{i: "{{{}}}", o: 1},
		{i: "{{},{}}", o: 5},
		{i: "{{{},{},{{}}}}", o: 16},
		{i: "{<a>,<a>,<a>,<a>}", o: 1},
		{i: "{{<ab>},{<ab>},{<ab>},{<ab>}}", o: 9},
		{i: "{<!!>},{<!!>},{<!!>},{<!!>}}", o: 9},
		{i: "{{<a!>},{<a!>},{<a!>},{<ab>}}", o: 3},
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
		{i: "", o: 0},
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

type garbageExample struct {
	i string
	o string
}

func TestRemoveGarbage(t *testing.T) {
	examples := []garbageExample{
		{i: "{}", o: "{}"},
		{i: "{{},{},{}}", o: "{{},{},{}}"},
		{i: "{{},<>,{}}", o: "{{},{}}"},
		{i: "{<>}", o: "{}"},
		{i: "{<>,<>}", o: "{}"},
		{i: "{{},<>}", o: "{{}}"},
		{i: "{<>,{}}", o: "{{}}"},
		{i: "{<!>>}", o: "{}"},
		{i: "{<!!>}", o: "{}"},
		{i: "{<a>,<a>,<a>,<a>}", o: "{}"},
		{i: "{{<ab>},{<ab>},{<ab>},{<ab>}}", o: "{{},{},{},{}}"},
		{i: "{{<!!>},{<!!>},{<!!>},{<!!>}}", o: "{{},{},{},{}}"},
		{i: "{{<a!>},{<a!>},{<a!>},{<ab>}}", o: "{{}}"},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e garbageExample) {

			defer wg.Done()
			got := removeGarbage(e.i)
			if got != e.o {
				t.Errorf("\n example (%v => %v)\nsolution (%v => %v)", e.i, e.o, e.i, got)
			}
		}(e)
	}

	wg.Wait()
}
