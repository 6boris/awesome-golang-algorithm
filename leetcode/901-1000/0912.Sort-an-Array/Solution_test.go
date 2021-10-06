package Solution

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//	solution func Info
type SolutionFuncType func([]int) []int

var SolutionFuncList = []SolutionFuncType{
	quickSort,
	heapSort,
	mergeSort,
	selectSort,
	insertSort,
	bubbleSort,
	shellSort,
	countingSort,
}

//	test info struct
type Case struct {
	name   string
	inputs []int
	expect []int
}

// 	Test case
var cases = []Case{
	{name: "TestCase 1", inputs: []int{5, 2, 3, 1}, expect: []int{1, 2, 3, 5}},
	{name: "TestCase 2", inputs: []int{5, 1, 1, 2, 0, 0}, expect: []int{0, 0, 1, 1, 2, 5}},
}

// TestSolution Example for solution test cases
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				nums := []int{}
				nums = append(nums, c.inputs...)
				got := f(nums)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}
