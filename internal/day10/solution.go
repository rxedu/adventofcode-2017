package day10

import (
	"strconv"
	"strings"
)

func SolvePartOne(input string) string {
	return serialize(solvePartOne(parse(input), 256))
}

func SolvePartTwo(input string) string {
	return serialize(solvePartTwo(parse(input)))
}

func solvePartOne(input []int, size int) int {
	list := make([]int, size)
	for i := 0; i < len(list); i++ {
		list[i] = i
	}

	cur := 0
	skip := 0

	for _, length := range input {
		length = length % size

		target := make([]int, length)
		mid := (cur + length) % size
		if cur <= mid {
			target = list[cur:mid]
			reversed := reverse(target)
			begin := list[:cur]
			end := append(reversed, list[mid:]...)
			list = append(begin, end...)
		} else {
			target = append(list[cur:], list[:mid]...)
			reversed := reverse(target)

			i := len(list[cur:])
			head := reversed[i:]
			tail := reversed[:i]

			begin := append(head, list[mid:cur]...)
			end := tail
			list = append(begin, end...)
		}

		cur += length + skip
		cur = cur % len(list)
		skip++
	}

	return list[0] * list[1]
}

func solvePartTwo(input []int) int {
	return 1
}

func reverse(list []int) []int {
	reversed := make([]int, len(list))
	for i, v := range list {
		reversed[len(list)-i-1] = v
	}
	return reversed
}

func parse(input string) []int {
	elements := strings.Split(input, ",")
	arr := make([]int, len(elements))
	for i, str := range elements {
		v, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		arr[i] = v
	}
	return arr
}

func serialize(output int) string {
	return strconv.Itoa(output)
}
