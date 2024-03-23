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
		expect []float64
	}{
		{"TestCase1", &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 9},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		}, []float64{3.00000, 14.50000, 11.00000}},
		{"TestCase", &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   9,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
			Right: &TreeNode{Val: 20},
		}, []float64{3.00000, 14.50000, 11.00000}},
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
