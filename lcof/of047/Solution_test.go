package Solution

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"testing"
)

//	solution func Info
type SolutionFuncType func([][]int) int

var SolutionFuncList = []SolutionFuncType{
	maxValue,
}

//	test info struct
type Case struct {
	name   string
	inputs [][]int
	expect int
}

// 	test case
var cases = []Case{
	{
		name: "TestCase 1",
		inputs: [][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		},
		expect: 12,
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
