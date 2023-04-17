package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name     string
		inputs   *TreeNode
		toDelete []int
		expect   []*TreeNode
	}{
		{"TestCase1", &TreeNode{
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
		}, []int{3, 5}, []*TreeNode{{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 4},
			},
		}, {Val: 6}, {Val: 7}}},
		{"TestCase2", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{Val: 4},
		}, []int{3}, []*TreeNode{{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.toDelete)
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
