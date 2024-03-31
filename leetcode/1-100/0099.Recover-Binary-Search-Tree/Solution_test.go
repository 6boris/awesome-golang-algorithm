package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs *TreeNode
		expect *TreeNode
	}{
		{"TestCacse1", &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}},
		}, &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
		}},
		{"TestCacse2", &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}},
		}, &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}},
		}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			Solution(c.inputs)
			if !reflect.DeepEqual(c.inputs, c.expect) {
				t.Fatalf("expected: %v, but got: %v",
					c.expect, c.inputs)
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
