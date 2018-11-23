package Solution

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		input1 *TreeNode
		input2 *TreeNode
		expect bool
	}{
		{"TestCacse 1",
			&TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
				Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
			&TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
				Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
			true},
		{"TestCacse 2",
			&TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
				Right: nil},
			&TreeNode{Val: 1,
				Left:  nil,
				Right: &TreeNode{Val: 2, Left: nil, Right: nil}},
			true},
		{"TestCacse 3",
			&TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
				Right: nil},
			&TreeNode{Val: 1,
				Left:  nil,
				Right: &TreeNode{Val: 2, Left: nil, Right: nil}},
			true},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isSameTree(c.input1, c.input2)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.input1)
			}
		})
	}
}

func TestIsSameTree(t *testing.T) {
	t.Run("test tree seilize", func(t *testing.T) {
		tree := conFromPreStr("124##5##36##")

		fmt.Println(dumpTreeToString(tree))
	})
}

//	序列化
func dumpTreeToString(tree *TreeNode) string {
	var str string = ""
	dumpTreeToStringHelper(tree, &str)
	return str
}

func dumpTreeToStringHelper(tree *TreeNode, str *string) {
	if tree == nil {
		*str += "#"
		return
	}
	*str += strconv.Itoa(tree.Val)
	dumpTreeToStringHelper(tree.Left, str)
	dumpTreeToStringHelper(tree.Right, str)
}

//	反序列化
func conFromPreStr(str string) *TreeNode {
	var strIndex int = 0
	return conFromPreStrHelper(str, &strIndex)
}

func conFromPreStrHelper(str string, index *int) *TreeNode {
	if len(str) == *index {
		return nil
	}
	if (str)[*index] == '#' {
		(*index) = *index + 1
		return nil
	}
	v, _ := strconv.Atoi(string((str)[*index]))
	root := &TreeNode{Val: v}
	(*index) = *index + 1
	root.Left = conFromPreStrHelper(str, index)
	root.Right = conFromPreStrHelper(str, index)

	return root
}
