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
		target int
		expect bool
	}{
		{"TestCase1", &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   6,
				Right: &TreeNode{Val: 7},
			},
		}, 9, true},
		{"TestCase2", &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   6,
				Right: &TreeNode{Val: 7},
			},
		}, 28, false},
		{"TestCase3", &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, 4, true},
		{"TestCase4", &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, 1, false},
		{"TestCase4", &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, 3, true},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.target)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v, target: %v",
					c.expect, got, c.inputs, c.target)
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
