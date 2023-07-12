package Solution

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]int
		expect []int
	}{
		{"TestCase", [][]int{
			{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {},
		}, []int{2, 4, 5, 6}},
		{"TestCase2", [][]int{
			{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {},
		}, []int{4}},
		{"TestCase3", [][]int{
			{}, {0, 2, 3, 4}, {3}, {4}, {},
		}, []int{0, 1, 2, 3, 4}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			sort.Ints(got)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
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
