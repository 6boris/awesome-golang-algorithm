package Solution

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"testing"
)

//	solution func Info
type SolutionFuncType func(*TreeNode) bool

var SolutionFuncList = []SolutionFuncType{
	isSymmetric,
}

//	test info struct
type Case struct {
	name   string
	root   *TreeNode
	expect bool
}

// 	test case
var cases = []Case{
	{
		name: "TestCase 1",
		root: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		expect: false,
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.root)
				ast.Equal(c.expect, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
