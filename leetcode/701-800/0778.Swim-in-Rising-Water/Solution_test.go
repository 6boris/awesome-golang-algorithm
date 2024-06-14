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
			{0, 2}, {1, 3},
		}, 3},
		{"TestCase1", [][]int{
			{0, 1, 2, 3, 4},
			{24, 23, 22, 21, 5},
			{12, 13, 14, 15, 16},
			{11, 17, 18, 19, 20},
			{10, 9, 8, 7, 6},
		}, 16},
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
