package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func(*ListNode, *ListNode) *ListNode

var SolutionFuncList = []SolutionFuncType{
	mergeTwoLists,
	// mergeTwoLists2,
}

//	test info struct
type Case struct {
	name   string
	l1     *ListNode
	l2     *ListNode
	expect *ListNode
}

// 	test case
var cases = []Case{
	{
		name:   "TestCase 1",
		l1:     &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}},
		l2:     &ListNode{Val: 2, Next: &ListNode{Val: 4}},
		expect: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.l1, c.l2)
				ast.Equal(c.expect, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
