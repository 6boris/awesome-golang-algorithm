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
type SolutionFuncType func(x int) string

var SolutionFuncList = []SolutionFuncType{
	Solution,
}

// Test case info struct
type Case struct {
	name   string
	input  int
	expect string
}

// Test case
var cases = []Case{
	{name: "TestCase 1", input: 2, expect: "0102"},
	{name: "TestCase 2", input: 5, expect: "0102030405"},
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
