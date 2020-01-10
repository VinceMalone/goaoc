package day2

import "strconv"
import "strings"

func toInts(strs []string) []int {
	ints := make([]int, len(strs))
	for i, str := range strs {
		ints[i], _ = strconv.Atoi(str)
	}
	return ints
}

func sum(input string, getRowResult func(nums []int) int) int {
	sum := 0
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		nums := toInts(strings.Fields(row))
		sum += getRowResult(nums)
	}
	return sum
}

func Part1(input string) int {
	return sum(input, func(nums []int) int {
		largest := nums[0]
		smallest := nums[0]
		for _, num := range nums {
			if num > largest {
				largest = num
			}
			if num < smallest {
				smallest = num
			}
		}
		return largest - smallest
	})
}

func Part2(input string) int {
	return sum(input, func(nums []int) int {
		for numeratorIndex, numerator := range nums {
			for denominatorIndex, denominator := range nums {
				if numeratorIndex != denominatorIndex && numerator % denominator == 0 {
					return numerator / denominator
				}
			}
		}
		return 0
	})
}
