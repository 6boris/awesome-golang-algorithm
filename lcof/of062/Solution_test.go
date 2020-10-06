package Solution

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"testing"
)

//	solution func Info
type SolutionFuncType func(int, int) int

var SolutionFuncList = []SolutionFuncType{
	lastRemaining,
}

//	test info struct
type Case struct {
	name   string
	inputs []int
	expect int
}

// 	test case
var cases = []Case{
	{
		name:   "TestCase 1",
		inputs: []int{5, 3},
		expect: 3,
	},
	{
		name:   "TestCase 2",
		inputs: []int{10, 17},
		expect: 2,
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.inputs[0], c.inputs[1])
				ast.Equal(c.expect, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
