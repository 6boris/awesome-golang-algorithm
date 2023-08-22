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
		t1, t2 *TreeNode
		expect bool
	}{
		{"TestCase1", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 4},
				Right: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 7},
					Right: &TreeNode{Val: 8},
				},
			},
			Right: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 6},
			},
		}, &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   3,
				Right: &TreeNode{Val: 6},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val:   5,
					Left:  &TreeNode{Val: 8},
					Right: &TreeNode{Val: 7},
				},
			},
		}, true},
		{"TestCase1", nil, nil, true},
		{"TestCase2", nil, &TreeNode{Val: 1}, false},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.t1, c.t2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.t1, c.t2)
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
