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
	list, _, _ := knotHash(input, size, 0, 0)
	return list[0] * list[1]
}

func solvePartTwo(input []int) string {
	size := 256
	rounds := 64
	cur := 0
	skip := 0

	var sparseHash []int
	for i := 0; i < rounds; i++ {
		sparseHash, cur, skip = knotHash(input, size, cur, skip)
	}

	hash := denseHash(sparseHash)
	return toHexString(hash)
}

func knotHash(lengths []int, size int, cur int, skip int) ([]int, int, int) {
	list := makeList(size)

	for _, length := range lengths {
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

func denseHash(sparseHash []int) []int {
	chunkSize := 16
	numChunks := len(sparseHash) / chunkSize
	hash := make([]int, numChunks)
	for i := 0; i < numChunks; i++ {
		start := chunkSize * i
		end := start + chunkSize
		for _, v := range sparseHash[start:end] {
			hash[i] = v ^ hash[i]
		}
	}
	return hash
}

func toHexString(denseHash []int) string {
	var hash string
	for _, v := range denseHash {
		hash += fmt.Sprintf("%02x", v)
	}
	return hash
}

func makeList(size int) []int {
	list := make([]int, size)
	for i := 0; i < len(list); i++ {
		list[i] = i
	}
	return list
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
