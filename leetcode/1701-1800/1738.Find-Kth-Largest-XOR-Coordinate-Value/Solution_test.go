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
		matrix [][]int
		k      int
		expect int
	}{
		{"TestCase1", [][]int{
			{5, 2}, {1, 6},
		}, 1, 7},
		{"TestCase2", [][]int{
			{5, 2}, {1, 6},
		}, 2, 5},
		{"TestCase3", [][]int{
			{5, 2}, {1, 6},
		}, 3, 4},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.matrix, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.matrix, c.k)
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
