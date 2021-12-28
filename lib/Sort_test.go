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
type SolutionFuncType func([]int) []int

var SolutionFuncList = []SolutionFuncType{
	BubbleSort,
	InsertSort,
	SelectSort,
	QuickSort,
	MergeSort,
	HeapSort,
	ShellSort,
	CocktailShakerSort,
	CombSort,
	CountingSort,
	GnomeSort,
	OddEvenSort,
	sortArray,
}

// Test case info struct
type Case struct {
	name   string
	input  []int
	expect []int
}

// Test case
var cases = []Case{
	{name: "TestCase 1", input: []int{2, 4, 1, 3, 8, 5, 9, 6, 7}, expect: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]

		for _, c := range cases {
			num := make([]int, len(c.input))
			copy(num, c.input)
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(num)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
