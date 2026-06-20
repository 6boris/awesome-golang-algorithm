package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name         string
		n            int
		restrictions [][]int
		expect       int
	}{
		{"TestCase1", 5, [][]int{{2, 1}, {4, 1}}, 2},
		{"TestCase2", 6, [][]int{}, 5},
		{"TestCase3", 10, [][]int{{5, 3}, {2, 5}, {7, 4}, {10, 3}}, 5},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.restrictions)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.n, c.restrictions)
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
