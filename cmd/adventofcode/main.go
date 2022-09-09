package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/rxedu/adventofcode-2017-go"
)

const INPUT_PATH = "input"
const OUTPUT_PATH = "solutions"

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

	if len(os.Args) == 2 {
		solveOne(day, 1)
		solveOne(day, 2)
		return
	}

	if len(os.Args) > 2 {
		solveOne(day, part)
		return
	}

	for i := 1; i <= adventofcode.NumSolvedDays(); i++ {
		solveOne(i, 1)
		solveOne(i, 2)
	}
}

func solveOne(day int, part int) {
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

	fmt.Printf("Solution for day %02d, part %d: %s\n", day, part, solution)
	err = saveOutput(day, part, solution)

	if err != nil {
		fmt.Printf("Error saving solution to day %d, part %d\n", day, part)
		log.Fatal(err)
	}
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
