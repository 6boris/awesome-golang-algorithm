package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs *TreeNode
		expect [][]int
	}{
		{"TestCase1", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 6},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 7},
			},
		}, [][]int{
			{4},
			{2},
			{1, 5, 6},
			{3},
			{7},
		}},
		{"TestCase2", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 7},
			},
		}, [][]int{
			{4},
			{2},
			{1, 5, 6},
			{3},
			{7},
		}},
		{"TestCase3", &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		}, [][]int{
			{9},
			{3, 15},
			{20},
			{7},
		}},
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

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
