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
		root   *TreeNode
		opts   []uint8
		expect []result
	}{
		{"TestCase1", &TreeNode{
			Val:  7,
			Left: &TreeNode{Val: 3},
			Right: &TreeNode{
				Val:   15,
				Left:  &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20},
			},
		}, []uint8{1, 1, 0, 1, 0, 1, 0, 1, 0}, []result{
			{intF: 3}, {intF: 7}, {boolF: true},
			{intF: 9}, {boolF: true}, {intF: 15},
			{boolF: true}, {intF: 20}, {boolF: false},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.root, c.opts)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.root, c.opts)
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
