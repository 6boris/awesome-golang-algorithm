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
		expect string
	}{
		{"TestCase1", &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		}, "2,1,3#1,2,3"},
		{"TestCase2", nil, ""},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			tree, str := Solution(c.inputs)
			if !reflect.DeepEqual(tree, c.inputs) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.inputs, tree, c.inputs)
			}
			if !reflect.DeepEqual(str, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, str, c.inputs)
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
