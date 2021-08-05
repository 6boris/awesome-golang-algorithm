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
		expect bool
	}{
		{"TestCase", &TreeNode{3, &TreeNode{Val: 9}, &TreeNode{20, &TreeNode{Val: 15}, &TreeNode{Val: 7}}}, true},
		{"TestCase", &TreeNode{1, &TreeNode{2, &TreeNode{3, &TreeNode{Val: 4}, &TreeNode{Val: 4}}, &TreeNode{Val: 3}}, &TreeNode{Val: 2}}, false},
		{"TestCase", &TreeNode{}, true},
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
