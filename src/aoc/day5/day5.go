package day5

import (
	"strconv"
	"strings"
)

type iInstructions interface {
	get(key int) int
	set(key int, value int)
	size() int
}

type iInstructions2 interface {
	iInstructions
	update(key int)
}

func countSteps(i iInstructions2) int {
	index := 0
	steps := 0
	for {
		steps++
		move := i.get(index)
		i.update(index)
		index += move
		if index >= i.size() {
			break
		}
	}
	return steps
}

// ################################################################################

type instructions []int

func newInstructions(input string) instructions {
	strs := strings.Split(input, "\n")
	ints := make([]int, len(strs))
	for i, str := range strs {
		ints[i], _ = strconv.Atoi(str)
	}
	return ints
}

func (i instructions) get(key int) int {
	return i[key]
}

func (i instructions) set(key int, value int) {
	i[key] = value
}

func (i instructions) size() int {
	return len(i)
}

// ################################################################################

// type part1Instructions instructions

// func (i part1Instructions) update(key int) {
// 	value := i.get(key)
// 	i.set(key, value+1)
// }

// type part2Instructions instructions

// func (i part2Instructions) update(index int) {
// 	value := i.get(key)
// 	if value > 2 {
// 		value--
// 	} else {
// 		value++
// 	}
// 	i.set(key, value)
// }

// ################################################################################

// type instructions interface {
// 	update(index int)
// }

// type part1Instructions []int

// func (i part1Instructions) update(index int) {
// 	i[index]++
// }

// type part2Instructions []int

// func (i part2Instructions) update(index int) {
// 	if i[index] > 2 {
// 		i[index]--
// 	} else {
// 		i[index]++
// 	}
// }

// ################################################################################

// func toInts(strs []string) []int {
// 	ints := make([]int, len(strs))
// 	for i, str := range strs {
// 		ints[i], _ = strconv.Atoi(str)
// 	}
// 	return ints
// }

// func newPart1Instructions(input string) part1Instructions {
// 	return newInstructions(input)
// }

// func newPart2Instructions(input string) part2Instructions {
// 	return newInstructions(input)
// }

// ################################################################################

func Part1(input string) int {
	i := newInstructions(input)

	// var i part1Instructions = toInts(strings.Split(input, "\n"))
	// i := newPart1Instructions(input)

	return countSteps(i)
}

func Part2(input string) int {
	i := newInstructions(input)

	// var i part2Instructions = toInts(strings.Split(input, "\n"))
	// i := newPart2Instructions(input)

	return countSteps(i)
}
