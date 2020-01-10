package day4

import (
	"sort"
	"strings"

	"./set"
)

func sum(input string, validate func(phrase string) bool) int {
	phrases := strings.Split(input, "\n")
	sum := len(phrases)
	for _, phrase := range phrases {
		valid := validate(phrase)
		if !valid {
			sum--
		}
	}
	return sum
}

func validate(hash func(value string) string) func(phrase string) bool {
	return func(phrase string) bool {
		wordSet := set.NewSet()
		words := strings.Fields(phrase)
		for _, word := range words {
			key := hash(word)
			if wordSet.Contains(key) {
				return false
			}
			wordSet.Add(key)
		}
		return true
	}
}

func Part1(input string) int {
	return sum(input, validate(func(value string) string {
		return value
	}))
}

func Part2(input string) int {
	return sum(input, validate(func(value string) string {
		letters := strings.Split(value, "")
		sort.Strings(letters)
		return strings.Join(letters, "")
	}))
}
