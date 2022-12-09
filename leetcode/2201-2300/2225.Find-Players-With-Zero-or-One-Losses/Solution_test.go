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
		expect [][]int
	}{
		{"TestCase1", [][]int{
			{1, 3},
			{2, 3},
			{3, 6},
			{5, 6},
			{5, 7},
			{4, 5},
			{4, 8},
			{4, 9},
			{10, 4},
			{10, 9},
		}, [][]int{
			{1, 2, 10},
			{4, 5, 7, 8},
		}},
		{"TestCase2", [][]int{
			{2, 3},
			{1, 3},
			{5, 4},
			{6, 4},
		}, [][]int{{1, 2, 5, 6}, {}}},
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
