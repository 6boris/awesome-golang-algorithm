package Solution

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"testing"
)

//	solution func Info
type SolutionFuncType func(node *ListNode) []int

var SolutionFuncList = []SolutionFuncType{
	printListFromTailToHead,
	printListFromTailToHead2,
}

//	test info struct
type Case struct {
	name   string
	inputs *ListNode
	expect []int
}

// 	test case
var cases = []Case{
	{
		name: "TestCase 1",
		inputs: &ListNode{Val: 1,
			Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		expect: []int{3, 2, 1},
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
