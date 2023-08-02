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
		val    int
		expect *TreeNode
	}{
		{"TestCase1", &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{Val: 7},
		}, 5, &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 5}},
		}},
		{"TestCase2", &TreeNode{
			Val: 40,
			Left: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 10},
				Right: &TreeNode{Val: 30},
			},
			Right: &TreeNode{
				Val:   60,
				Left:  &TreeNode{Val: 50},
				Right: &TreeNode{Val: 70},
			},
		}, 25, &TreeNode{
			Val: 40,
			Left: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 10},
				Right: &TreeNode{Val: 30, Left: &TreeNode{Val: 25}},
			},
			Right: &TreeNode{
				Val:   60,
				Left:  &TreeNode{Val: 50},
				Right: &TreeNode{Val: 70},
			},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.val)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.val)
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
