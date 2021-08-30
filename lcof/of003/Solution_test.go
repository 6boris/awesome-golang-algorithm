package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SolutionFuncType func([]int) int

//	Solution func
var SolutionFuncList = []SolutionFuncType{
	findRepeatNumber,
	findRepeatNumber2,

	findRepeatNumber3,
}

// test  info struct
type Case struct {
	name   string
	inputs []int
	expect int
}

// test case
var cases = []Case{
	{
		name:   "TestCase 1",
		inputs: []int{},
		expect: 0,
	},
	{
		name:   "TestCase 2",
		inputs: []int{2, 1, 1, 0, 2, 5, 3},
		expect: 1,
	},
	{
		name:   "TestCase 2",
		inputs: []int{1, 1, 0, 2, 5, 3},
		expect: 1,
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.inputs)
				ast.Equal(c.expect, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}

}
