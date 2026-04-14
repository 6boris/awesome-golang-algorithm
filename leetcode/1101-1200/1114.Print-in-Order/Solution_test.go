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
type SolutionFuncType func([]int) string

var SolutionFuncList = []SolutionFuncType{
	Solution,
}

// Test case info struct
type Case struct {
	name   string
	input  []int
	expect string
}

// Test case
var cases = []Case{
	{name: "TestCase 1", input: []int{1, 2, 3}, expect: "firstsecondthird"},
	{name: "TestCase 2", input: []int{1, 3, 2}, expect: "firstsecondthird"},
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
