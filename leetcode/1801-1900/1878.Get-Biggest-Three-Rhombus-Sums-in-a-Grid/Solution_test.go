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
		inputs [][]int
		expect []int
	}{
		{"TestCase1", [][]int{
			{3, 4, 5, 1, 3},
			{3, 3, 4, 2, 3},
			{20, 30, 200, 40, 10},
			{1, 5, 5, 4, 1},
			{4, 3, 2, 2, 5},
		}, []int{228, 216, 211}},
		{"TestCase2", [][]int{
			{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
		}, []int{20, 9, 8}},
		{"TestCase3", [][]int{{7, 7, 7}}, []int{7}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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
