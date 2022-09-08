package day07

import (
	"sync"
	"testing"
)

type exampleOne struct {
	i []Node
	o string
}

type exampleTwo struct {
	i []Node
	o int
}

func TestPartOneExamples(t *testing.T) {
	examples := []exampleOne{
		{i: []Node{}, o: ""},
		{i: []Node{{name: "a", weight: 0}}, o: "a"},
		{i: []Node{
			{name: "pbga", weight: 66},
			{name: "xhth", weight: 57},
			{name: "ebii", weight: 61},
			{name: "havc", weight: 66},
			{name: "ktlj", weight: 57},
			{name: "fwft", weight: 72, links: []string{"ktlj", "cntj", "xhth"}},
			{name: "qoyq", weight: 66},
			{name: "padx", weight: 45, links: []string{"pbga", "havc", "qoyq"}},
			{name: "tknk", weight: 41, links: []string{"ugml", "padx", "fwft"}},
			{name: "jptl", weight: 61},
			{name: "ugml", weight: 68, links: []string{"gyxo", "ebii", "jptl"}},
			{name: "gyxo", weight: 61},
			{name: "cntj", weight: 57},
		}, o: "tknk"},
	}

	var wg sync.WaitGroup
	for _, e := range examples {
		wg.Add(1)
		go func(e exampleOne) {

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
	examples := []exampleTwo{
		{i: []Node{
			{name: "pbga", weight: 66},
			{name: "xhth", weight: 57},
			{name: "ebii", weight: 61},
			{name: "havc", weight: 66},
			{name: "ktlj", weight: 57},
			{name: "fwft", weight: 72, links: []string{"ktlj", "cntj", "xhth"}},
			{name: "qoyq", weight: 66},
			{name: "padx", weight: 45, links: []string{"pbga", "havc", "qoyq"}},
			{name: "tknk", weight: 41, links: []string{"ugml", "padx", "fwft"}},
			{name: "jptl", weight: 61},
			{name: "ugml", weight: 68, links: []string{"gyxo", "ebii", "jptl"}},
			{name: "gyxo", weight: 61},
			{name: "cntj", weight: 57},
		}, o: 60},
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
