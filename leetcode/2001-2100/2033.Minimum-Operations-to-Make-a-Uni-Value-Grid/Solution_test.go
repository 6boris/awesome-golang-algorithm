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
		grid   [][]int
		x      int
		expect int
	}{
		{"TestCase1", [][]int{{2, 4}, {6, 8}}, 2, 4},
		{"TestCase2", [][]int{{1, 5}, {2, 3}}, 1, 5},
		{"TestCase3", [][]int{{1, 2}, {3, 4}}, 2, -1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.grid, c.x)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.grid, c.x)
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
