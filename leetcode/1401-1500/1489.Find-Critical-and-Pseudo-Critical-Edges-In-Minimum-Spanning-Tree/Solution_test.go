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
		expect [][]int
	}{
		{"TestCase1", 5, [][]int{
			{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6},
		}, [][]int{
			{0, 1}, {2, 3, 4, 5},
		}},
		{"TestCase2", 4, [][]int{
			{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 1},
		}, [][]int{
			{}, {0, 1, 2, 3},
		}},
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
