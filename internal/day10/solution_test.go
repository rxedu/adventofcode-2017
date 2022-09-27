package day10

import (
	"sync"
	"testing"
)

type exampleOne struct {
	i []int
	o int
}

type exampleTwo struct {
	i []int
	o string
}

func TestPartOneExamples(t *testing.T) {
	examples := []exampleOne{
		{i: []int{1, 1}, o: 0},
		{i: []int{3, 4, 1, 5}, o: 12},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e exampleOne) {

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
	examples := []exampleTwo{
		{i: []int{}, o: "a2582a3a0e66e6e86e3812dcb672a272"},
		{i: []int{65, 111, 67, 32, 50, 48, 49, 55}, o: "33efeb34ea91902bb2f59c9920caa6cd"},
		{i: []int{49, 44, 50, 44, 51}, o: "3efbe78a8d82f29979031a4aa0b16a9d"},
		{i: []int{49, 44, 50, 44, 52}, o: "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e exampleTwo) {

			defer wg.Done()
			got := solvePartTwo(e.i)
			if got != e.o {
				t.Errorf("\n example (%v => %v)\nsolution (%v => %v)", e.i, e.o, e.i, got)
			}
		}(e)
	}

	wg.Wait()
}

func TestDenseHash(t *testing.T) {
	input := []int{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22,
		65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 23,
	}
	got := denseHash(input)

	if got[0] != 64 {
		t.Errorf("%d; want 64", got[0])
	}

	if got[1] != 65 {
		t.Errorf("%d; want 65", got[1])
	}
}

func TestToHexString(t *testing.T) {
	got := toHexString([]int{64, 7, 255})

	if got != "4007ff" {
		t.Errorf("%s; want 4007ff", got)
	}

}
