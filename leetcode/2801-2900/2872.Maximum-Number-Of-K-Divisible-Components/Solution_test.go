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
		values []int
		k      int
		expect int
	}{
		{"TestCase1", 5, [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}}, []int{1, 8, 1, 4, 4}, 6, 2},
		{"TestCase2", 7, [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}}, []int{3, 0, 6, 1, 5, 2, 1}, 3, 3},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.edges, c.values, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.n, c.edges, c.values, c.k)
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
