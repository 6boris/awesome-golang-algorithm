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
		start  int
		expect int
	}{
		{"TestCase1", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 10,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
		}, 3, 4},
		{"TestCase2", &TreeNode{Val: 1}, 1, 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.start)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.start)
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
