package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name              string
		inputs            *TreeNode
		low, high, expect int
	}{
		{"TestCase1", &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 7},
			},
			Right: &TreeNode{
				Val:   15,
				Right: &TreeNode{Val: 18},
			},
		}, 7, 15, 32},
		{"TestCase2", &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val:  7,
					Left: &TreeNode{Val: 6},
				},
			},
			Right: &TreeNode{
				Val:   15,
				Left:  &TreeNode{Val: 13},
				Right: &TreeNode{Val: 18},
			},
		}, 6, 10, 23},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.low, c.high)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %d %d",
					c.expect, got, c.inputs, c.low, c.high)
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
