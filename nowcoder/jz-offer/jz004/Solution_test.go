package Solution

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)
	//	test cases
	cases := []struct {
		name   string
		pre    []int
		in     []int
		expect *TreeNode
	}{
		{"TestCase",
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
			&TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7}},
			},
		},
		{"TestCase",
			[]int{1, 2, 4, 7, 3, 5, 6, 8},
			[]int{4, 7, 2, 1, 5, 3, 8, 6},
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 5},
					Right: &TreeNode{
						Val:  6,
						Left: &TreeNode{Val: 8},
					}},
			},
		},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := reConstructBinaryTree(c.pre, c.in)
			ast.Equal(c.expect, actual)
		})
	}
}
