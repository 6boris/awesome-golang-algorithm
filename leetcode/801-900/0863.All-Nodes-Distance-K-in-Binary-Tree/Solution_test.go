package Solution

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	trees := []*TreeNode{
		{
			Val: 3,
			Left: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 6},
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
		},
		{Val: 1},
	}
	cases := []struct {
		name       string
		tree, node *TreeNode
		k          int
		expect     []int
	}{
		{"TestCase1", trees[0], trees[0].Left, 2, []int{1, 4, 7}},
		{"TestCase2", trees[1], trees[1], 3, []int{}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.tree, c.node, c.k)
			sort.Ints(got)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.tree, c.node, c.k)
			}
		})
	}
}

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
