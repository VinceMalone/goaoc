package stack

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	input    Stack
	resolver OffsetResolver
	expected int
}{
	{Stack{0, 3, 0, 1, -3}, SimpleIncrementer, 5},
	{Stack{0, 2, 0, 1, -3}, SimpleIncrementer, 10},
	{Stack{0, 3, 0, 1, -3}, BiasedDecrementer, 10},
}

func TestTrace(t *testing.T) {
	for _, c := range testCases {
		if actual := Stack(c.input).Trace(c.resolver); actual != c.expected {
			t.Errorf("test failed for %v; wanted:%v but got:%v", c.input, c.expected, actual)
		}
	}
}

func BenchmarkTraceWithSimpleIncrementer(b *testing.B) {
	b.StopTimer()
	for _, c := range testCases {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			Stack(c.input).Trace(SimpleIncrementer)
		}
		b.StopTimer()
	}
}

func BenchmarkTraceWithBiasedDecrementer(b *testing.B) {
	b.StopTimer()
	for _, c := range testCases {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			Stack(c.input).Trace(BiasedDecrementer)
		}
		b.StopTimer()
	}
}

func TestNewStack(t *testing.T) {
	inputCases := []struct {
		input    string
		expected Stack
	}{
		{"0\n2\n2\n-1\n-1\n-4\n-2\n-6", []int{0, 2, 2, -1, -1, -4, -2, -6}},
	}

	for _, c := range inputCases {
		if actual := NewStack(c.input); !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("failed '%v'; expected:%v, got:%v", c.input, c.expected, actual)
		}
	}
}
