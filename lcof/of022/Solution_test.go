package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func(*ListNode, int) *ListNode

var SolutionFuncList = []SolutionFuncType{
	getKthFromEnd_1,
	getKthFromEnd_2,
}

//	test info struct
type Case struct {
	name   string
	inputs *ListNode
	val    int
	expect *ListNode
}

// 	test case
var cases = []Case{
	{
		name:   "TestCase 1",
		inputs: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9}}}},
		val:    2,
		expect: &ListNode{Val: 1},
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.inputs, c.val)
				ast.Equal(c.expect.Val, actual.Val,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
