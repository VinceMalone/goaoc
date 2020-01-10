package memory

import "testing"

var cases = []struct {
	in   State
	want int
}{
	{State{0, 2, 7, 0}, 5},
}

func TestReallocate(t *testing.T) {
	for _, c := range cases {
		if got := Reallocate(c.in); got != c.want {
			t.Errorf("test failed for %v; wanted:%v but got:%v", c.in, c.want, got)
		}
	}
}
