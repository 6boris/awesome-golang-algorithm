package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func([][]int, int) bool

var SolutionFuncList = []SolutionFuncType{
	findNumberIn2DArray,
}

//	test info struct
type Case struct {
	name   string
	board  [][]int
	target int
	expect bool
}

// 	test case
var cases = []Case{
	{
		name: "TestCase 1",
		board: [][]int{
			{1, 2, 8, 9},
		},
		target: 7,
		expect: false,
	},
	{
		name: "TestCase 2",
		board: [][]int{
			{1, 2, 8, 9},
			{2, 4, 9, 12},
			{4, 7, 10, 13},
			{6, 8, 11, 15},
		},
		target: 7,
		expect: true,
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.board, c.target)
				ast.Equal(c.expect, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
