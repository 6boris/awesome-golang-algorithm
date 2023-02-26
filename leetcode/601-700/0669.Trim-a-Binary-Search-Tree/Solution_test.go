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
		l, h   int
		expect *TreeNode
	}{
		{"TestCase1", &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 2},
		}, 1, 2, &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}},
		{"TestCase2", &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 0,
				Right: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 1},
				},
			},
			Right: &TreeNode{Val: 4},
		}, 1, 3, &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 1},
			},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.l, c.h)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.inputs, c.l, c.h)
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
