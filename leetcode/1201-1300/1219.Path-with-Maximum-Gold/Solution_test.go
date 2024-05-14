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
			{0, 6, 0}, {5, 8, 7}, {0, 9, 0},
		}, 24},
		{"TestCase2", [][]int{
			{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20},
		}, 28},
		{"TestCase3", [][]int{
			{1, 0, 7, 0, 0, 0}, {2, 0, 6, 0, 1, 0}, {3, 5, 6, 7, 4, 2}, {4, 3, 1, 0, 2, 0}, {3, 0, 5, 0, 20, 0},
		}, 60},
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
