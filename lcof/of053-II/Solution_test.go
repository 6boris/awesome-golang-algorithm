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
	missingNumber_1,
	missingNumber_2,
	missingNumber_3,
	missingNumber_4,
}

// Test case info struct
type Case struct {
	name   string
	input  []int
	expect int
}

// Test case
var cases = []Case{
	{"TestCase 1", []int{3, 0, 1}, 2},
	{"TestCase 2", []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	{"TestCase 2", []int{0, 1}, 2},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				nums := make([]int, 0)
				for _, v := range c.input {
					nums = append(nums, v)
				}
				got := f(nums)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
