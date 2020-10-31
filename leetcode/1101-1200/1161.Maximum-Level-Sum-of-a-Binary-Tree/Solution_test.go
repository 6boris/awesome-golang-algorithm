package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	root *TreeNode
}

func (t *BinaryTree) insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &TreeNode{Val: data, Left: nil, Right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *TreeNode) insert(data int) {
	if n == nil {
		return
	} else if data <= n.Val {
		if n.Left == nil {
			n.Left = &TreeNode{Val: data, Left: nil, Right: nil}
		} else {
			n.Left.insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Val: data, Left: nil, Right: nil}
		} else {
			n.Right.insert(data)
		}
	}
}

func main() {

}

func TestSolution(t *testing.T) {
	//	测试用例
	// The original Problem requires the binary tree constructed from array. Please refer to test cases produced at Leetcode problem https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
	tree1 := &BinaryTree{}
	tree1.insert(1).
		insert(7).
		insert(0).
		insert(7).
		insert(-8)

	cases := []struct {
		name   string
		inputs *BinaryTree
		expect int
	}{
		{"TestCase1", tree1, 2},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {

}

//	使用案列
func ExampleSolution() {

}
