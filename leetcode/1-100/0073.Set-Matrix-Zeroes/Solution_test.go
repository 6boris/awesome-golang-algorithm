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
type SolutionFuncType func([][]int)

var SolutionFuncList = []SolutionFuncType{
	setZeroes,
}

// Test case info struct
type Case struct {
	name   string
	input  [][]int
	expect [][]int
}

// Test case
var cases = []Case{
	{name: "TestCase 1", input: [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}, expect: [][]int{
		{1, 0, 1},
		{0, 0, 0},
		{1, 0, 1},
	}},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				f(c.input)
				ast.Equal(c.expect, c.input,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
