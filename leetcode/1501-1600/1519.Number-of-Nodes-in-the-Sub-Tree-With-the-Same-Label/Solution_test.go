package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		n      int
		edges  [][]int
		labels string
		expect []int
	}{
		{"TestCase1", 7, [][]int{
			{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6},
		}, "abaedcd", []int{2, 1, 1, 1, 1, 1, 1}},
		{"TestCase2", 4, [][]int{
			{0, 1}, {1, 2}, {0, 3},
		}, "bbbb", []int{4, 2, 1, 1}},
		{"TestCase3", 5, [][]int{
			{0, 1}, {0, 2}, {1, 3}, {0, 4},
		}, "aabab", []int{3, 2, 1, 1, 1}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.edges, c.labels)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.n, c.edges, c.labels)
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
