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
		inputs []byte
		expect []byte
	}{
		{"TestCase",
			[]byte{'a', ' ', 'b', ' ', 'c'},
			[]byte{'a', '%', '2', '0', 'b', '%', '2', '0', 'c'}},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := replaceSpace(c.inputs)
			ast.Equal(actual, c.expect, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.inputs)
		})
	}
}
