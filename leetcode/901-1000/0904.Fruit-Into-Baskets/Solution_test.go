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
	totalFruit_1,
}

// Test case info struct
type Case struct {
	name   string
	input  []int
	expect int
}

// Test case
var cases = []Case{
	{name: "TestCase 1", input: []int{1, 2, 1}, expect: 3},
	{name: "TestCase 2", input: []int{0, 1, 2, 2}, expect: 3},
	{name: "TestCase 3", input: []int{1, 2, 3, 2, 2}, expect: 4},
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
