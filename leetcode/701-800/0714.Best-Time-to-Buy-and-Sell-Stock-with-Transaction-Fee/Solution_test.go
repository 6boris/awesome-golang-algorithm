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
type SolutionFuncType func([]int, int) int

var SolutionFuncList = []SolutionFuncType{
	maxProfit_1,
	maxProfit_2,
}

// Test case info struct
type Case struct {
	name   string
	prices []int
	fee    int
	expect int
}

// Test case
var cases = []Case{
	{"TestCase", []int{1, 3, 2, 8, 4, 9}, 2, 8},
	{"TestCase", []int{1, 3, 7, 5, 10, 3}, 3, 6},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.prices, c.fee)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
