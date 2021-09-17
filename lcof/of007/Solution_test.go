package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func([]int, []int) *TreeNode

var SolutionFuncList = []SolutionFuncType{
	buildTree,
}

//	test info struct
type Case struct {
	name   string
	pre    []int
	in     []int
	expect *TreeNode
}

// 	test case
var cases = []Case{
	{
		name: "TestCase 1",
		pre:  []int{3, 9, 20, 15, 7},
		in:   []int{9, 3, 15, 20, 7},
		expect: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		},
	},
	{
		name: "TestCase 1",
		pre:  []int{1, 2, 4, 7, 3, 5, 6, 8},
		in:   []int{4, 7, 2, 1, 5, 3, 8, 6},
		expect: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			Right: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 5},
				Right: &TreeNode{
					Val:  6,
					Left: &TreeNode{Val: 8},
				},
			},
		},
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.pre, c.in)
				ast.Equal(c.expect, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
