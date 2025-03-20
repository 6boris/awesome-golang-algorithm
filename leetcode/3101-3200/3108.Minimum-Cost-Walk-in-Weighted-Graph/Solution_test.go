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
		query  [][]int
		expect []int
	}{
		{"TestCase1", 5, [][]int{{0, 1, 7}, {1, 3, 7}, {1, 2, 1}}, [][]int{{0, 3}, {3, 4}}, []int{1, -1}},
		{"TestCase2", 3, [][]int{{0, 2, 7}, {0, 1, 15}, {1, 2, 6}, {1, 2, 1}}, [][]int{{1, 2}}, []int{0}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.edges, c.query)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.n, c.edges, c.query)
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
