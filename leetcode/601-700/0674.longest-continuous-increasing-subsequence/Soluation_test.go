package Soluation

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
	findLengthOfLCIS_1,
}

// Test case info struct
type Case struct {
	name   string
	inputs []int
	expect int
}

// Test case
var cases = []Case{
	{"TestCase 1", []int{1, 3, 5, 4, 7}, 3},
	{"TestCase 2", []int{2, 2, 2, 2, 2}, 1},
	{"TestCase 3", []int{1}, 1},
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
