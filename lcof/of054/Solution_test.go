package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func(*TreeNode, int) int

var SolutionFuncList = []SolutionFuncType{
	kthLargest,
}

//	test info struct
type Case struct {
	name   string
	root   *TreeNode
	k      int
	expect int
}

// 	test case
var cases = []Case{
	{
		name: "TestCase 1",
		root: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		k:      1,
		expect: 4,
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.root, c.k)
				ast.Equal(c.expect, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
