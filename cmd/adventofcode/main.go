package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/rxedu/adventofcode-2017-go"
)

const INPUT_PATH = "input"
const OUTPUT_PATH = "solutions"

var wg sync.WaitGroup

func main() {
	day := 1
	if len(os.Args) > 1 {
		v, err := strconv.Atoi(os.Args[1])
		if err != nil || day < 1 {
			fmt.Printf("Day must be positive integer, got %v\n", day)
			os.Exit(2)
		}
		day = v
	}

	part := 1
	if len(os.Args) > 2 {
		v, err := strconv.Atoi(os.Args[2])
		if err != nil || part > 2 {
			fmt.Printf("Part must be 1 or 2, got %v\n", part)
			os.Exit(2)
		}
		part = v
	}

	ch := make(chan Solution)

	var solutions []Solution
	go func() {
		for s := range ch {
			solutions = append(solutions, s)
		}
		wg.Done()
	}()

	if len(os.Args) == 2 {
		wg.Add(2)
		go solveOne(day, 1, ch)
		go solveOne(day, 2, ch)
	}

	if len(os.Args) > 2 {
		wg.Add(1)
		go solveOne(day, part, ch)
	}

	if len(os.Args) == 1 {
		for i := 1; i <= adventofcode.NumSolvedDays(); i++ {
			wg.Add(2)
			go solveOne(i, 1, ch)
			go solveOne(i, 2, ch)
		}
	}

	wg.Wait()

	sort.SliceStable(solutions, func(i, j int) bool {
		if solutions[i].day != solutions[j].day {
			return solutions[i].day < solutions[j].day
		}
		return solutions[i].part < solutions[j].part
	})

	for _, s := range solutions {
		fmt.Printf("Solution for day %02d, part %d: %s\n", s.day, s.part, s.answer)
	}
}

type Solution struct {
	day    int
	part   int
	answer string
}

func solveOne(day int, part int, ch chan Solution) {
	input, err := loadInput(day)
	if err != nil {
		fmt.Printf("Error loading input for day %v\n", day)
		log.Fatal(err)
	}

	solution, ok := adventofcode.SolveDay(day, part, input)
	if !ok {
		fmt.Printf("No solution for Day %d, Part %d\n", day, part)
		os.Exit(1)
	}

	ch <- Solution{day, part, solution}
	err = saveOutput(day, part, solution)

	if err != nil {
		fmt.Printf("Error saving solution to day %d, part %d\n", day, part)
		log.Fatal(err)
	}

	wg.Done()
}

func loadInput(day int) (string, error) {
	name := fmt.Sprintf("%02d.txt", day)
	data, err := os.ReadFile(path.Join(INPUT_PATH, name))
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(data), "\n"), nil
}

func saveOutput(day int, part int, solution string) error {
	name := fmt.Sprintf("%02d-%d.txt", day, part)
	err := os.MkdirAll(OUTPUT_PATH, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(path.Join(OUTPUT_PATH, name), []byte(solution), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
