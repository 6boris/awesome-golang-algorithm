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
		{"TestCase", []int{-10,-3,0,5,9}, &TreeNode{0, &TreeNode{Val: -3, Left: &TreeNode{Val: -10}}, &TreeNode{Val: 9, Left: &TreeNode{Val: 5}}}},
		{"TestCase", []int{1,3}, &TreeNode{Val: 3, Left: &TreeNode{Val: 1}}},
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
