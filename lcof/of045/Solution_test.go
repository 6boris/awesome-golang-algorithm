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
type SolutionFuncType func([]int) string

var SolutionFuncList = []SolutionFuncType{
	minNumber,
}

//	test info struct
type Case struct {
	name   string
	inputs []int
	expect string
}

// 	Test case
var cases = []Case{
	{name: "TestCase 1", inputs: []int{10, 2}, expect: "102"},
	{name: "TestCase 2", inputs: []int{3, 30, 34, 5, 9}, expect: "3033459"},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.inputs)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
