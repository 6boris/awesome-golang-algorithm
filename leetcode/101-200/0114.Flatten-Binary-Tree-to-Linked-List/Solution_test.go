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
		expect *TreeNode
	}{
		{
			"TestCase", &TreeNode{1, &TreeNode{2, &TreeNode{Val: 3}, &TreeNode{Val: 4}}, &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}},
			&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}}}},
		},
		{"TestCase", &TreeNode{}, &TreeNode{}},
		{"TestCase", &TreeNode{Val: 0}, &TreeNode{Val: 0}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			temp := c.inputs
			Solution(c.inputs)
			if !reflect.DeepEqual(c.inputs, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, c.inputs, temp)
			}
		})
	}
}
func TestSolution2(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs *TreeNode
		expect *TreeNode
	}{
		{
			"TestCase", &TreeNode{1, &TreeNode{2, &TreeNode{Val: 3}, &TreeNode{Val: 4}}, &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}},
			&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}}}},
		},
		{"TestCase", &TreeNode{}, &TreeNode{}},
		{"TestCase", &TreeNode{Val: 0}, &TreeNode{Val: 0}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			temp := c.inputs
			Solution2(c.inputs)
			if !reflect.DeepEqual(c.inputs, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, c.inputs, temp)
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
