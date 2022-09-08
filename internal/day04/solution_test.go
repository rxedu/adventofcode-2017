package day04

import (
	"sync"
	"testing"
)

type example struct {
	i []string
	o int
}

func TestPartOneExamples(t *testing.T) {
	examples := []example{
		{i: []string{""}, o: 1},
		{i: []string{"aa bb cc dd ee"}, o: 1},
		{i: []string{"aa bb cc dd aa"}, o: 0},
		{i: []string{"aa bb cc dd aaa"}, o: 1},
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
		{i: []string{""}, o: 1},
		{i: []string{"abcde fghij"}, o: 1},
		{i: []string{"abcde xyz ecdab"}, o: 0},
		{i: []string{"a ab abc abd abf abj"}, o: 1},
		{i: []string{"iiii oiii ooii oooi oooo"}, o: 1},
		{i: []string{"oiii ioii iioi iiio"}, o: 0},
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
