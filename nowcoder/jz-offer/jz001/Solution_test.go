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
		target int
		expect bool
	}{
		{"TestCase", [][]int{
			[]int{1, 2, 8, 9},
			[]int{2, 4, 9, 12},
			[]int{4, 7, 10, 13},
			[]int{6, 8, 11, 15},
		}, 7, true},
		{"TestCase", [][]int{
			[]int{1, 2, 8, 9},
			[]int{2, 4, 9, 12},
			[]int{4, 6, 10, 13},
			[]int{6, 8, 11, 15},
		}, 7, false},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := Find(c.inputs, c.target)
			ast.Equal(actual, c.expect, "expected: %v, but actual: %v, with inputs: %v",
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
