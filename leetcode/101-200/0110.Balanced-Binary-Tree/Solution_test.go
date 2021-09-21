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
type SolutionFuncType func(root *TreeNode) bool

var SolutionFuncList = []SolutionFuncType{
	isBalanced_1,
	isBalanced_2,
}

// Test case info struct
type Case struct {
	name   string
	input  *TreeNode
	expect bool
}

// Test case
var cases = []Case{
	{"TestCase 1", &TreeNode{3, &TreeNode{Val: 9}, &TreeNode{20, &TreeNode{Val: 15}, &TreeNode{Val: 7}}}, true},
	{"TestCase 2", &TreeNode{1, &TreeNode{2, &TreeNode{3, &TreeNode{Val: 4}, &TreeNode{Val: 4}}, &TreeNode{Val: 3}}, &TreeNode{Val: 2}}, false},
	{"TestCase 3", &TreeNode{}, true},
	{"TestCase 4",
		&TreeNode{Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{Val: 2},
		},

		false},
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
