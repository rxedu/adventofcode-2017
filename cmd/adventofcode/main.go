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
	if len(os.Args) < 2 {
		fmt.Printf("Error: must pass day to solve as first argument\n")
		os.Exit(2)
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

	day, err := strconv.Atoi(os.Args[1])
	if err != nil || day < 1 {
		fmt.Printf("Day must be positive integer, got %v\n", day)
		os.Exit(2)
	}

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

	fmt.Printf("Solution for day %02d, part %d: %s", day, part, solution)
	err = saveOutput(day, solution)

	if err != nil {
		fmt.Printf("Error saving solution to day %d, part %d\n", day, part)
		log.Fatal(err)
	}
}

func loadInput(day int) (string, error) {
	name := getName(day)
	data, err := os.ReadFile(path.Join(INPUT_PATH, name))
	if err != nil {
		return "", err
	}

	return strings.TrimSuffix(string(data), "\n"), nil
}

func saveOutput(day int, solution string) error {
	name := getName(day)
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

func getName(day int) string {
	return fmt.Sprintf("%02d.txt", day)
}
