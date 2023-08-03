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
		n        int
		requests [][]int
		expect   int
	}{
		{"TestCase1", 5, [][]int{
			{0, 1}, {1, 0}, {0, 1}, {1, 2}, {2, 0}, {3, 4},
		}, 5},
		{"TestCase2", 3, [][]int{
			{0, 0}, {1, 2}, {2, 1},
		}, 3},
		{"TestCase3", 4, [][]int{
			{0, 3}, {3, 1}, {1, 2}, {2, 0},
		}, 4},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.requests)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.n, c.requests)
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
