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
type SolutionFuncType func([][]int, []int, []int) bool

var SolutionFuncList = []SolutionFuncType{
	hasPath,
}

// Test case info struct
type Case struct {
	name        string
	maze        [][]int
	start       []int
	destination []int
	expect      bool
}

// Test case
var cases = []Case{
	{
		name: "TestCase 1",
		maze: [][]int{
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0},
			{1, 1, 0, 1, 1},
			{0, 0, 0, 0, 0},
		},
		start:       []int{0, 4},
		destination: []int{4, 4},
		expect:      true,
	},
	{
		name: "TestCase 2",
		maze: [][]int{
			{0, 0, 0, 0, 0},
			{1, 1, 0, 0, 1},
			{0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1},
			{0, 1, 0, 0, 0},
		},
		start:       []int{4, 3},
		destination: []int{0, 1},
		expect:      false,
	},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.maze, c.start, c.destination)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
