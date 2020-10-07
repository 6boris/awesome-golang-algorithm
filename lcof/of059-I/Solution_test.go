package Solution

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"testing"
)

//	solution func Info
type SolutionFuncType func([]int, int) []int

var SolutionFuncList = []SolutionFuncType{
	maxSlidingWindow,
	maxSlidingWindow2,
	maxSlidingWindow3,
}

//	test info struct
type Case struct {
	name   string
	inputs []int
	k      int
	expect []int
}

// 	test case
var cases = []Case{
	{
		name:   "TestCase 1",
		inputs: []int{1, 3, -1, -3, 5, 3, 6, 7},
		k:      3,
		expect: []int{3, 3, 5, 5, 6, 7},
	},
	{
		name:   "TestCase 2",
		inputs: []int{7, 2, 4, 6},
		k:      2,
		expect: []int{7, 4, 6},
	},
	{
		name:   "TestCase 3",
		inputs: []int{7, 2, 4, 6},
		k:      1,
		expect: []int{7, 2, 4, 6},
	},
	{
		name:   "TestCase 3",
		inputs: []int{},
		k:      1,
		expect: []int{},
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.inputs, c.k)
				ast.Equal(c.expect, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
