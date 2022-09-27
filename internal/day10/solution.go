package day10

import (
	"fmt"
	"strconv"
	"strings"
)

func SolvePartOne(input string) string {
	return serializePartOne(solvePartOne(parsePartOne(input), 256))
}

func SolvePartTwo(input string) string {
	return serializePartTwo(solvePartTwo(parsePartTwo(input)))
}

func solvePartOne(input []int, size int) int {
	list, _, _ := knotHash(size, input, 0, 0)
	return list[0] * list[1]
}

func solvePartTwo(input []int) string {
	size := 256
	chunkSize := 16
	rounds := 64
	cur := 0
	skip := 0
	var sparseHash []int

	for i := 0; i < rounds; i++ {
		sparseHash, cur, skip = knotHash(size, input, cur, skip)
	}

	denseHash := make([]int, size/chunkSize)
	for i := 0; i < chunkSize; i++ {
		for _, v := range sparseHash[i : i+chunkSize] {
			denseHash[i] = v ^ denseHash[i]
		}
	}

	var hash string
	for _, v := range denseHash {
		hash += fmt.Sprintf("%02x", v)
	}
	fmt.Println(sparseHash)
	fmt.Println(denseHash)

	return hash
}

func knotHash(size int, input []int, cur int, skip int) ([]int, int, int) {
	list := make([]int, size)
	for i := 0; i < len(list); i++ {
		list[i] = i
	}

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

	return list, cur, skip
}

func reverse(list []int) []int {
	reversed := make([]int, len(list))
	for i, v := range list {
		reversed[len(list)-i-1] = v
	}
	return reversed
}

func parsePartOne(input string) []int {
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

func parsePartTwo(input string) []int {
	lengthSuffix := []int{17, 31, 73, 47, 23}
	arr := make([]int, len(input))
	for i, str := range input {
		arr[i] = int(str)
	}
	return append(arr, lengthSuffix...)
}

func serializePartOne(output int) string {
	return strconv.Itoa(output)
}

func serializePartTwo(output string) string {
	return output
}
