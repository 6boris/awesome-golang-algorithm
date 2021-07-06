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
		root *TreeNode
		expect int
	}{
		{"TestCase", &TreeNode{
			Val: 1, 
			Left: &TreeNode{Val: 2}, 
			Right: &TreeNode{Val: 3},
		}, 6},
		{"TestCase", &TreeNode{
			-10,
			&TreeNode{Val: 9},
			&TreeNode{
				20,
				&TreeNode{Val: 15},
				&TreeNode{Val: 7},
			},
		}, 42},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.root)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.root)
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
