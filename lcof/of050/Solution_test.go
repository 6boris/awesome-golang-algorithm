package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func(string) byte

var SolutionFuncList = []SolutionFuncType{
	firstUniqChar,
}

//	test info struct
type Case struct {
	name   string
	inputs string
	expect byte
}

// 	test case
var cases = []Case{
	{
		name:   "TestCase 1",
		inputs: "abaccdeff",
		expect: byte('b'),
	},
	{
		name:   "TestCase 1",
		inputs: "",
		expect: byte(' '),
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
