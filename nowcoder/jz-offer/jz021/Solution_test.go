package Solution

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)
	//	test cases
	cases := []struct {
		name   string
		pushA  []int
		popA   []int
		expect bool
	}{
		{"TestCase", []int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, true},
		{"TestCase", []int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, true},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := validateStackSequences(c.pushA, c.popA)
			ast.Equal(c.expect, actual, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.pushA, c.popA)
		})
	}
}
