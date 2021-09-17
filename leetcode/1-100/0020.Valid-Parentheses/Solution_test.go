package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Solution func Info
type SolutionFuncType func(x string) bool

var SolutionFuncList = []SolutionFuncType{
	isValid_1,
	isValid_2,
}

// Test case info struct
type Case struct {
	name   string
	input  string
	expect bool
}

// Test case
var cases = []Case{
	{name: "TestCase 1", input: "()", expect: true},
	{name: "TestCase 2", input: "()[]{}", expect: true},
	{name: "TestCase 3", input: "(]", expect: false},
	{name: "TestCase  4", input: "([)]", expect: false},
	{name: "TestCase 5", input: "{[]}", expect: true},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				got := f(c.input)
				ast.Equal(c.expect, got,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
