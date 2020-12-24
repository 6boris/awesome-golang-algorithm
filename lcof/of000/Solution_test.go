package Solution

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"testing"
)

//	solution func Info
type SolutionFuncType func(bool) bool

var SolutionFuncList = []SolutionFuncType{
	Solution,
}

//	test info struct
type Case struct {
	name   string
	inputs bool
	expect bool
}

// 	test case
var cases = []Case{
	{
		name:   "TestCase 1",
		inputs: true,
		expect: true,
	},
	{
		name:   "TestCase 2",
		inputs: false,
		expect: false,
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.inputs)
				ast.Equal(c, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
