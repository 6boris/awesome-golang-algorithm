package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name                         string
		k                            int
		rowConditions, colConditions [][]int
		expect                       [][]int
	}{
		{"TestCase1", 3, [][]int{{1, 2}, {3, 2}}, [][]int{{2, 1}, {3, 2}}, [][]int{
			{0, 0, 1},
			{3, 0, 0},
			{0, 2, 0},
		}},
		{"TestCase2", 3, [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 3}}, [][]int{{2, 1}}, nil},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.k, c.rowConditions, c.colConditions)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.k, c.rowConditions, c.colConditions)
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
