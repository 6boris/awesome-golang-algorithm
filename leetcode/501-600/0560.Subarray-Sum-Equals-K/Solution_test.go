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
	subarraySum,
	subarraySum_2,
}

// Test case info struct
type Case struct {
	name   string
	nums   []int
	k      int
	expect int
}

// Test case
var cases = []Case{
	{name: "TestCase 1", nums: []int{1, 1, 1}, k: 2, expect: 2},
	{name: "TestCase 2", nums: []int{1, 2, 3}, k: 3, expect: 2},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.nums, c.k)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
