package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name      string
		n         int
		edges     [][]int
		hasApples []bool
		expect    int
	}{
		{"TestCase", 7, [][]int{
			{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6},
		}, []bool{false, false, true, false, false, true, false}, 6},
		{"TestCase", 7, [][]int{
			{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6},
		}, []bool{false, false, true, false, true, true, false}, 8},
		{"TestCase", 1, [][]int{}, []bool{true}, 0},
		{"TestCase", 4, [][]int{{0, 1}, {1, 2}, {0, 3}}, []bool{true, true, true, true}, 6},
		{"TestCase", 4, [][]int{{0, 2}, {0, 3}, {1, 2}}, []bool{false, true, false, false}, 4},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.edges, c.hasApples)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.n, c.edges, c.hasApples)
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
