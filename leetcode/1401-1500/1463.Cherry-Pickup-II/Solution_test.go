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
			{3, 1, 1}, {2, 5, 1}, {1, 5, 5}, {2, 1, 1},
		}, 24},
		{"TestCase2", [][]int{
			{1, 0, 0, 0, 0, 0, 1},
			{2, 0, 0, 0, 0, 3, 0},
			{2, 0, 9, 0, 0, 0, 0},
			{0, 3, 0, 5, 4, 0, 0},
			{1, 0, 2, 3, 0, 0, 6},
		}, 28},
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
