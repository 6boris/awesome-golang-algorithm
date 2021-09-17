package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func([][]byte, string) bool

var SolutionFuncList = []SolutionFuncType{
	exist,
}

//	test info struct
type Case struct {
	name   string
	board  [][]byte
	word   string
	expect bool
}

// 	test case
var cases = []Case{
	{
		name: "TestCase 1",
		board: [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		},
		word:   "ABCCED",
		expect: true,
	},
	{
		name: "TestCase 2",
		board: [][]byte{
			{'a', 'b'},
			{'c', 'd'},
		},
		word:   "abcd",
		expect: false,
	},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				actual := f(c.board, c.word)
				ast.Equal(c.expect, actual,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
