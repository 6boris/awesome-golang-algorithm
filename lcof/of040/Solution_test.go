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
type SolutionFuncType func([]int, int) []int

var SolutionFuncList = []SolutionFuncType{
	getLeastNumbers,
}

//	test info struct
type Case struct {
	name   string
	arr    []int
	k      int
	expect []int
}

// 	Test case
var cases = []Case{
	{name: "TestCase", arr: []int{3, 2, 1}, k: 2, expect: []int{1, 2}},
	{name: "TestCase", arr: []int{0, 1, 2, 1}, k: 1, expect: []int{0}},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.arr, c.k)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
