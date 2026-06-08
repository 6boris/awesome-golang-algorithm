package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name    string
		inputs  [][]int
		x, y, k int
		expect  [][]int
	}{
		{"TestCase1", [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		}, 1, 0, 3, [][]int{
			{1, 2, 3, 4},
			{13, 14, 15, 8},
			{9, 10, 11, 12},
			{5, 6, 7, 16},
		}},
		{"TestCase2", [][]int{
			{3, 4, 2, 3}, {2, 3, 4, 2},
		}, 0, 2, 2, [][]int{
			{3, 4, 4, 2}, {2, 3, 2, 3},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.x, c.y, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.inputs, c.x, c.y, c.k)
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
