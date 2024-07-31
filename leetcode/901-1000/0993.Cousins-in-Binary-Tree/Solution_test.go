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
		tree   *TreeNode
		x, y   int
		expect bool
	}{
		{"TestCase1", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 4},
			},
			Right: &TreeNode{Val: 3},
		}, 4, 3, false},
		{"TestCase2", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 5},
			},
		}, 5, 4, true},
		{"TestCase3", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{Val: 3},
		}, 2, 3, false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.tree, c.x, c.y)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.tree, c.x, c.y)
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
