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
		inputs *NodeList
		expect []int
	}{
		{"TestCase", &NodeList{Val: 1,
			Next: &NodeList{Val: 2, Next: &NodeList{Val: 3}}},
			[]int{3, 2, 1},
		}}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := printListFromTailToHead(c.inputs)

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
		inputs *NodeList
		expect []int
	}{
		{"TestCase", &NodeList{Val: 1,
			Next: &NodeList{Val: 2, Next: &NodeList{Val: 3}}},
			[]int{3, 2, 1},
		}}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := printListFromTailToHead2(c.inputs)

			ast.Equal(c.expect, actual, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.inputs)
		})
	}
}
