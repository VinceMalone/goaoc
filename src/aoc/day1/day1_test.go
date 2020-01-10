package day1

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
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
		{getInput(), 1393},
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
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
		{getInput(), 1292},
	}
	for _, c := range cases {
		got := Part2(c.in)
		if got != c.want {
			t.Errorf("Part2(%q) == %d, want %d", c.in, got, c.want)
		}
	}
}
