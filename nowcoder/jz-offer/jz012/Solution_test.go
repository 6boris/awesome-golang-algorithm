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
		b      float64
		n      int
		expect float64
	}{
		{"TestCase", 2, 3, 8},
		{"TestCase", -2, 3, -8},
		{"TestCase", 2, -3, 0.12500},
		{"TestCase", -2, -3, -0.12500},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := Power(c.b, c.n)
			ast.Equal(c.expect, actual, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.b)
		})
	}
}
