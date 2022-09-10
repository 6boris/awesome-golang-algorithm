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
type SolutionFuncType func(int, []int) int

var SolutionFuncList = []SolutionFuncType{
	maxProfit_1,
	maxProfit_2,
	maxProfit_3,
}

// Test case info struct
type Case struct {
	name   string
	prices []int
	k      int
	expect int
}

// Test case
var cases = []Case{
	{"TestCase 1", []int{2, 4, 1}, 2, 2},
	{"TestCase 2", []int{3, 2, 6, 5, 0, 3}, 2, 7},
	{"TestCase 3", []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0}, 2, 13},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.k, c.prices)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
