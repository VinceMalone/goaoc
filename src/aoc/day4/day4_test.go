package day4

import "io/ioutil"
import "strings"
import "testing"

func getInput() string {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data))
}

func TestPart1(t *testing.T) {
	cases := []struct {
		in string
		want int
	}{
		{"aa bb cc dd ee\naa bb cc dd aa\naa bb cc dd aaa", 2},
		{getInput(), 383},
	}
	for _, c := range cases {
		got := Part1(c.in)
		if got != c.want {
			t.Errorf("Part1(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}

func TestPart2(t *testing.T) {
	cases := []struct {
		in string
		want int
	}{
		{"abcde fghij\nabcde xyz ecdab\na ab abc abd abf abj\niiii oiii ooii oooi oooo\noiii ioii iioi iiio", 3},
		{getInput(), 265},
	}
	for _, c := range cases {
		got := Part2(c.in)
		if got != c.want {
			t.Errorf("Part2(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
