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
type SolutionFuncType func(int) int

var SolutionFuncList = []SolutionFuncType{
	fib_1,
	fib_2,
	fib_3,
	fib_4,
}

//	Test info struct
type Case struct {
	name   string
	input  int
	expect int
}

// 	test case
var cases = []Case{
	{"TestCase 1", 2, 1},
	{"TestCase 2", 3, 2},
	{"TestCase 3", 4, 3},
	{"TestCase 4", 0, 0},
	{"TestCase 5", 30, 832040},
}

// TestSolution Example for solution test cases
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
