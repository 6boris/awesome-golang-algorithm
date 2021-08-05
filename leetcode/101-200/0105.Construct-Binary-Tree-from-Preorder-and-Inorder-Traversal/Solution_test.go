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
		preorder []int 
		inorder []int
		expect *TreeNode
	}{
		{"TestCase", []int{3,9,20,15,7}, []int{9,3,15,20,7}, &TreeNode{3, &TreeNode{Val: 9}, &TreeNode{20, &TreeNode{Val: 15}, &TreeNode{Val: 7}}}},
		{"TestCase", []int{-1}, []int{-1}, &TreeNode{Val: -1}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.preorder, c.inorder)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.preorder, c.inorder)
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
