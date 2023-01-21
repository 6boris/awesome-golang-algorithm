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
		expect []int
	}{
		{"TestCase1", 6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}, []int{8, 12, 6, 10, 10, 10}},
		{"TestCase2", 1, [][]int{}, []int{0}},
		{"TestCase3", 2, [][]int{{1, 0}}, []int{1, 1}},
		{"TestCase4", 4, [][]int{{2, 0}, {3, 1}, {2, 1}}, []int{6, 4, 4, 6}},
		{"TestCase5", 5, [][]int{{0, 4}, {1, 3}, {1, 2}, {0, 2}}, []int{7, 7, 6, 10, 10}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.edges)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.n, c.edges)
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
