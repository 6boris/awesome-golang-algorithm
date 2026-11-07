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
		edges   [][]int
		s       string
		queries []string
		expect  []bool
	}{
		{"TestCase1", 3, [][]int{{0, 1}, {1, 2}}, "aac", []string{"query 0 2", "update 1 b", "quer 0 2"}, []bool{true, false}},
		{"TestCase2", 4, [][]int{{0, 1}, {0, 2}, {0, 3}}, "abca", []string{"query 1 2", "update 0 b", "query 2 3", "update 3 a", "query 1 3"}, []bool{false, false, true}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.edges, c.s, c.queries)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.n, c.edges, c.s, c.queries)
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
