package Solution

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Solution func Info
type SolutionFuncType func(root *TreeNode) []int

var SolutionFuncList = []SolutionFuncType{
	inorderTraversal,
	inorderTraversal_2,
	inorderTraversal_3,
}

// Test case info struct
type Case struct {
	name   string
	input  *TreeNode
	expect []int
}

// Test case
var cases = []Case{
	{
		name: "TestCase 1",
		input: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: nil,
			},
		},
		expect: []int{1, 3, 2},
	},
	{name: "TestCase 2", input: &TreeNode{Val: 1}, expect: []int{1}},
	{name: "TestCase 3", input: nil, expect: []int{}},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)
	fmt.Println(2<<30 - 1)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				got := f(c.input)
				ast.Equal(c.expect, got,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
