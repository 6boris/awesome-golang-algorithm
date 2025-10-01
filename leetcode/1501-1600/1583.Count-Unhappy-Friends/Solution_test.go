package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name        string
		n           int
		preferences [][]int
		pairs       [][]int
		expect      int
	}{
		{"TestCase1", 4, [][]int{
			{1, 2, 3}, {3, 2, 0}, {3, 1, 0}, {1, 2, 0},
		}, [][]int{
			{0, 1}, {2, 3},
		}, 2},
		{"TestCase2", 2, [][]int{
			{1}, {0},
		}, [][]int{{1, 0}}, 0},
		{"TestCase3", 4, [][]int{
			{1, 3, 2}, {2, 3, 0}, {1, 3, 0}, {0, 2, 1},
		}, [][]int{{1, 3}, {0, 2}}, 4},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.preferences, c.pairs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.n, c.preferences, c.pairs)
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
