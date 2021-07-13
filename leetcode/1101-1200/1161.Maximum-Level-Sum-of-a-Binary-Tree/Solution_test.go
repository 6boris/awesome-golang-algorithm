package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SolutionFuncType func(*TreeNode) int

var SolutionFuncList = []SolutionFuncType{
	maxLevelSum,
	//maxLevelSum2,
}

var DefaultValue int = -1024

func InsertNodeToTree(tree *TreeNode, node *TreeNode) {
	if tree == nil {
		return
	}
	if tree.Val == DefaultValue {
		tree.Val = node.Val
		return
	}
	if node.Val > tree.Val {
		if tree.Right == nil {
			tree.Right = &TreeNode{Val: DefaultValue}
		}
		InsertNodeToTree(tree.Right, node)
	}
	if node.Val < tree.Val {
		if tree.Left == nil {
			tree.Left = &TreeNode{Val: DefaultValue}
		}
		InsertNodeToTree(tree.Left, node)
	}
}

func InitTree(values ...int) (root *TreeNode) {
	rootNode := TreeNode{Val: DefaultValue, Right: nil, Left: nil}
	for _, value := range values {
		node := TreeNode{Val: value}
		InsertNodeToTree(&rootNode, &node)
	}
	return &rootNode
}

func TestSolution(t *testing.T) {
	//	测试用例
	ast := assert.New(t)

	// The original Problem requires the binary tree constructed from array. Please refer to test cases produced at Leetcode problem https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
	//treeNode := InitTree(1, 7, 0, 7, -8)
	treeNode := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 7, Left: &TreeNode{Val: 7, Right: &TreeNode{Val: -8}}},
		Right: &TreeNode{Val: 0},
	}
	var cases = []struct {
		name   string
		inputs *TreeNode
		expect int
	}{
		{"TestCase1", treeNode, 2},
	}
	//ast.Equal(treeNode, &TreeNode{})
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
