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
		connections [][]int
		expect      int
	}{
		{"TestCase1", 6, [][]int{
			{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5},
		}, 3},
		{"TestCase2", 5, [][]int{
			{1, 0}, {1, 2}, {3, 2}, {3, 4},
		}, 2},
		{"TestCase3", 3, [][]int{
			{1, 0}, {2, 0},
		}, 0},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.connections)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.n, c.connections)
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
