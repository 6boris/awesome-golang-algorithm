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
		inputs int
		edges  [][]int
		labels string
		k      int
		expect int
	}{
		{"TestCase1", 3, [][]int{{0, 1, 1}, {1, 2, 1}, {0, 2, 3}}, "aab", 1, 3},
		{"TestCase2", 3, [][]int{{0, 1, 1}, {1, 2, 1}, {0, 2, 3}}, "aab", 2, 2},
		{"TestCase3", 3, [][]int{{0, 1, 1}, {1, 2, 1}}, "aaa", 2, -1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.edges, c.labels, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.inputs, c.edges, c.labels, c.k)
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
