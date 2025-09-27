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
		m       int
		entries [][]int
		ops     []op
		expect  []any
	}{
		{"TestCase1", 3, [][]int{
			{0, 1, 5}, {0, 2, 6}, {0, 3, 7}, {1, 1, 4}, {1, 2, 7}, {2, 1, 5},
		}, []op{
			{name: "search", movie: 1},
			{name: "rent", shop: 0, movie: 1},
			{name: "rent", shop: 1, movie: 2},
			{name: "report"},
			{name: "drop", shop: 1, movie: 2},
			{name: "search", movie: 2},
		}, []any{[]int{1, 0, 2}, [][]int{{0, 1}, {1, 2}}, []int{0, 1}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.m, c.entries, c.ops)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.m, c.entries, c.ops)
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
