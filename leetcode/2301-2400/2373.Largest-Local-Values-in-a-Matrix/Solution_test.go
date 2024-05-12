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
			{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2},
		}, [][]int{{9, 9}, {8, 6}}},
		{"TestCase2", [][]int{
			{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1},
			{1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1},
		}, [][]int{
			{2, 2, 2}, {2, 2, 2}, {2, 2, 2},
		}},
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
