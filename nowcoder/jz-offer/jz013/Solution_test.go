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
		inputs []int
		expect []int
	}{
		{"TestCase", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 3, 5, 7, 2, 4, 6}},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := reOrderArray(c.inputs)
			ast.Equal(c.expect, actual, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.inputs)
		})
	}
}

// TestSolution Example for solution test cases
func TestSolution2(t *testing.T) {
	ast := assert.New(t)
	//	test cases
	cases := []struct {
		name   string
		inputs []int
		expect []int
	}{
		{"TestCase", []int{1, 2, 3, 4, 5, 6, 7}, []int{1, 3, 5, 7, 2, 4, 6}},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := reOrderArray2(c.inputs)
			ast.Equal(c.expect, actual, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.inputs)
		})
	}
}
