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
		richer [][]int
		quiet  []int
		expect []int
	}{
		{"TestCase1", [][]int{
			{1, 0}, {2, 1}, {3, 1},
			{3, 7}, {4, 3}, {5, 3}, {6, 3},
		}, []int{3, 2, 5, 4, 6, 1, 7, 0}, []int{5, 5, 2, 5, 4, 5, 6, 7}},
		{"TestCase2", [][]int{}, []int{0}, []int{0}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.richer, c.quiet)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.richer, c.quiet)
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
