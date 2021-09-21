package Solution

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Solution func Info
type SolutionFuncType func(root, p, q *TreeNode) *TreeNode

var SolutionFuncList = []SolutionFuncType{
	lowestCommonAncestor_1,
	lowestCommonAncestor_2,
}

// Test case info struct
type Case struct {
	name               string
	root, q, p, expect *TreeNode
}

var baseTree = &TreeNode{
	Val: 3,
	Left: &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: 4},
		},
	},
	Right: &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{Val: 8},
	},
}

// Test case
var cases = []Case{
	{name: "TestCase 1", root: baseTree, p: baseTree.Left, q: baseTree.Right, expect: baseTree},
	{name: "TestCase 2", root: baseTree, p: baseTree.Left, q: baseTree.Left.Right.Right, expect: baseTree.Left},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.root, c.p, c.q)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
