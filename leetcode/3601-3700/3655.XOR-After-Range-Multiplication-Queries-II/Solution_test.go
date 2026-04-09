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
		nums    []int
		queries [][]int
		expect  int
	}{
		{"TestCase1", []int{1, 1, 1}, [][]int{{0, 2, 1, 4}}, 4},
		{"TestCase2", []int{2, 3, 1, 5, 4}, [][]int{
			{1, 4, 2, 3}, {0, 2, 1, 2},
		}, 31},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.nums, c.queries)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.nums, c.queries)
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
