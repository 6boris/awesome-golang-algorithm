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
type SolutionFuncType func([]int) int

var SolutionFuncList = []SolutionFuncType{
	singleNumber_1,
	singleNumber_2,
	singleNumber_3,
}

//	test info struct
type Case struct {
	name   string
	inputs []int
	expect int
}

// 	test case
var cases = []Case{
	{name: "TestCase 1", inputs: []int{2, 2, 1}, expect: 1},
	{name: "TestCase 2", inputs: []int{4, 1, 2, 1, 2}, expect: 4},
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
