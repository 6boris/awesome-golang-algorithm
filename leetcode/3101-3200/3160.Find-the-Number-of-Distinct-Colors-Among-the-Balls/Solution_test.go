package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name    string
		limit   int
		queries [][]int
		expect  []int
	}{
		{"TestCase1", 4, [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}}, []int{1, 2, 2, 3}},
		{"TestCase2", 4, [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}}, []int{1, 2, 2, 3, 4}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.limit, c.queries)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.limit, c.queries)
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
