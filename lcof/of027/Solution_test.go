package Solution

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"testing"
)

//	solution func Info
type SolutionFuncType func(*TreeNode) *TreeNode

var SolutionFuncList = []SolutionFuncType{
	mirrorTree,
}

//	test info struct
type Case struct {
	name   string
	root   *TreeNode
	expect *TreeNode
}

// 	test case
var cases = []Case{
	{
		name: "TestCase 1",
		root: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 0},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 5},
				},
			},
			Right: &TreeNode{
				Val:   8,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 9},
			},
		},
		expect: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   8,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 7},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{Val: 0},
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
				actual := f(c.root)
				ast.Equal(c.expect, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
