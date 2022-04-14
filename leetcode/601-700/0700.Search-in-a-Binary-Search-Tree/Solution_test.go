package Solution

import (
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		tree   *TreeNode
		val    int
		expect *TreeNode
	}{
		{"TestCase1", nil, 1, nil},
		{"Testcase2", &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{Val: 7},
		}, 2, &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}},
		{"TestCase3", &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{Val: 7},
		}, 7, &TreeNode{Val: 7}},
		{"TestCase4", &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{Val: 7},
		}, 5, nil},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.tree, c.val)
			if !isSameTree(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v, val: %d",
					c.expect, got, c.tree, c.val)
			}
		})
	}
}

func TestIsSameTree(t *testing.T) {
	t1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 7},
	}
	if !isSameTree(t1, t1) {
		t.Fatalf("expect true get false")
	}

	var t2 *TreeNode
	if isSameTree(t1, t2) {
		t.Fatalf("expect false get ture")
	}

	t2 = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 8},
	}
	if isSameTree(t1, t2) {
		t.Fatalf("expect false get true")
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
