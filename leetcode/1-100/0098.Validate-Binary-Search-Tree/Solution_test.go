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
type SolutionFuncType func(*TreeNode) bool

var SolutionFuncList = []SolutionFuncType{
	isValidBST_1,
	isValidBST_2,
}

// Test case info struct
type Case struct {
	name   string
	input  *TreeNode
	expect bool
}

// Test case
var cases = []Case{
	{name: "TestCase 1", input: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, expect: true},
	{
		name: "TestCase 2",
		input: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 0},
				Right: &TreeNode{Val: 2},
			},
			Right: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 6}},
		},
		expect: true},
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
