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
		n      int
		inputs [][]int
		expect []int
	}{
		{"TestCase1", 6, [][]int{
			{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2},
		}, []int{0, 3}},
		{"TestCase2", 5, [][]int{
			{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4},
		}, []int{0, 2, 3}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.n, c.inputs)
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
