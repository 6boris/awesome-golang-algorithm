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
		target int
		expect int
	}{
		{"TestCase1", &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 2},
			},
			Right: &TreeNode{
				Val:   -1,
				Right: &TreeNode{Val: 1},
			},
		}, 5, 4},
		{"TestCase2", &TreeNode{Val: 1}, 1, 1},
		{"TestCase3", &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: -2},
				},
				Right: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 1},
				},
			},
			Right: &TreeNode{
				Val:   -3,
				Right: &TreeNode{Val: 11},
			},
		}, 8, 3},
		{"TestCase4", &TreeNode{Val: 1}, 0, 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.tree, c.target)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v, target: %v",
					c.expect, got, c.tree, c.target)
			}

			got = Solution2(c.tree, c.target)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v, target: %v",
					c.expect, got, c.tree, c.target)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
