package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func(uint32) int

var SolutionFuncList = []SolutionFuncType{
	hammingWeight,
}

//	test info struct
type Case struct {
	name   string
	inputs uint32
	expect int
}

// 	test case
var cases = []Case{
	{
		name:   "TestCase 1",
		inputs: uint32(0o0000000000000000000000000001011),
		expect: 3,
	},
	{
		name:   "TestCase 2",
		inputs: uint32(0o0000000000000000000000010000000),
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
