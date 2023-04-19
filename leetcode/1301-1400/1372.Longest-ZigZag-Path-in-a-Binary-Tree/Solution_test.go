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
		expect int
	}{
		{"TestCase1", &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 1},
					Left: &TreeNode{
						Val: 1,
						Right: &TreeNode{
							Val:   1,
							Right: &TreeNode{Val: 1},
						},
					},
				},
			},
		}, 3},
		{"TestCase1", &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 1},
			Left: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 1},
					Left: &TreeNode{
						Val:   1,
						Right: &TreeNode{Val: 1},
					},
				},
			},
		}, 4},
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
