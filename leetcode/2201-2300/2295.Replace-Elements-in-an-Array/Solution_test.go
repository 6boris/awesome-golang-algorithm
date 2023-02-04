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
		ops    [][]int
		expect []int
	}{
		{"TestCase1", []int{1, 2, 4, 6}, [][]int{{1, 3}, {4, 7}, {6, 1}}, []int{3, 2, 7, 1}},
		{"TestCase2", []int{1, 2}, [][]int{{1, 3}, {2, 1}, {3, 2}}, []int{2, 1}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums, c.ops)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.nums, c.ops)
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
