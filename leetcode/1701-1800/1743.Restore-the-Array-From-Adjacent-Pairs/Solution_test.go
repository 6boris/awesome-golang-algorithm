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
		inputs  [][]int
		expect  []int
		expect1 []int
	}{
		{"TestCase1", [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"TestCase2", [][]int{{-2, 4}, {1, 4}, {-3, 1}}, []int{-2, 4, 1, -3}, []int{-3, 1, 4, -2}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) && !reflect.DeepEqual(got, c.expect1) {
				t.Fatalf("expected: %v or %v, but got: %v, with inputs: %v",
					c.expect, c.expect1, got, c.inputs)
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
