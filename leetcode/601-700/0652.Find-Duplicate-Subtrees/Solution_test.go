package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例

	n1 := &TreeNode{Val: 4}
	n2 := &TreeNode{Val: 2, Left: n1}

	n3 := &TreeNode{Val: 1}

	n4 := &TreeNode{Val: 3}
	n5 := &TreeNode{Val: 2, Left: n4}
	cases := []struct {
		name   string
		inputs *TreeNode
		expect []*TreeNode
	}{
		{"TestCase", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 4},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  n2,
				Right: &TreeNode{Val: 4},
			},
		}, []*TreeNode{n1, n2}},
		{"TestCase2", &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: n3,
		}, []*TreeNode{n3}},
		{"TestCase3", &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 3},
			},
			Right: n5,
		}, []*TreeNode{n4, n5}},
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

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
