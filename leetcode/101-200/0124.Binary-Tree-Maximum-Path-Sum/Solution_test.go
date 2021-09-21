package Solution

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Solution func Info
type SolutionFuncType func(node *TreeNode) int

var SolutionFuncList = []SolutionFuncType{
	maxPathSum_1,
}

// Test case info struct
type Case struct {
	name   string
	input  *TreeNode
	expect int
}

// Test case
var cases = []Case{
	{
		name: "TestCase 1",
		input: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		},
		expect: 6,
	},
	{
		name: "TestCase 2",
		input: &TreeNode{
			-10,
			&TreeNode{Val: 9},
			&TreeNode{
				20,
				&TreeNode{Val: 15},
				&TreeNode{Val: 7},
			},
		},
		expect: 42,
	},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.input)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
