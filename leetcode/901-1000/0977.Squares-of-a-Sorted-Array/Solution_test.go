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
type SolutionFuncType func([]int) []int

var SolutionFuncList = []SolutionFuncType{
	sortedSquares_1,
	sortedSquares_2,
}

// Test case info struct
type Case struct {
	name   string
	inputs []int
	expect []int
}

// Test case
var cases = []Case{
	{"TestCase 1", []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
	{"TestCase 2", []int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	{"TestCase 3", []int{}, []int{}},
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
