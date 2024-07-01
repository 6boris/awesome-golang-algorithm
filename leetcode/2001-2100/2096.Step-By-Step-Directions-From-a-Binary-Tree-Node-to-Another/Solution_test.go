package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name        string
		tree        *TreeNode
		start, dest int
		expect      string
	}{
		{"TestCase1", &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 3},
			},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 4},
			},
		}, 3, 6, "UURL"},
		{"TestCase2", &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, 2, 1, "L"},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.tree, c.start, c.dest)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.tree, c.start, c.dest)
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
