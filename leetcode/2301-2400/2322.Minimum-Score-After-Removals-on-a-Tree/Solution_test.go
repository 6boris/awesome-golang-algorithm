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
		nums   []int
		edges  [][]int
		expect int
	}{
		{"TestCase1", []int{1, 5, 5, 4, 11}, [][]int{
			{0, 1}, {1, 2}, {1, 3}, {3, 4},
		}, 9},
		{"TestCase2", []int{5, 5, 2, 4, 4, 2}, [][]int{
			{0, 1}, {1, 2}, {5, 2}, {4, 3}, {1, 3},
		}, 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums, c.edges)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.nums, c.edges)
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
