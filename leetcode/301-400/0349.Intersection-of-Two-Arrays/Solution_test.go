package Solution

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func([]int, []int) []int

var SolutionFuncList = []SolutionFuncType{
	intersection_1,
	intersection_2,
}

//	test info struct
type Case struct {
	name   string
	num1   []int
	num2   []int
	expect []int
}

// 	test case
var cases = []Case{
	{name: "TestCase 1", num1: []int{1, 2, 2, 1}, num2: []int{2, 2}, expect: []int{2}},
	{name: "TestCase 2", num1: []int{4, 9, 5}, num2: []int{9, 4, 9, 8, 4}, expect: []int{9, 4}},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.num1, c.num2)
				sort.Ints(got)
				sort.Ints(c.expect)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
