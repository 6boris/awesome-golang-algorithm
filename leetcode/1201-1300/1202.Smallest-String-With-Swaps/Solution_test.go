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
		inputs string
		pairs  [][]int
		expect string
	}{
		{"TestCase1", "dcab", [][]int{{0, 3}, {1, 2}}, "bacd"},
		{"TestCase2", "dcab", [][]int{
			{0, 3}, {1, 2}, {0, 2},
		}, "abcd"},
		{"TestCase3", "cba", [][]int{
			{0, 1}, {1, 2},
		}, "abc"},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.pairs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.inputs, c.pairs)
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
