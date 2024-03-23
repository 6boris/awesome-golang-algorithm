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
		meetings [][]int
		expect   int
	}{
		{"TestCase1", 2, [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}, 0},
		{"TestCase2", 3, [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}, 1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.meetings)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.n, c.meetings)
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
