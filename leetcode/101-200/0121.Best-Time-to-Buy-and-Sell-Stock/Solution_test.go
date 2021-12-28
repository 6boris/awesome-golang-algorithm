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
type SolutionFuncType func([]int) int

var SolutionFuncList = []SolutionFuncType{
	maxProfit_1,
	maxProfit_2,
	maxProfit_3,
}

// Test case info struct
type Case struct {
	name   string
	inputs []int
	expect int
}

// Test case
var cases = []Case{
	{"TestCase", []int{7, 1, 5, 3, 6, 4}, 5},
	{"TestCase", []int{1, 2, 3, 4, 5}, 4},
	{"TestCase", []int{7, 6, 4, 3, 1}, 0},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.inputs)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
