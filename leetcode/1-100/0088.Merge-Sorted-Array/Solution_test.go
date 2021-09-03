package Solution

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Solution func Info
type SolutionFuncType func(nums1 []int, m int, nums2 []int, n int)

var SolutionFuncList = []SolutionFuncType{
	merge_1,
	merge_2,
	merge_3,
}

// Test case info struct
type Case struct {
	name   string
	nums1  []int
	m      int
	n      int
	nums2  []int
	expect []int
}

// Test case
var cases = []Case{
	{
		name:   "TestCase 1",
		nums1:  []int{1, 2, 3, 0, 0, 0},
		m:      3,
		nums2:  []int{2, 5, 6},
		n:      3,
		expect: []int{1, 2, 2, 3, 5, 6},
	},
	{
		name:   "TestCase 2",
		nums1:  []int{1},
		m:      1,
		nums2:  []int{0},
		n:      0,
		expect: []int{1},
	},
	{
		name:   "TestCase 3",
		nums1:  []int{0},
		m:      0,
		nums2:  []int{1},
		n:      1,
		expect: []int{1},
	},
}

// TestSolution Run test case for all solutions
func TestSolution(t *testing.T) {
	ast := assert.New(t)

	for _, f := range SolutionFuncList {
		for _, c := range cases {
			nums1 := make([]int, len(c.nums1))
			nums2 := make([]int, len(c.nums2))
			copy(nums1, c.nums1)
			copy(nums2, c.nums2)
			t.Run(c.name, func(t *testing.T) {
				f(nums1, c.m, nums2, c.n)
				ast.Equal(c.expect, nums1,
					"func: %v case: %v ",
					runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), c.name)
			})
		}
	}
}
