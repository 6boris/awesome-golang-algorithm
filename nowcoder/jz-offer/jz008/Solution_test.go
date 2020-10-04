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
		inputs int
		expect int
	}{
		{"TestCase", 1, 1},
		{"TestCase", 2, 2},
		{"TestCase", 4, 5},
		{"TestCase", 5, 8},
		{"TestCase", 6, 13},
		{"TestCase", 11, 144},
		{"TestCase", 14, 610},
		{"TestCase", 27, 317811},
		{"TestCase", 43, 701408733},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := jumpFloor(c.inputs)
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
		inputs int
		expect int
	}{
		{"TestCase", 1, 1},
		{"TestCase", 2, 2},
		{"TestCase", 4, 5},
		{"TestCase", 5, 8},
		{"TestCase", 6, 13},
		{"TestCase", 11, 144},
		{"TestCase", 14, 610},
		{"TestCase", 27, 317811},
		{"TestCase", 43, 701408733},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := jumpFloor2(c.inputs)
			ast.Equal(c.expect, actual, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.inputs)
		})
	}
}

// TestSolution Example for solution test cases
func TestSolution3(t *testing.T) {
	ast := assert.New(t)
	//	test cases
	cases := []struct {
		name   string
		inputs int
		expect int
	}{
		{"TestCase", 1, 1},
		{"TestCase", 2, 2},
		{"TestCase", 4, 5},
		{"TestCase", 5, 8},
		{"TestCase", 6, 13},
		{"TestCase", 11, 144},
		{"TestCase", 14, 610},
		{"TestCase", 27, 317811},
		{"TestCase", 43, 701408733},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := jumpFloor3(c.inputs)
			ast.Equal(c.expect, actual, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.inputs)
		})
	}
}

// TestSolution Example for solution test cases
func TestSolution4(t *testing.T) {
	ast := assert.New(t)
	//	test cases
	cases := []struct {
		name   string
		inputs int
		expect int
	}{
		{"TestCase", 1, 1},
		{"TestCase", 2, 2},
		{"TestCase", 4, 5},
		{"TestCase", 5, 8},
		{"TestCase", 6, 13},
		{"TestCase", 11, 144},
		{"TestCase", 14, 610},
		{"TestCase", 27, 317811},
		{"TestCase", 43, 701408733},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := jumpFloor4(c.inputs)
			ast.Equal(c.expect, actual, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.inputs)
		})
	}
}
