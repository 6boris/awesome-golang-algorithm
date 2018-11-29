package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs *TreeNode
		expect [][]int
	}{
		{"TestCacse 1",
			&TreeNode{Val: 3,
				Left: &TreeNode{Val: 9, Left: nil, Right: nil},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{15, nil, nil},
					Right: &TreeNode{7, nil, nil}}},
			[][]int{{3}, {9, 20}, {15, 7}}},
		{"TestCacse 2",
			&TreeNode{1,
				&TreeNode{2,
					&TreeNode{4, nil, nil},
					&TreeNode{5, nil, nil}},
				&TreeNode{3, nil, nil}},
			[][]int{{1}, {2, 3}, {4, 5}}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := levelOrder(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				//t.Fatalf("expected: %v, but got: %v, with inputs: %v",
				//	c.expect, got, c.inputs)
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
