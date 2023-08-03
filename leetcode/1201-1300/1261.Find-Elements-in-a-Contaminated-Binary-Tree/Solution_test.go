package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name    string
		inputs  *TreeNode
		targets []int
		expect  []bool
	}{
		{"TestCase1", &TreeNode{Val: -1, Right: &TreeNode{Val: -1}}, []int{1, 2}, []bool{false, true}},
		{"TestCase2", &TreeNode{
			Val: -1,
			Left: &TreeNode{
				Val:   -1,
				Left:  &TreeNode{Val: -1},
				Right: &TreeNode{Val: -1},
			},
			Right: &TreeNode{Val: -1},
		}, []int{1, 3, 5}, []bool{true, true, false}},
		{"TestCase3", &TreeNode{
			Val: -1,
			Right: &TreeNode{
				Val: -1,
				Left: &TreeNode{
					Val:  -1,
					Left: &TreeNode{Val: -1},
				},
			},
		}, []int{2, 3, 4, 5}, []bool{true, false, false, true}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.targets)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.targets)
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
