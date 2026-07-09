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
		expect  []bool
	}{
		{"TestCase1", 2, []int{1, 3}, 1, [][]int{{0, 0}, {0, 1}}, []bool{true, false}},
		{"TestCase2", 4, []int{2, 5, 6, 8}, 2, [][]int{{0, 1}, {0, 2}, {1, 3}, {2, 3}}, []bool{false, false, true, true}},
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
