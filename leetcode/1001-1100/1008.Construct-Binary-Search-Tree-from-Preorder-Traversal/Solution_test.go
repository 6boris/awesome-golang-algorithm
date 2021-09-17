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
		inputs []int
		expect *TreeNode
	}{
		{"TestCase1", []int{1}, &TreeNode{Val: 1}},
		{"TestCase2", []int{8, 9, 10}, &TreeNode{Val: 8, Right: &TreeNode{Val: 9, Right: &TreeNode{Val: 10}}}},
		{"TestCase3", []int{10, 9, 8}, &TreeNode{Val: 10, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 8}}}},
		{"TestCase4", []int{8, 5, 1, 7, 10, 12}, &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 7},
			},
			Right: &TreeNode{
				Val:   10,
				Right: &TreeNode{Val: 12},
			},
		}},
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

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
