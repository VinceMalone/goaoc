package set

import "testing"

func Test_AddSet(t *testing.T) {
	set := NewSet()

	set.Add(1)
	set.Add(2)
	set.Add(3)

	if set.Size() != 3 {
		t.Error("AddSet does not have a size of 3 even though 3 items were added to a new set")
	}
}

func Test_ContainsSet(t *testing.T) {
	set := NewSet()

	set.Add(2)

	if !set.Contains(2) {
		t.Error("ContainsSet should contain 2")
	}

	set.Remove(2)

	if set.Contains(2) {
		t.Error("ContainsSet should not contain 2")
	}

	set.Add(5)
	set.Add(4)
	set.Add(9)

	if !set.Contains(5, 4, 9) {
		t.Error("ContainsSet should contain 5, 4, and 9")
	}
}
