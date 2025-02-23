package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name                string
		preorder, postorder []int
		expect              *TreeNode
	}{
		{"TestCase1", []int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 7, 3, 1}, &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{Val: 5},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{Val: 7},
			},
		}},
		{"TestCase2", []int{1}, []int{1}, &TreeNode{Val: 1}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.preorder, c.postorder)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.preorder, c.postorder)
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
