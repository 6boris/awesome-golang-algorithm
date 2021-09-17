package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func([][]int) []int

var SolutionFuncList = []SolutionFuncType{
	spiralOrder,
}

//	test info struct
type Case struct {
	name   string
	inputs [][]int
	expect []int
}

// 	test case
var cases = []Case{
	{
		name: "TestCase 1",
		inputs: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		expect: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
	},
	{
		name: "TestCase 2",
		inputs: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		expect: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
	},
	{
		name: "TestCase 3",
		inputs: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		},
		expect: []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10},
	},
	{
		name: "TestCase 3",
		inputs: [][]int{
			{1, 2, 3, 4},
		},
		expect: []int{1, 2, 3, 4},
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.inputs)
				ast.Equal(c.expect, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
