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
	moveZeroes,
	moveZeroes_2,
}

// Test case info struct
type Case struct {
	name   string
	input  []int
	expect []int
}

// Test case
var cases = []Case{
	{"TestCase 1", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	{"TestCase 2", []int{0}, []int{0}},
	{"TestCase 3", []int{1}, []int{1}},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				i := make([]int, 0)
				for _, v := range c.input {
					i = append(i, v)
				}
				got := f(i)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
