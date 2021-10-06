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
type SolutionFuncType func([][]int, int, int, int) [][]int

var SolutionFuncList = []SolutionFuncType{
	Solution,
	floodFill_dfs,
	floodFill_bfs,
}

// Test case info struct
type Case struct {
	name             string
	image            [][]int
	sr, sc, newColor int
	expect           [][]int
}

// Test case
var cases = []Case{
	{"TestCase 1", [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
	{"TestCase 2", [][]int{{1}}, 0, 0, 2, [][]int{{2}}},
	{"TestCase 3", [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, 1, 1, 2, [][]int{{1, 1, 1}, {1, 2, 1}, {1, 1, 1}}},
	{"TestCase 4", [][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 2, [][]int{{2, 2, 2}, {2, 2, 2}}},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				image := [][]int{}
				for _, v := range c.image {
					image = append(image, append([]int{}, v...))
				}
				got := f(image, c.sc, c.sc, c.newColor)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
