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
		r1, r2 *TreeNode
		expect *TreeNode
	}{
		{"TestCase1", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 5},
			},
			Right: &TreeNode{Val: 2},
		}, &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 7},
			},
		}, &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 5},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   5,
				Right: &TreeNode{Val: 7},
			},
		}},
		{"TestCase2", &TreeNode{Val: 1}, &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.r1, c.r2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.r1, c.r2)
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
