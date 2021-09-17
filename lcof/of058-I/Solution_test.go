package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func(string) string

var SolutionFuncList = []SolutionFuncType{
	reverseWords,
}

//	test info struct
type Case struct {
	name   string
	inputs string
	expect string
}

// 	test case
var cases = []Case{
	{
		name:   "TestCase 1",
		inputs: "the sky is blue",
		expect: "blue is sky the",
	},

	{
		name:   "前后有多个空格",
		inputs: "  hello world!  ",
		expect: "world! hello",
	},
	{
		name:   "连续多个空格",
		inputs: "a good   example",
		expect: "example good a",
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
