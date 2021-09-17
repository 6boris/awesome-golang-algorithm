package Solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]int
		expect [][]int
	}{
		{"TestCase 1", [][]int{
			{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
		}, [][]int{
			{7, 4, 1}, {8, 5, 2}, {9, 6, 3},
		}},
		{"TestCase 2", [][]int{
			{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16},
		}, [][]int{
			{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11},
		}},
	}

	//	开始测试
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := [][]int{}
			copy(input, c.inputs)
			Solution(c.inputs)
			if !reflect.DeepEqual(c.inputs, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, c.inputs, input)
			}
		})
	}
}

//	压力测试
func BenchmarkSolution(b *testing.B) {
}

//	使用案列
func ExampleSolution() {
}
