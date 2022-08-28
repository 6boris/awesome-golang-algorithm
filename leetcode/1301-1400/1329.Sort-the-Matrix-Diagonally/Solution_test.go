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
		inputs [][]int
		expect [][]int
	}{
		{"TestCase1", [][]int{
			{3, 3, 1, 1},
			{2, 2, 1, 2},
			{1, 1, 1, 2},
		}, [][]int{
			{1, 1, 1, 1},
			{1, 2, 2, 2},
			{1, 2, 3, 3},
		}},
		{"TestCase2", [][]int{
			{11, 25, 66, 1, 69, 7},
			{23, 55, 17, 45, 15, 52},
			{75, 31, 36, 44, 58, 8},
			{22, 27, 33, 25, 68, 4},
			{84, 28, 14, 11, 5, 50},
		}, [][]int{
			{5, 17, 4, 1, 52, 7},
			{11, 11, 25, 45, 8, 69},
			{14, 23, 25, 44, 58, 15},
			{22, 27, 31, 36, 50, 66},
			{84, 28, 75, 33, 55, 68},
		}},
		{"TestCase3", [][]int{{1}}, [][]int{{1}}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
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
