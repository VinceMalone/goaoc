package day5

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
		in   string
		want int
	}{
		{"0\n3\n0\n1\n-3", 5},
		{getInput(), 351282},
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
		in   string
		want int
	}{
		{"0\n3\n0\n1\n-3", 10},
		{getInput(), 24568703},
	}
	for _, c := range cases {
		got := Part2(c.in)
		if got != c.want {
			t.Errorf("Part2(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
