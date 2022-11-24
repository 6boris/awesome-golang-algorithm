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
		expect []int
	}{
		{"TestCase1", [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}, []int{1, 2, 4, 7, 5, 3, 6, 8, 9}},
		{"TestCase2", [][]int{
			{1, 2},
			{3, 4},
		}, []int{1, 2, 3, 4}},
		{"TestCase3", [][]int{{1}}, []int{1}},
		{"TestCase4", [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		}, []int{1, 2, 5, 6, 3, 4, 7, 8}},
		{"TestCase5", [][]int{
			{1, 2},
			{3, 4},
			{5, 6},
			{7, 8},
		}, []int{1, 2, 3, 5, 4, 6, 7, 8}},
		{"TestCase6", [][]int{
			{1},
			{2},
			{3},
			{4},
		}, []int{1, 2, 3, 4}},
		{"TestCase7", [][]int{
			{1, 2, 3, 4},
		}, []int{1, 2, 3, 4}},
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
