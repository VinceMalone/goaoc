package memory

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/go-cmp/cmp"
)

// State represents a list of banks
type State []int

func (s State) findMost() (int, int) {
	index := 0
	most := s[index]
	for i, blocks := range s {
		if blocks > most {
			index = i
			most = blocks
		}
	}
	return index, most
}

func (s State) hash() string {
	// TODO
}

// NewState creates a new State from an input string
func NewState(input string) State {
	fields := strings.Fields(input)
	banks := make([]int, len(fields))
	for i, field := range fields {
		banks[i], _ = strconv.Atoi(field)
	}
	return banks
}

// Reallocate
func Reallocate(blocks State) int {
	cycles := 0
	set := make(map[string]struct{})
	i, most := blocks.findMost()
	fmt.Println(i, most)

	a := State{0, 1, 2}
	b := State{0, 1, 2}
	fmt.Println(cmp.Equal(a, b))

	return cycles
}
