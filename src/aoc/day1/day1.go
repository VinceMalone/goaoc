package day1

import "strconv"

func sum(input string, getNextIndex func(size, index int) int) int {
	size := len(input)
	sum := 0

	for index := 0; index < len(input); index++ {
		nextIndex := getNextIndex(size, index)

		next := input[nextIndex]
		current := input[index]

		if current == next {
			currentN, _ := strconv.Atoi(string(current))
			sum += currentN
		}
	}

	return sum
}

func Part1(input string) int {
	return sum(input, func(size, index int) int {
		if index == size - 1 {
			return 0
		}
		return index + 1
	})
}

func Part2(input string) int {
	return sum(input, func(size, index int) int {
		return (index + size / 2) % size
	})
}
