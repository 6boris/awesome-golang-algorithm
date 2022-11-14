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
		expect int
	}{
		{"TestCase1", [][]int{
			{0, 0},
			{0, 1},
			{1, 0},
			{1, 2},
			{2, 1},
			{2, 2},
		}, 5},
		{"TestCase2", [][]int{
			{0, 0},
			{0, 2},
			{1, 1},
			{2, 0},
			{2, 2},
		}, 3},
		{"TestCase3", [][]int{
			{0, 0},
		}, 0},
		{"TestCase4", [][]int{
			{3, 2},
			{3, 1},
			{4, 4},
			{1, 1},
			{0, 2},
			{4, 0},
		}, 4},
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
