package Solution

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Solution func Info
type SolutionFuncType func(bool) bool

var SolutionFuncList = []SolutionFuncType{
	Solution_1,
	Solution_2,
}

// Test case info struct
type Case struct {
	name   string
	input  bool
	expect bool
}

// Test case
var cases = []Case{
	{name: "TestCase 1", input: true, expect: true},
	{name: "TestCase 2", input: false, expect: false},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.input)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
