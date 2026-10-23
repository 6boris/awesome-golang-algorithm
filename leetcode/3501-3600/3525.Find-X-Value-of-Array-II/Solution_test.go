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
		nums    []int
		k       int
		queries [][]int
		expect  []int
	}{
		{"TestCase1", []int{1, 2, 3, 4, 5}, 3, [][]int{{2, 2, 0, 2}, {3, 3, 3, 0}, {0, 1, 0, 1}}, []int{2, 2, 2}},
		{"TestCase2", []int{1, 2, 4, 8, 16, 32}, 4, [][]int{{0, 2, 0, 2}, {0, 2, 0, 1}}, []int{1, 0}},
		{"TestCase3", []int{1, 1, 2, 1, 1}, 2, [][]int{{2, 1, 0, 1}}, []int{5}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums, c.k, c.queries)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.nums, c.k, c.queries)
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
