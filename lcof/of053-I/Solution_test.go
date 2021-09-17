package Solution

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func([]int, int) int

var SolutionFuncList = []SolutionFuncType{
	search,
}

//	Test info struct
type Case struct {
	name   string
	nums   []int
	target int
	expect int
}

// 	test case
var cases = []Case{
	{name: "TestCase 1", nums: []int{5, 7, 7, 8, 8, 10}, target: 8, expect: 2},
	{name: "TestCase 2", nums: []int{5, 7, 7, 8, 8, 10}, target: 6, expect: 0},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.nums, c.target)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
