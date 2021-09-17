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
type SolutionFuncType func([]int, int) []int

var SolutionFuncList = []SolutionFuncType{
	TwoSum_1,
	TwoSum_2,
	TwoSum_3,
}

// Test case info struct
type Case struct {
	name   string
	nums   []int
	target int
	expect []int
}

// Test case
var cases = []Case{
	{"TestCase 1", nil, 9, nil},
	{"TestCase 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
	{"TestCase 2", []int{3, 2, 4}, 6, []int{1, 2}},
	{"TestCase 3", []int{7, 6, 5, 3, 2, 1, 4, 9, 10}, 17, []int{0, 8}},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
		for _, c := range cases {
			t.Run(fmt.Sprintf("%s %s", funcName, c.name), func(t *testing.T) {
				got := f(c.nums, c.target)
				ast.Equal(c.expect, got,
					"func: %v case: %v ", funcName, c.name)
			})
		}
	}
}

// BenchmarkSolution Run Benchmark for all solutions
func BenchmarkSolution(b *testing.B) {
	ast := assert.New(b)
	for _, f := range SolutionFuncList {
		for i := 0; i < b.N; i++ {
			funcName := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1]
			for _, c := range cases {
				b.ResetTimer()
				b.Run(funcName, func(b *testing.B) {
					got := f(c.nums, c.target)
					ast.Equal(c.expect, got,
						"func: %v case: %v ", funcName, c.name)
				})
			}
		}
	}
}
