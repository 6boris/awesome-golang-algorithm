package Solution

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Print(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	Print(root.Left)
	Print(root.Right)
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)
	//	test cases
	cases := []struct {
		name   string
		inputs *TreeNode
		expect *TreeNode
	}{
		{"TestCase",
			&TreeNode{1,
				&TreeNode{2,
					&TreeNode{4, nil, nil},
					&TreeNode{5, nil, nil},
				},
				&TreeNode{3, nil, nil},
			},
			&TreeNode{1,
				&TreeNode{3, nil, nil},
				&TreeNode{2,
					&TreeNode{5, nil, nil},
					&TreeNode{4, nil, nil},
				},
			},
		},
	}

	//	Start test
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			actual := Mirror(c.inputs)
			Print(c.inputs)
			fmt.Println()
			Print(c.expect)
			fmt.Println()
			Print(actual)
			ast.Equal(c.expect, actual, "expected: %v, but actual: %v, with inputs: %v",
				c.expect, actual, c.inputs)
		})
	}
}
