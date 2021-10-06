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
type SolutionFuncType func([]int, []int) *TreeNode

var SolutionFuncList = []SolutionFuncType{
	buildTree_1,
	buildTree_2,
}

// Test case info struct
type Case struct {
	name     string
	preorder []int
	inorder  []int
	expect   *TreeNode
}

// Test case
var cases = []Case{
	{
		name:     "TestCase 1",
		preorder: []int{3, 9, 20, 15, 7},
		inorder:  []int{9, 3, 15, 20, 7},
		expect: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7}},
		},
	},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.preorder, c.inorder)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
