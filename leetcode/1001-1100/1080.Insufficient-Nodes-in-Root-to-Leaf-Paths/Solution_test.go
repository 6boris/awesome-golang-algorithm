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
		limit  int
		expect *TreeNode
	}{
		{"TestCase1", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 8,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
				Right: &TreeNode{
					Val:   -99,
					Left:  &TreeNode{Val: -99},
					Right: &TreeNode{Val: -99},
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   -99,
					Left:  &TreeNode{Val: 12},
					Right: &TreeNode{Val: 13},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: -99},
					Right: &TreeNode{Val: 14},
				},
			},
		}, 1, &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 8},
					Right: &TreeNode{Val: 9},
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val:   7,
					Right: &TreeNode{Val: 14},
				},
			},
		}},
		{"TestCase2", &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   11,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 1},
				},
			},
			Right: &TreeNode{
				Val:  8,
				Left: &TreeNode{Val: 17},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 3},
				},
			},
		}, 22, &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:  11,
					Left: &TreeNode{Val: 7},
				},
			},
			Right: &TreeNode{
				Val:   8,
				Left:  &TreeNode{Val: 17},
				Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}},
			},
		}},
		{"TestCase3", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: -5},
			},
			Right: &TreeNode{Val: -3, Left: &TreeNode{Val: 4}},
		}, -1, &TreeNode{Val: 1, Right: &TreeNode{Val: -3, Left: &TreeNode{Val: 4}}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.limit)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.limit)
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
