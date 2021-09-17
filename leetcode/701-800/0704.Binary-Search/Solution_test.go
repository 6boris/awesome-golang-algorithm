package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Solution func Info
type SolutionFuncType func(nums []int, target int) int

var SolutionFuncList = []SolutionFuncType{
	search_1,
}

// Test case info struct
type Case struct {
	name   string
	nums   []int
	target int
	expect int
}

// Test case
var cases = []Case{
	{name: "TestCase 1", nums: []int{-1, 0, 3, 5, 9, 12}, target: 9, expect: 4},
	{name: "TestCase 2", nums: []int{-1, 0, 3, 5, 9, 12}, target: 2, expect: -1},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				got := f(c.nums, c.target)
				ast.Equal(c.expect, got,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
