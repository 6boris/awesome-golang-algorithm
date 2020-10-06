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
		inputs [][]int
		expect []int
	}{
		{"TestCase",
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			[]int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10}},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := printMatrix(c.inputs)
			ast.Equal(c.expect, actual, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.inputs)
		})
	}
}

//	Benchmark Test
func BenchmarkSolution(b *testing.B) {

}

//	Example Test
func ExampleSolution() {

}
