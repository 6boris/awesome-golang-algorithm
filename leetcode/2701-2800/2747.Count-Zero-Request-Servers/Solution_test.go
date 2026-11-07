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
		n       int
		logs    [][]int
		x       int
		queries []int
		expect  []int
	}{
		{"TestCase1", 3, [][]int{{1, 3}, {2, 6}, {1, 5}}, 5, []int{10, 11}, []int{1, 2}},
		{"TestCase2", 3, [][]int{{2, 4}, {2, 1}, {1, 2}, {3, 1}}, 2, []int{3, 4}, []int{0, 1}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.logs, c.x, c.queries)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.n, c.logs, c.x, c.queries)
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
