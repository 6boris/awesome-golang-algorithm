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
		expect []int
	}{
		{"TestCase1", &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 4},
		}, &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0},
			Right: &TreeNode{Val: 3},
		}, []int{0, 1, 1, 2, 3, 4}},
		{"TestCase2", &TreeNode{Val: 1, Right: &TreeNode{Val: 8}}, &TreeNode{Val: 8, Left: &TreeNode{Val: 1}}, []int{1, 1, 8, 8}},
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
