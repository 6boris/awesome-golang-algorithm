package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name    string
		n       int
		nums    []int
		maxDiff int
		queries [][]int
		expect  []int
	}{
		{"TestCase1", 5, []int{1, 8, 3, 4, 2}, 3, [][]int{{0, 3}, {2, 4}}, []int{1, 1}},
		{"TestCase2", 5, []int{5, 3, 1, 9, 10}, 2, [][]int{{0, 1}, {0, 2}, {2, 3}, {4, 3}}, []int{1, 2, -1, 1}},
		{"TestCase3", 3, []int{3, 6, 1}, 1, [][]int{{0, 0}, {0, 1}, {1, 2}}, []int{0, -1, -1}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.nums, c.maxDiff, c.queries)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.n, c.nums, c.maxDiff, c.queries)
			}
		})
	}
}

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
