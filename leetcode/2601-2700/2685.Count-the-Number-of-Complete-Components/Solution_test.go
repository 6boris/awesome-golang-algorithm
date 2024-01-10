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
		edges  [][]int
		expect int
	}{
		{"TestCase1", 6, [][]int{
			{0, 1}, {0, 2}, {1, 2}, {3, 4},
		}, 3},
		{"TestCase2", 6, [][]int{
			{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5},
		}, 1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.edges)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.n, c.edges)
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
