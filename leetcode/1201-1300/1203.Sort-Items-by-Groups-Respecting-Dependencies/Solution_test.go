package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name        string
		n, m        int
		group       []int
		beforeItems [][]int
		expect      []int
	}{
		{"TestCase1", 8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, [][]int{
			{}, {6}, {5}, {6}, {3, 6}, {}, {}, {},
		}, []int{6, 3, 4, 5, 2, 0, 7, 1}},
		{"TestCase2", 8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, [][]int{
			{}, {6}, {5}, {6}, {3}, {}, {4}, {},
		}, nil},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.m, c.group, c.beforeItems)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v %v",
					c.expect, got, c.n, c.m, c.group, c.beforeItems)
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
