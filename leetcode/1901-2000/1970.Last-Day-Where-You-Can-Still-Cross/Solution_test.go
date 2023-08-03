package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name     string
		row, col int
		cells    [][]int
		expect   int
	}{
		{"TestCase1", 2, 2, [][]int{
			{1, 1}, {2, 1}, {1, 2}, {2, 2},
		}, 2},
		{"TestCase2", 2, 2, [][]int{
			{1, 1}, {1, 2}, {2, 1}, {2, 2},
		}, 1},
		{"TestCase3", 3, 3, [][]int{
			{1, 2}, {2, 1}, {3, 3}, {2, 2}, {1, 1}, {1, 3}, {2, 3}, {3, 2}, {3, 1},
		}, 3},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.row, c.col, c.cells)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.row, c.col, c.cells)
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
