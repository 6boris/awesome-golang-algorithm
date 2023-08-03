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
		expect *TreeNode
	}{
		{"TestCase1", &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 3},
				},
			},
			Right: &TreeNode{
				Val:  6,
				Left: &TreeNode{Val: 5},
				Right: &TreeNode{
					Val:   7,
					Right: &TreeNode{Val: 8},
				},
			},
		}, &TreeNode{
			Val: 30,
			Left: &TreeNode{
				Val: 36,
				Left: &TreeNode{
					Val: 36,
				},
				Right: &TreeNode{
					Val:   35,
					Right: &TreeNode{Val: 33},
				},
			},
			Right: &TreeNode{
				Val:  21,
				Left: &TreeNode{Val: 26},
				Right: &TreeNode{
					Val:   15,
					Right: &TreeNode{Val: 8},
				},
			},
		}},
		{"TestCase2", &TreeNode{Val: 0, Right: &TreeNode{Val: 1}}, &TreeNode{Val: 1, Right: &TreeNode{Val: 1}}},
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
