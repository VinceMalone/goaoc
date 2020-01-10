package day2

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
		{"5 1 9 5\n7 5 3\n2 4 6 8", 18},
		{getInput(), 47136},
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
		{"5 9 2 8\n9 4 7 3\n3 8 6 5", 9},
		{getInput(), 250},
	}
	for _, c := range cases {
		got := Part2(c.in)
		if got != c.want {
			t.Errorf("Part2(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
