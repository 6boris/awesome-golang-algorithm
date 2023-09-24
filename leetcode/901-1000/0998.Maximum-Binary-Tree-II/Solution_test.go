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
		val    int
		expect *TreeNode
	}{
		{"TestCase1", &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 2},
			},
		}, 5, &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 2},
				},
			},
		}},
		{"TestCase2", &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{Val: 4},
		}, 3, &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 3}},
		}},
		{"TestCase3", &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		}, 4, &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
				},
			},
			Right: &TreeNode{
				Val:  4,
				Left: &TreeNode{Val: 3},
			},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.val)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.val)
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
