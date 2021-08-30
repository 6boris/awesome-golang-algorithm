package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func(m int, n int, k int) int

var SolutionFuncList = []SolutionFuncType{
	movingCount1,
	movingCount2,
}

//	test info struct
type Case struct {
	name    string
	input_m int
	input_n int
	input_k int
	expect  int
}

// 	test case
var cases = []Case{
	{
		name:    "TestCase 1",
		input_m: 2,
		input_n: 3,
		input_k: 1,
		expect:  3,
	},
	{
		name:    "TestCase 2",
		input_m: 3,
		input_n: 1,
		input_k: 0,
		expect:  1,
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.input_m, c.input_n, c.input_k)
				ast.Equal(c.expect, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
