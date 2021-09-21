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
type SolutionFuncType func(root *TreeNode) [][]int

var SolutionFuncList = []SolutionFuncType{
	zigzagLevelOrder_1,
	zigzagLevelOrder_2,
	zigzagLevelOrder_3,
}

// Test case info struct
type Case struct {
	name   string
	input  *TreeNode
	expect [][]int
}

// Test case
var cases = []Case{
	{name: "TestCase 1", input: &TreeNode{Val: 1}, expect: [][]int{{1}}},
	{name: "TestCase 2", input: nil, expect: [][]int{}},
	{
		name: "TestCase 3",
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
		expect: [][]int{{1}, {2}, {3}},
	},
	{
		name: "TestCase 4",
		input: &TreeNode{Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   6,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 7},
			},
		},
		expect: [][]int{{4}, {6, 2}, {1, 3, 5, 7}},
	},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.input)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
