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
		opts   []opt
		expect []int
	}{
		{"TestCase1", 4, [][]int{
			{0, 2, 5}, {0, 1, 2}, {1, 2, 1}, {3, 0, 3},
		}, []opt{{"s", []int{3, 2}}, {"s", []int{0, 3}}, {"a", []int{1, 3, 4}}, {"s", []int{0, 3}}}, []int{6, -1, 6}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.edges, c.opts)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.n, c.edges, c.opts)
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
