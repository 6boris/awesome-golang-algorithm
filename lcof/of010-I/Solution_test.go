package Solution

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"testing"
)

//	solution func Info
type SolutionFuncType func(int) int

var SolutionFuncList = []SolutionFuncType{
	Fibonacci,
	Fibonacci2,
	Fibonacci3,
}

//	test info struct
type Case struct {
	name   string
	inputs int
	expect int
}

// 	test case
var cases = []Case{
	{
		name:   "TestCase 1",
		inputs: 2,
		expect: 1,
	},
	{
		name:   "TestCase 2",
		inputs: 5,
		expect: 5,
	},
	{
		name:   "TestCase 3",
		inputs: 13,
		expect: 233,
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
