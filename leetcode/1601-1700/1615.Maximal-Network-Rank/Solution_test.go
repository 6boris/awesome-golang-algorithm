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
		n      int
		roads  [][]int
		expect int
	}{
		{"TestCase1", 4, [][]int{
			{0, 1}, {0, 3}, {1, 2}, {1, 3},
		}, 4},
		{"TestCase2", 5, [][]int{
			{0, 1}, {0, 3}, {1, 2}, {1, 3}, {2, 3}, {2, 4},
		}, 5},
		{"TestCase3", 8, [][]int{
			{0, 1}, {1, 2}, {2, 3}, {2, 4}, {5, 6}, {5, 7},
		}, 5},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.roads)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.n, c.roads)
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
