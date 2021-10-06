package Solution

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test case info struct
type Case struct {
	name string
	s    string
	root *TreeNode
}

// Test case
var cases = []Case{
	{
		name: "TestCase 1",
		s:    "1,2,null,null,3,4,null,null,5,null,null",
		root: &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5}},
		},
	},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)
	con := Constructor()
	for _, c := range cases {
		t.Run(fmt.Sprintf("%s %s", "", c.name), func(t *testing.T) {
			ast.Equal(c.s, con.serialize(con.deserialize(c.s)),
				"func: %v case: %v ", "serialize", c.name)
			ast.Equal(c.root, con.deserialize(con.serialize(c.root)),
				"func: %v case: %v ", "deserialize", c.name)
		})
	}
}
