package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name      string
		n         int
		red, blue [][]int
		expect    []int
	}{
		{"TestCase1", 3, [][]int{{0, 1}, {1, 2}}, [][]int{}, []int{0, 1, -1}},
		{"TestCase2", 3, [][]int{{0, 1}}, [][]int{{2, 1}}, []int{0, 1, -1}},
		{"TestCase3", 1, [][]int{{0, 0}}, [][]int{{0, 0}}, []int{0}},
		{"TestCase4", 5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, [][]int{{1, 2}, {2, 3}, {3, 1}}, []int{0, 1, 2, 3, 7}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.red, c.blue)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.n, c.red, c.blue)
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
