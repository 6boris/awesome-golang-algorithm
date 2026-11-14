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
		m, n   int
		ops    [][]int
		expect int
	}{
		{"TestCase1", 3, 3, [][]int{{2, 2}, {3, 3}}, 4},
		{"TestCase2", 3, 3, [][]int{
			{2, 2}, {3, 3}, {3, 3}, {3, 3},
			{2, 2}, {3, 3}, {3, 3}, {3, 3},
			{2, 2}, {3, 3}, {3, 3}, {3, 3},
		}, 4},
		{"TestCase3", 3, 3, [][]int{}, 9},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.m, c.n, c.ops)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.m, c.n, c.ops)
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
