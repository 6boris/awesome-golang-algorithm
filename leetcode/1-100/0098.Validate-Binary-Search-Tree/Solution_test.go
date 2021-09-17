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
		expect bool
	}{
		{
			"TestCacse 1",
			&TreeNode{
				2,
				&TreeNode{1, nil, nil},
				&TreeNode{3, nil, nil},
			},
			true,
		},
		{
			"TestCacse 2",
			&TreeNode{
				5,
				&TreeNode{1, nil, nil},
				&TreeNode{
					4,
					&TreeNode{3, nil, nil},
					&TreeNode{6, nil, nil},
				},
			},
			false,
		},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := isValidBST(c.inputs)
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, ret, c.inputs)
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
