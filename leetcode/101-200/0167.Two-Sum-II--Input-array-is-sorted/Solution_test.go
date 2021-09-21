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
type SolutionFuncType func([]int, int) []int

var SolutionFuncList = []SolutionFuncType{
	twoSum_1,
	twoSum_2,
}

// Test case info struct
type Case struct {
	name    string
	numbers []int
	target  int
	expect  []int
}

// Test case
var cases = []Case{
	{"TestCase", []int{2, 7, 11, 15}, 9, []int{1, 2}},
	{"TestCase", []int{2, 7, 11, 15}, 22, []int{2, 4}},
	{"TestCase", []int{2, 7, 11, 15}, 26, []int{3, 4}},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.numbers, c.target)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
