package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Solution func Info
type SolutionFuncType func(x, y string) string

var SolutionFuncList = []SolutionFuncType{
	// addStrings_1,
	addStrings_2,
}

// Test case info struct
type Case struct {
	name       string
	num1, num2 string
	expect     string
}

// Test case
var cases = []Case{
	{"TestCase1", "11", "123", "134"},
	{"TestCase2", "456", "77", "533"},
	{"TestCase3", "0", "0", "0"},
	{"TestCase4", "86043", "5582", "91625"},
	{"TestCase5", "51189", "967895", "1019084"},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				got := f(c.num1, c.num2)
				ast.Equal(c.expect, got,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
