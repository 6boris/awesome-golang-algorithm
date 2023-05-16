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
		{"TestCase1", &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 8},
				Right: &TreeNode{Val: 13},
			},
			Right: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 21},
				Right: &TreeNode{Val: 34},
			},
		}, &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 8},
				Right: &TreeNode{Val: 13},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 21},
				Right: &TreeNode{Val: 34},
			},
		}},
		{"TestCase2", &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 13},
			Right: &TreeNode{Val: 11},
		}, &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 11},
			Right: &TreeNode{Val: 13},
		}},
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
func TestSolution1(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs *TreeNode
		expect *TreeNode
	}{
		{"TestCase1", &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 8},
				Right: &TreeNode{Val: 13},
			},
			Right: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 21},
				Right: &TreeNode{Val: 34},
			},
		}, &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 8},
				Right: &TreeNode{Val: 13},
			},
			Right: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 21},
				Right: &TreeNode{Val: 34},
			},
		}},
		{"TestCase2", &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 13},
			Right: &TreeNode{Val: 11},
		}, &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 11},
			Right: &TreeNode{Val: 13},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution1(c.inputs)
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
