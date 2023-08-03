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
		inputs string
		expect *TreeNode
	}{
		{"TestCase1", "1-2--3--4-5--6--7", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 6},
				Right: &TreeNode{Val: 7},
			},
		}},
		{"TestCase2", "1-2--3---4-5--6---7", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
				},
			},
			Right: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 6,
					Left: &TreeNode{
						Val: 7,
					},
				},
			},
		}},
		{"TestCase3", "1-401--349---90--88", &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 401,
				Right: &TreeNode{
					Val: 88,
				},
				Left: &TreeNode{
					Val: 349,
					Left: &TreeNode{
						Val: 90,
					},
				},
			},
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

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
